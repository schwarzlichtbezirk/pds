@echo off
cd /d %GOPATH%/bin/
set pbimport=github.com/schwarzlichtbezirk/pds/pb
protoc -I=%GOPATH%/src/%pbimport%/^
 --go_out=%GOPATH%/src/ --go-grpc_out=%GOPATH%/src/^
 %GOPATH%/src/%pbimport%/pds.proto
