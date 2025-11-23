#!/bin/bash

rm -rf gen-go
thrift -r --gen go schema.thrift