@echo off
docker stop pds-grpc_server_1
docker rm pds-grpc_server_1
docker rmi pds-server
cd /d %GOPATH%/src/github.com/schwarzlichtbezirk/pds-grpc/server
docker build --rm -t pds-server .
