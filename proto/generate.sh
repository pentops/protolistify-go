#!/bin/bash

buf generate
rm -rf ../listify
mv ./gen/github.com/pentops/protoc-gen-listify/listify ../
cp ./static/v1/* ../listify/v1/
rm -rf ./gen
