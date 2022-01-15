#!/bin/bash
wsdir=$(dirname $0)/..
protoc --proto_path=$wsdir\
 --go_out=$wsdir --go_opt paths=source_relative\
 --go-grpc_out=$wsdir --go-grpc_opt paths=source_relative\
 --grpc-gateway_out=$wsdir\
 --grpc-gateway_opt logtostderr=true\
 --grpc-gateway_opt paths=source_relative\
 --grpc-gateway_opt generate_unbound_methods=true\
 $wsdir/pb/pds.proto
