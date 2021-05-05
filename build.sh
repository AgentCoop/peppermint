#!/usr/bin/env bash

PROJECT_ROOT=~/go/src/github.com/AgentCoop/peppermint
GO_IMPORT_PREFIX=github.com/AgentCoop/peppermint/internal/api
PROTO_ROOT=$PROJECT_ROOT/api
PROTO_BUILD=./build/proto
GEN_OUTPUT=$PROJECT_ROOT/internal/api/peppermint

mkdir -p $PROTO_BUILD
mkdir -p $GEN_OUTPUT

rm -rf $PROTO_BUILD/*
rm -rf $GEN_OUTPUT/*

cp -r $PROTO_ROOT/* $PROTO_BUILD/
find $PROTO_BUILD -name '*.proto' -exec sed -r -i "s@option go_package = \"(.*)\";@option go_package = \"$GO_IMPORT_PREFIX/\1\";@g" {} \;

for file in $(find $PROTO_BUILD -name '*.proto' -type f);
do
  echo Compiling ...$(basename $file)
  protoc -I="$PROTO_BUILD" \
    --go_opt=paths=source_relative \
    --go-grpc_opt=paths=source_relative \
    --go_out=$GEN_OUTPUT \
    --go-grpc_out=$GEN_OUTPUT $file
done

