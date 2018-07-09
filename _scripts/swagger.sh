#!/usr/bin/env bash
# set -e

BASE_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )/.." && pwd )"

cd $BASE_DIR

# Setup
# git clone git@github.com:swagger-api/swagger-ui.git ./tmp/

# cd ./tmp 
# git checkout v3.17.2

# rm -Rf ./static/swagger-ui

# mv ./tmp/dist ./static/swagger-ui

# rm -Rf ./tmp


cp protos/*swagger.json ./static/swagger-ui/.

go-bindata-assetfs -o "swagger/data.go" -pkg swagger -ignore=\\.sh -ignore=\\.go ./static/swagger-ui/...