#!/bin/sh
# pg_dump -Fc -c -Umdgkb  -h45.67.57.208 -dmdgkb | pg_restore -Umdgkb -hlocalhost -Fc -c --if-exists -dmdgkb

#PGPASSWORD=123 dropdb -h localhost -Umdgkb mdgkb
#PGPASSWORD=123 createdb -h localhost -Umdgkb mdgkb
ssh root@45.67.57.208 "pg_dump -h 45.67.57.208 -d mdgkb -U mdgkb -Fc --password"  | PGPASSWORD=123 pg_restore -Umdgkb -hlocalhost --format=c --clean -dmdgkb

