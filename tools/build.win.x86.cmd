@echo off
cd /d %~dp0..
xcopy .\config %GOPATH%\bin\pds-config /f /d /i /e /k /y
go env -w GOOS=windows GOARCH=386
go build -v -o %GOPATH%/bin/pds.server.x86.exe ./server
go build -v -o %GOPATH%/bin/pds.client.x86.exe ./client
