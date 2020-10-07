#!/bin/bash

TLD=$(git rev-parse --show-toplevel)

MIGRATION_DIR=${TLD}/internal/pkg/db/migrations

# 1 = user
USER=pp
if [[ ! -z "$1" ]]; then
    USER="$1"
fi

# 2 = pass
PASS=password
if [[ ! -z "$2" ]]; then
    PASS="$2"
fi

# 3 = db
DB=plantparenthood
if [[ ! -z "$3" ]]; then
    DB="$3"
fi

# 4 = port
PORT=5432
if [[ ! -z "$4" ]]; then
    PORT="$4"
fi

migrate -database "postgres://${USER}:${PASS}@0.0.0.0:${PORT}/${DB}?sslmode=disable" -path ${MIGRATION_DIR} up
