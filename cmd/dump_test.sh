#!/bin/bash

DB_NAME=$1
DB_TEST_NAME=$1_test
DB_USER=$2
DB_PASSWORD=$3
DB_REMOTE_USER=$4
DB_REMOTE_PASSWORD=$5

psql -Umdgkb -c "SELECT pg_terminate_backend(pg_stat_activity.pid) FROM pg_stat_activity WHERE pg_stat_activity.datname = '$DB_TEST_NAME' AND pid <> pg_backend_pid();"

PGPASSWORD=$DB_PASSWORD createdb -Umdgkb $DB_TEST_NAME
ssh root@45.67.57.208 "pg_dump -s -C -h 45.67.57.208 -d $DB_NAME -U $DB_REMOTE_USER -Fc --password" | PGPASSWORD=$DB_PASSWORD pg_restore -U$DB_USER -hlocalhost --format=c -d$DB_TEST_NAME