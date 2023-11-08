#!/bin/bash

rm -rf api

buf -v generate proto
