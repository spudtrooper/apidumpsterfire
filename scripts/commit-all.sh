#!/bin/sh
#
# Commits everything.
#
set -e

SCRIPTS=$(dirname $0)

. $SCRIPTS/common.sh

for dir in $DIRS; do
  pushd $dir
  ./scripts/commit.sh "$@"
  popd
done

$SCRIPTS/common.sh "$@"