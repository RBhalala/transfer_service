#!/bin/bash
set -e

PGHOST=localhost
PGPORT=5432
PGUSER=root
PGPASSWORD=root
PGDATABASE=transfer_service

export PGPASSWORD

echo "Running migrations..."
psql -h "$PGHOST" -p "$PGPORT" -U "$PGUSER" -d "$PGDATABASE" -f migrations/001_create_tables.sql

echo "Migrations completed."
