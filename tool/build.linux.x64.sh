#!/bin/bash
go env -w GOOS=linux GOARCH=amd64
cd $GOPATH/src/github.com/schwarzlichtbezirk/pds
go build -o $GOPATH/bin/pds.client.x64 -v ./client
go build -o $GOPATH/bin/pds.server.x64 -v ./server
cp -r -u ./config $GOPATH/bin/pds-config
