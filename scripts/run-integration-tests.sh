#!/bin/bash
MODULE=$1
PORT=$2
ENTRY=$3
MAX_TRIES=5
VERSION=$(cat .version)
GOOS=linux GOARCH=amd64 go build \
    -tags release \
	-o build/macaque-$MODULE-$VERSION $ENTRY

docker-compose -f testInt/$MODULE/res/docker-compose.yml pull 
echo "Building containers"
docker-compose -f testInt/$MODULE/res/docker-compose.yml build
echo "Gonna up the compose"
docker-compose -f testInt/$MODULE/res/docker-compose.yml up -d

#echo "Gonna check the port ${PORT}"
#while ! nc localhost $PORT; do 
#    if [ $MAX_TRIES -eq 0 ]
#    then
#        docker-compose -f testInt/$MODULE/docker-compose.yml ps
#        docker-compose -f testInt/$MODULE/docker-compose.yml logs
#        docker-compose -f testInt/$MODULE/docker-compose.yml stop
#        docker-compose -f testInt/$MODULE/docker-compose.yml rm -f
#        exit 1
#    fi
#    echo "Checking port $PORT"
#    let MAX_TRIES-=1
#    sleep 5; 
#done;
echo 'The docker-compose is already up!'
echo 'Running integration tests'
go test -v github.com/wildnature/macaque/testInt/$MODULE/src 
echo 'Integration tests were already run.'
docker-compose -f testInt/$MODULE/res/docker-compose.yml ps
docker-compose -f testInt/$MODULE/res/docker-compose.yml logs
docker-compose -f testInt/$MODULE/res/docker-compose.yml stop
docker-compose -f testInt/$MODULE/res/docker-compose.yml rm -f
