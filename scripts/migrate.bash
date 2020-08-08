#!/bin/bash

TLD=$(git rev-parse --show-toplevel)

MIGRATION_DIR=${TLD}/internal/pkg/db/migrations

migrate -database "postgres://pp:password@0.0.0.0:5432/plantparenthood?sslmode=disable" -path ${MIGRATION_DIR} up
