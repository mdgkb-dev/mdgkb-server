#!/bin/bash

./cmd/server/get_params.sh

DEPLOY_PATH=$1

PROCESS_NAME=mdgkb-server
PID_FILE=$DEPLOY_PATH/$PROCESS_NAME.pid
PROCESS_FILE=$DEPLOY_PATH/$PROCESS_NAME

echo "$PID_FILE"
GOOS=linux GOARCH=amd64 go build -o "$PROCESS_FILE" ./cmd/server/*.go

if [ -f "$PID_FILE" ]; then
    kill -9 "$(cat "$PID_FILE")" && rm -f "$PID_FILE"
fi

nohup "$PROCESS_FILE" > /dev/null 2>&1 & echo $! > "$PID_FILE"
exit
