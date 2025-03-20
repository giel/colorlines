#!/usr/bin/env bash

source ./build_part_start.sh

# go build
CGO_ENABLED=0 go build -ldflags "-X 'colorlines/cl.Version=$version'"

ShowInfo "Build version"
./colorlines --version
