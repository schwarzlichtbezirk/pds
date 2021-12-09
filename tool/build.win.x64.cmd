@echo off
cd /d %GOPATH%\src\github.com\schwarzlichtbezirk\pds
xcopy .\config %GOPATH%\bin\pds-config /f /d /i /e /k /y
go env -w GOOS=windows GOARCH=amd64
go build -v -o %GOPATH%/bin/pds.server.x64.exe ./server
go build -v -o %GOPATH%/bin/pds.client.x64.exe ./client
