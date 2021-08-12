@echo off
go env -w GOOS=windows GOARCH=386
cd /d %GOPATH%/bin/
go build -v -o pds.server.x86.exe github.com/schwarzlichtbezirk/pds/server
go build -v -o pds.client.x86.exe github.com/schwarzlichtbezirk/pds/client
xcopy %GOPATH%\src\github.com\schwarzlichtbezirk\pds\config pds-config /f /d /i /s /e /k /y
