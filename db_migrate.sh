#!/bin/bash

source .env

MIGRATIONS_DIR="./dao/migrations"

run_migration() {
    migrate -database "postgres://${DB_USER}:${DB_PASSWORD}@localhost:5432/${DB_NAME}?sslmode=disable" -path "${MIGRATIONS_DIR}" "$1"
}

if ! command -v migrate &> /dev/null; then
    echo "Error: 'migrate' command not found. Please install it by following the instructions on https://github.com/golang-migrate/migrate."
    exit 1
fi

if [ "$#" -ne 1 ]; then
    echo "Usage: $0 [up|down]"
    exit 1
fi

case $1 in
    up)
        echo "Migrating up..."
        run_migration up
        ;;
    down)
        echo "Migrating down..."
        run_migration down
        ;;
    *)
        echo "Invalid argument. Use 'up' or 'down'."
        exit 1
        ;;
esac

echo "Migration completed."
