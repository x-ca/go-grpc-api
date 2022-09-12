#!/bin/bash
# https://www.xiexianbin.cn/golang/net/grpc/index.html

SRC_DIR=..
DST_DIR=..
protoc -I=$SRC_DIR --go_opt=paths=source_relative --go_out=$DST_DIR $SRC_DIR/grpc/xca.proto

protoc -I=$SRC_DIR --go-grpc_out=$DST_DIR --go-grpc_opt=paths=source_relative $SRC_DIR/grpc/xca.proto
