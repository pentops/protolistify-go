#!/bin/bash

rm -rf listify

buf -v generate proto

cp proto/static/v1/* listify/v1/
