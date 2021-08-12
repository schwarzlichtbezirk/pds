@echo off
cd /d %GOPATH%/bin/
start "server" pds.server.x86.exe
start "client" pds.client.x86.exe
