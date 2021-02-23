@echo off
set GOARCH=amd64
go build -v -o %GOPATH%/bin/pds.server.exe github.com/schwarzlichtbezirk/pds-grpc/server
go build -v -o %GOPATH%/bin/pds.client.exe github.com/schwarzlichtbezirk/pds-grpc/client
xcopy %GOPATH%\src\github.com\schwarzlichtbezirk\pds-grpc\config %GOPATH%\bin\pds-config /f /d /i /s /e /k /y