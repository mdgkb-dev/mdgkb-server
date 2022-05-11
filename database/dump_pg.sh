#!/bin/sh

psql -Umdgkb -c "SELECT pg_terminate_backend(pg_stat_activity.pid) FROM pg_stat_activity WHERE pg_stat_activity.datname = 'mdgkb' AND pid <> pg_backend_pid();"

PGPASSWORD=$1 dropdb -Umdgkb -hlocalhost mdgkb
PGPASSWORD=$1 createdb -Umdgkb mdgkb
ssh root@45.67.57.208 "pg_dump -C -h 45.67.57.208 -d mdgkb -U mdgkb -Fc --password" | PGPASSWORD=$1 pg_restore -Umdgkb -hlocalhost --format=c -dmdgkb