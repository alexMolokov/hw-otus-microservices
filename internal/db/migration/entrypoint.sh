#!/bin/bash

DBSTRING="host=$DB_HOST user=$DB_USER password=$DB_PASSWORD dbname=$DB_NAME sslmode=$DB_SSL"

goose postgres "$DBSTRING" up