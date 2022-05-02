#!/bin/bash

#branch=$1
#if [ -z $branch ]
# then
#    branch="develop"
#fi
#echo $branch


for ARGUMENT in "$@"
do
   KEY=$(echo $ARGUMENT | cut -f1 -d=)
   KEY_LENGTH=${#KEY}
   VALUE="${ARGUMENT:$KEY_LENGTH+1}"
   export "$KEY"="$VALUE"
done

PROCESS_NAME=mdgkb-server
PIDFILE=${BIN_PATH}/${PROCESS_NAME}.pid
PROCESS_FILE=${BIN_PATH}/${PROCESS_NAME}
##git reset --hard && \
##git pull --all && \
##git checkout $branch && \
#export GO111MODULE=on && \
go build -o $PROCESS_FILE ./cmd/server/main.go && \
#if [ -f "$PIDFILE" ]; then
#    echo "$FILE exists."
#    kill -9 `cat ${PIDFILE}` && rm -f "${PIDFILE}"
#fi
#echo $PROCESS_FILE
#
nohup "$PROCESS_FILE" &
exit
