@echo off
set pbsrc=%GOPATH%/src
set pbpkg=github.com/schwarzlichtbezirk/pds/pb
protoc -I=%pbsrc%/%pbpkg% --go_out=%pbsrc% --go-grpc_out=%pbsrc%^
 %pbsrc%/%pbpkg%/helloworld.proto
