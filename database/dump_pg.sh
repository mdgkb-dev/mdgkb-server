#!/bin/sh
pg_dump -Fc -c -Umdgkb  -h45.67.57.208 -dmdgkb | pg_restore -Umdgkb -hlocalhost -Fc -c --if-exists -dmdgkb
