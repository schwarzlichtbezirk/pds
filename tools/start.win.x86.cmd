@echo off
start "server" %GOPATH%/bin/pds.server.x86.exe
start "client" %GOPATH%/bin/pds.client.x86.exe
