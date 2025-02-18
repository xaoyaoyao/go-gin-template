#!/bin/bash
set -e

readonly service="$1"
readonly output_dir="$2"
readonly package="$3"

oapi-codegen -generate types -o "$output_dir/api_types.gen.go" -package "$package" "api/$service.yaml"
oapi-codegen -generate gin -o "$output_dir/api_gin.gen.go" -package "$package" "api/$service.yaml"
oapi-codegen -generate client -o "$output_dir/api_client.gen.go" -package "$package" "api/$service.yaml"
