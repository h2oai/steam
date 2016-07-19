#!/usr/bin/env bash

go build
mkdir -p ../../docs/cli
rm -f ../../docs/cli/*.md
./cli-md ../../docs/cli/
mv ../../docs/cli/steam.md ../../docs/cli/README.md
