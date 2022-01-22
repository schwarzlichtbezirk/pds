#!/bin/bash
cd $(dirname $0)/..
git describe --tags > buildvers.txt # puts version to file for docker image building
buildvers=`cat buildvers.txt`
builddate=$(date +%F)
mkdir -pv $GOPATH/bin/config
cp -ruv ./config/* $GOPATH/bin/config
go env -w GOOS=linux GOARCH=amd64
go build -o $GOPATH/bin/pds.client.x64 -v -ldflags="-X 'main.buildvers=$buildvers' -X 'main.builddate=$builddate'" ./client
go build -o $GOPATH/bin/pds.server.x64 -v -ldflags="-X 'main.buildvers=$buildvers' -X 'main.builddate=$builddate'" ./server
