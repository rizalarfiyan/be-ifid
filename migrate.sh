#!/bin/bash

# Check Command sql-migrate
if ! [ -x "$(command -v sql-migrate)" ]; then
    echo "Command sql-migrate could not be found"
    echo "Installing sql-migrate..."
    go install github.com/rubenv/sql-migrate/...@latest
fi

# Load .env
ENV_VARS="$(cat .env | awk '!/^\s*#/' | awk '!/^\s*$/')"

eval "$(
    printf '%s\n' "$ENV_VARS" | while IFS='' read -r line; do
        key=$(printf '%s\n' "$line"| sed 's/"/\\"/g' | cut -d '=' -f 1)
        value=$(printf '%s\n' "$line" | cut -d '=' -f 2- | sed 's/"/\\\"/g')
        printf '%s\n' "export $key=\"$value\""
    done
)"

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
