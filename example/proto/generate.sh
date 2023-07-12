#!/bin/bash

buf generate
rm -rf ../api
mv ./gen/github.com/pentops/protoc-gen-listify/example/api ../
rm -rf ./gen
