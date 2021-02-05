#!/usr/bin/env bash

mkdir -p build/proto

GITHUB_PKG=github.com/AgentCoop
PROTO_ROOT=~/repos/github/peppermint/api/strawman
PROTO_BUILD=./build/proto
GEN_OUTPUT=./internal/api

rm -rf $PROTO_BUILD/*
cp -r $PROTO_ROOT/* $PROTO_BUILD/
find $PROTO_BUILD -name '*.proto' -exec sed -r -i "s@option go_package = \"(.*)\";@option go_package = \"$GITHUB_PKG/\1\";@g" {} \;

for file in $(find $PROTO_BUILD -name '*.proto' -type f);
do
  echo Compiling ...$(basename $file)
  protoc -I="$PROTO_BUILD" \
    --go_opt=paths=source_relative \
    --go-grpc_opt=paths=source_relative \
    --go_out=$GEN_OUTPUT \
    --go-grpc_out=$GEN_OUTPUT $file
done

