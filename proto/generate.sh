#!/bin/bash

buf generate
rm -rf ../listify/v1/*.pb.go
mv ./gen/github.com/pentops/protoc-gen-listify/listify ../
rm -rf ./gen
