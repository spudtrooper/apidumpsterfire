#!/bin/sh
#
# Deploy to heroku in another directory
#
set -e

SCRIPTS=$(dirname $0)
DIRS="../lyft ../resy ../opentable ../opensecrets"

DEPLOY_DIR=../apidumpsterfire-deploy
mkdir -p $DEPLOY_DIR
mkdir -p $DEPLOY_DIR/scripts
cp $SCRIPTS/deploy.sh.txt $DEPLOY_DIR/scripts/deploy-to-keroku.sh

for dir in $DIRS; do
  pushd $dir
  go generate ./...
  popd
done

pushd $DEPLOY_DIR
scripts/deploy-to-keroku.sh
popd