#!/bin/sh

DIRS="../lyft ../resy ../opentable ../opensecrets ../minimalcli ../uber ../spotifydown"

function go_generate() {
  for dir in $DIRS; do
    pushd $dir
    go generate ./...
    popd
  done
}
