#!/bin/bash

set -e

psql -v ON_ERROR_STOP=1 --user "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
	DROP DATABASE IF EXISTS j0eppp;
	CREATE DATABASE j0eppp;
EOSQL
