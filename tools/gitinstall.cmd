@echo off
if not defined GOPATH (
	echo Golang https://go.dev/dl/ should been installed
	goto:eof
)
if not exist %GOPATH%\bin\protoc.exe (
	echo Install protocol buffers compiler https://github.com/protocolbuffers/protobuf/releases
	goto:eof
)

echo.
echo STAGE#1: install proto plugins
go install -v^
 google.golang.org/protobuf/cmd/protoc-gen-go@latest^
 google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest^
 github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest^
 github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest
if %errorlevel% neq 0 (
	goto:eof
)

echo.
echo STAGE#2: download and installs dependencies
cd %~dp0..
go mod download
if %errorlevel% neq 0 (
	goto:eof
)

echo.
echo STAGE#3: compile proto-files
call tools/pb.cmd
rem call tools/swagger.cmd
if %errorlevel% neq 0 (
	goto:eof
)

echo.
echo STAGE#4: compile executable binaries
if "%PROCESSOR_ARCHITECTURE%" equ "amd64" (
	call tools/build.win.x64.cmd
	goto endarch
)
if "%PROCESSOR_ARCHITECTURE%" equ "AMD64" (
	call tools/build.win.x64.cmd
	goto endarch
)
if "%PROCESSOR_ARCHITECTURE%" equ "x86" (
	call tools/build.win.x86.cmd
	goto endarch
)
if "%PROCESSOR_ARCHITECTURE%" equ "X86" (
	call tools/build.win.x86.cmd
	goto endarch
)
echo processor architecture %PROCESSOR_ARCHITECTURE% does not supported
:endarch

echo.
echo STAGE#5: build docker images
docker build --pull --rm -f "Dockerfile.server" -t pds-server:latest "."
docker build --pull --rm -f "Dockerfile.client" -t pds-client:latest "."
