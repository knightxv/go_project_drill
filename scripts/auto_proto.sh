#!/usr/bin/env bash

SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
source $SCRIPT_DIR/proto_dir.cfg

cd $SCRIPT_DIR/../pkg/proto
for ((i = 0; i < ${#all_proto[*]}; i++)); do
  proto=${all_proto[$i]}
  protoc -I ../../../  -I ./ --go_out=. --go-grpc_out=require_unimplemented_servers=false:. $proto
  echo "protoc --go_out=plugins=grpc:." $proto
done
echo "proto file generate success"





