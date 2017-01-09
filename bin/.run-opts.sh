#!/usr/bin/env bash

options=''

# running interactive?
tty -s
TTY=$?

if [ $TTY -eq 0 ]; then
    options+=' -t'
fi