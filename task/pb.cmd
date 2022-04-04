@echo off

set wsdir=%~dp0..
set pbdir=%wsdir%/pb
set apidir=%wsdir%/api/export

protoc -I=%apidir% -I=%wsdir%/api/import^
 --go_out %pbdir%^
 --go_opt paths=source_relative^
 --go-grpc_out %pbdir%^
 --go-grpc_opt paths=source_relative^
 --grpc-gateway_out %pbdir%^
 --grpc-gateway_opt logtostderr=true^
 --grpc-gateway_opt paths=source_relative^
 --grpc-gateway_opt generate_unbound_methods=true^
 pds.proto
