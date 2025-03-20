#!/usr/bin/env bash

CGO_ENABLED=0 go build -ldflags "-X 'colorlines/cl.Version=$version'"

ShowInfo "Build version"
./colorlines --version

ShowInfo "Zip"
tar -cavf colorlines-$GOOS-$GOARCH.tar.gz ./colorlines
