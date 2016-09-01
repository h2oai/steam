#!/bin/bash

if [[ "$OSTYPE" == "linux"* ]]; then
	./.run-linux.sh
elif [[ "$OSTYPE" == "darwin"* ]]; then
	./.run-darwin.sh
fi

