#!/bin/bash

output_file="db/schema.sql"
build_dir="backend/infra/repository/db/build"
sqlc_config="backend/infra/repository/db/sqlc.yaml"

mkdir -p "$build_dir"
cp "$output_file" "$build_dir/"

sqlc generate -f "$sqlc_config"

cd backend 
swag init
