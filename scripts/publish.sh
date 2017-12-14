#!/bin/bash
DOCKER_REGISTRY="wildnature"
MODULE=$1
VERSION=$(cat .version)

docker login -u "$(DOCKER_USER)" -p "$(DOCKER_PASS)"
docker push $DOCKER_REGISTRY/macaque-$MODULE:$VERSION