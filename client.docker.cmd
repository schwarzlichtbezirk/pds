@echo off
docker stop pds-grpc_client_1
docker rm pds-grpc_client_1
docker rmi pds-client
cd /d %GOPATH%/src/github.com/schwarzlichtbezirk/pds-grpc/client
docker build --rm -t pds-client .
