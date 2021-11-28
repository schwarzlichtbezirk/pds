@echo off
go env -w GOOS=windows GOARCH=386
cd /d %GOPATH%\src\github.com\schwarzlichtbezirk\pds
go build -v -o %GOPATH%/bin/pds.server.x86.exe ./server
go build -v -o %GOPATH%/bin/pds.client.x86.exe ./client
xcopy .\config %GOPATH%\bin\pds-config /f /d /i /s /e /k /y
