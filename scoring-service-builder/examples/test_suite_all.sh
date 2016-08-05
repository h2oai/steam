#!/usr/bin/env bash


(cd pojo-server && sh test_suite.sh) || exit 1

(cd spam-detection-python && sh test_suite.sh) || exit 1

(cd spam-prejar  && sh test_suite.sh) || exit 1

