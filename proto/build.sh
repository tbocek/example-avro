#!/bin/bash

rm -rf pb
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
protoc-gen-go --version
protoc --go_out=. schema_v1.proto
protoc --go_out=. schema_v2.proto
go run protobuf.go