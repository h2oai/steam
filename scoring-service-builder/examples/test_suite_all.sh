#!/usr/bin/env bash

echo "Building and testing pojo-server"
(cd pojo-server && bash example.sh) || exit 1
(cd pojo-server && sh test_suite.sh) || exit 1

echo "Building and testing spam-detection-python"
(cd spam-detection-python && bash example-python.sh) || exit 1
(cd spam-detection-python && bash test_suite.sh) || exit 1

echo "Building and testing spam-prejar"
(cd spam-prejar  && bash build.sh) || exit 1
(cd spam-prejar  && bash test_suite.sh) || exit 1
