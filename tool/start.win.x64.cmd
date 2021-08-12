@echo off
cd /d %GOPATH%/bin/
start "PDS server" pds.server.x64.exe
start "PDS client" pds.client.x64.exe
