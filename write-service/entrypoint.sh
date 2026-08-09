#!/bin/sh
set -e

if [ -n "$DATABASE_URL" ]; then
    echo "Running migrations..."
    # migrate exits 0 and logs "no change" when nothing pending -- safe to run every start
    if ! migrate -path /app/migrations -database "$DATABASE_URL" up; then
        echo "Migration failed"
        exit 1
    fi
else
    echo "DATABASE_URL not set, skipping migrations"
fi

exec "$@"
