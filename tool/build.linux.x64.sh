#!/bin/bash
cd $(dirname $0)/..
mkdir -pv $GOPATH/bin/pds-config
cp -ruv ./config/* $GOPATH/bin/pds-config
go env -w GOOS=linux GOARCH=amd64
go build -o $GOPATH/bin/pds.client.x64 -v ./client
go build -o $GOPATH/bin/pds.server.x64 -v ./server
