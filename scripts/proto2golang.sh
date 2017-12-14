#!/bin/bash

go get -u github.com/golang/protobuf/protoc-gen-go

cd protobuf

for r in * ; do
  echo "Generating code from protos in $r into ../pkg/pb/$r"
  if [[ ! -e ../pkg/pb/$r ]]; then
    mkdir ../pkg/pb/$r
  fi
  rm -rf ../pkg/pb/$r/*.proto
  protoc -I.\
  --go_out=plugins=grpc:../pkg/pb ./$r/*.proto
done
