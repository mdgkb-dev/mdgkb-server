#!/bin/sh
PGPASSWORD=$1 dropdb -Umdgkb -hlocalhost mdgkb
PGPASSWORD=$1 createdb -Umdgkb mdgkb
ssh root@45.67.57.208 "pg_dump -C -h 45.67.57.208 -d mdgkb -U mdgkb -Fc --password" | PGPASSWORD=$1 pg_restore -Umdgkb -hlocalhost --format=c -dmdgkb


