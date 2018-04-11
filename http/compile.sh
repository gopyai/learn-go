#!/bin/bash

docker run \
--rm \
-v $GOPATH/src:/go/src \
-w /go/src/learn/http \
golang go build -o http-linux

scp -i $it3s http-linux ubuntu@gw-a:.
scp -i $it3s http-linux ubuntu@gw-b:.