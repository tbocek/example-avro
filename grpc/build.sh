#!/bin/bash

rm -rf pb
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
protoc-gen-go --version
protoc --go_out=. --go-grpc_out=. schema_grpc.proto
