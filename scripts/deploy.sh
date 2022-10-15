#!/bin/sh
#
# Deploy to heroku in another directory
#
set -e

SCRIPTS=$(dirname $0)

. $SCRIPTS/common.sh

DEPLOY_DIR=../apidumpsterfire-deploy
mkdir -p $DEPLOY_DIR
mkdir -p $DEPLOY_DIR/scripts
cp $SCRIPTS/deploy.sh.txt $DEPLOY_DIR/scripts/deploy-to-keroku.sh

go_generate

pushd $DEPLOY_DIR
scripts/deploy-to-keroku.sh
popd