#!/usr/bin/env bash

go build
mkdir -p ../../doc/cli
rm -f ../../doc/cli/*.md
./cli-md ../../doc/cli/
mv ../../doc/cli/steam.md ../../doc/cli/README.md
