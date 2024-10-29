@echo off
start "PDS server" %GOPATH%/bin/pds.server.x64.exe
start "PDS client" %GOPATH%/bin/pds.client.x64.exe
