#!/bin/bash

CONTAINER_NAME=test_pp_db

PP_USER=test_pp
PP_PASS=test_pass
PP_DB=test_db

function clean() {
    docker kill ${CONTAINER_NAME}
    docker rm ${CONTAINER_NAME}
}

function wait_for_db() {
    RETRIES=5

    until PGPASSWORD=$PP_PASS psql -h 0.0.0.0 -p 5433 -U $PP_USER -d $PP_DB -c "select 1" > /dev/null 2>&1 || [[ $RETRIES -eq 0 ]]; do
        echo "Waiting for postgres server, $((RETRIES--)) remaining attempts..."
        sleep 1
    done

    if [[ $RETRIES -eq 0 ]]; then
        echo "failed to start database"
        exit 1
    fi

    echo "connected"
}

function migrate() {
    pushd $(git rev-parse --show-toplevel)/scripts
    ./migrate.bash ${PP_USER} ${PP_PASS} ${PP_DB} 5433
    popd
}

clean

docker run --name ${CONTAINER_NAME} -p "5433:5432" -e POSTGRES_USER=${PP_USER} -e POSTGRES_PASSWORD=${PP_PASS} -e POSTGRES_DB=${PP_DB} -d postgres

wait_for_db

migrate

if [[ "$1" == "docker" ]]; then
    echo "Not running tests..."
    exit 0
fi

pushd $(git rev-parse --show-toplevel)
PP_E2E_TEST=1 go test ./...
popd

 clean
