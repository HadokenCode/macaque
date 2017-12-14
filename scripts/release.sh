#!/bin/bash
PACKAGE="github.com/wildnature/macaque"
DOCKER_REGISTRY="wildnature"
MODULE=$1
ENTRY=$2
DATE=$(date +%FT%T%z)
VERSION=$(cat .version)
echo "Go building. $ENTRY"
#go get google.golang.org/grpc
GOOS=linux GOARCH=amd64 go build \
    -tags release \
	-o build/macaque-$MODULE-$VERSION $ENTRY
echo "Go built completely."
docker build -f Dockerfile.$MODULE -t $DOCKER_REGISTRY/macaque-$MODULE:$VERSION . 
docker image tag $DOCKER_REGISTRY/macaque-$MODULE:$VERSION $DOCKER_REGISTRY/macaque-$MODULE:latest