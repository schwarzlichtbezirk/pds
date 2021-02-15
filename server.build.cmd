@echo off
set GOARCH=amd64
go build -v -o %GOPATH%/bin/pds.server.exe github.com/schwarzlichtbezirk/pds-grpc/server