@echo off
go env -w GOOS=windows GOARCH=amd64
cd /d %GOPATH%\src\github.com\schwarzlichtbezirk\pds
go build -v -o %GOPATH%/bin/pds.server.x64.exe github.com/schwarzlichtbezirk/pds/server
go build -v -o %GOPATH%/bin/pds.client.x64.exe github.com/schwarzlichtbezirk/pds/client
xcopy .\config %GOPATH%\bin\pds-config /f /d /i /s /e /k /y
