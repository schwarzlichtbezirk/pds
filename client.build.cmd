@echo off
set GOARCH=amd64
go build -v -o %GOPATH%/bin/pds.client.exe github.com/schwarzlichtbezirk/pds-grpc/client