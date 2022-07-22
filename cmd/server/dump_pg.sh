#!/bin/bash

DB_NAME=$1
DB_USER=$2
DB_PASSWORD=$3
REMOTE_DB_PASSWORD=$4
REMOTE_DB_USER=$5

psql -Umdgkb -c "SELECT pg_terminate_backend(pg_stat_activity.pid) FROM pg_stat_activity WHERE pg_stat_activity.datname = '$DB_NAME' AND pid <> pg_backend_pid();"

PGPASSWORD=$DB_PASSWORD dropdb -Umdgkb -hlocalhost $DB_NAME
PGPASSWORD=$DB_PASSWORD createdb -Umdgkb $DB_NAME
ssh root@45.67.57.208 "pg_dump -C -h 45.67.57.208 -d $REMOTE_DB_PASSWORD -U $REMOTE_DB_USER -Fc --password" | PGPASSWORD=$DB_PASSWORD pg_restore -U$DB_USER -hlocalhost --format=c -d$DB_NAME