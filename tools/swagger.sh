#!/bin/bash
wsdir=$(dirname $0)/..
protoc --proto_path=$wsdir\
 --openapiv2_out=$wsdir\
 --openapiv2_opt logtostderr=true\
 --openapiv2_opt allow_merge=true,merge_file_name=pds\
 $wsdir/pb/pds.proto
