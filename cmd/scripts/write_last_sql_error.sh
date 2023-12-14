#!/bin/bash

act_log_file=./logs/errors/_actual.log
last_error_file=./logs/errors/_last_error.sql

tail -c+1 $act_log_file | grep -E 'SQLSTATE' | grep -E 'msg' | awk -F: '{print $3}' | grep -v ERROR | tail -n 1 >$last_error_file

cat $last_error_file
