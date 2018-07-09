#!/usr/bin/env bash

set -e

CUR_DIR=$(pwd)

# this installs globally 
# go install ./vendor/github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
# go install ./vendor/github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
cd $GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway && go install
cd $GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger && go install

cd $CUR_DIR

ORG_PATH="github.com/stevepartridge/"

# roll that goodness dough
protoc -I=./ -I=$GOPATH/src/$ORG_PATH -I=$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --go_out=plugins=grpc:. ./protos/*.proto
protoc -I=./ -I=$GOPATH/src/$ORG_PATH -I=$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --grpc-gateway_out=logtostderr=true:. ./protos/*.proto
protoc -I=./ -I=$GOPATH/src/$ORG_PATH -I=$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --proto_path=. --swagger_out=logtostderr=true:. ./protos/*.proto

