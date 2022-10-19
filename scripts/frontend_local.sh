#!/bin/sh
#
# Runs the frontend server locally.
#
set -e

SCRIPTS=$(dirname $0)

. $SCRIPTS/common.sh

go_generate

go generate ./...
$SCRIPTS/just_frontend_local.sh "$@"