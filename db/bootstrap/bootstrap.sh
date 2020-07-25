#!/usr/bin/env bash

set -euf -o pipefail

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
    CREATE USER expsvcuser with password '$EXPUSER_PW';
    GRANT ALL PRIVILEGES ON DATABASE expense_db TO expsvcuser;
    CREATE SCHEMA IF NOT EXISTS expense_schema AUTHORIZATION expsvcuser;
EOSQL
