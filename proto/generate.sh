#!/bin/bash

buf generate
rm -rf ../listify
mv ./gen/github.com/pentops/protoc-gen-listify/listify ../
rm -rf ./gen
