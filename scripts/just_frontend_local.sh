#!/bin/sh
#
# Just runs the frontend server locally.
#
set -e

go mod tidy
go run frontend_main.go --port_for_testing 8080 --host localhost "$@"