#!/usr/bin/env bash

source ./build_part_start.sh

export GOOS=darwin
export GOARCH=arm64

source ./build_part_end.sh
