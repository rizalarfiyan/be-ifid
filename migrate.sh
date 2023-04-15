#!/bin/bash

# Check Command sql-migrate
if ! [ -x "$(command -v sql-migrate)" ]; then
    echo "Command sql-migrate could not be found"
    echo "Installing sql-migrate..."
    go install github.com/rubenv/sql-migrate/...@latest
fi

# Load .env
unamestr=$(uname)
if [ "$unamestr" = 'Linux' ]; then
    export $(grep -v '^#' .env | xargs -d '\n')
elif [ "$unamestr" = 'FreeBSD' ] || [ "$unamestr" = 'Darwin' ]; then
    export $(grep -v '^#' .env | xargs -0)
fi

# Do Action
OPTIONS="-config=dbconfig.yml -env=be_ifid"

case "$1" in
    "new")
    sql-migrate new $OPTIONS $2
    ;;
    "up")
    sql-migrate up $OPTIONS
    ;;
    "redo")
    sql-migrate redo $OPTIONS
    ;;
    "status")
    sql-migrate status $OPTIONS
    ;;
    "down")
    sql-migrate down $OPTIONS
    ;;
    *)
    echo "Usage: $(basename "$0") new {name}/up/status/down"
    exit 1
esac
