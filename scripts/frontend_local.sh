#!/bin/sh
#
# Runs the frontend server locally.
#
set -e

SCRIPTS=$(dirname $0)

. $SCRIPTS/common.sh

go_generate

go generate ./...
go mod tidy
go run frontend_main.go --port_for_testing 8080 --host localhost "$@"