#!/usr/bin/env bash

source ../bash_unit_tests.sh

pre_tests $0

assert_equals_string \
"Command line JFK" \
'java -jar example.jar example.jar gbm_3f258f27_f0ad_4520_b6a5_3d2bb4a9b0ff "{Dest:JFK}"' \
'{"labelIndex":1,"label":"Y","classProbabilities":[0.026513747179178093,0.9734862528208219]}' \
$0:$LINENO

assert_equals_string \
"Command line SFO" \
'java -jar example.jar example.jar gbm_3f258f27_f0ad_4520_b6a5_3d2bb4a9b0ff "{Dest:SFO}"' \
'{"labelIndex":1,"label":"Y","classProbabilities":[0.008905417583984554,0.9910945824160154]}' \
$0:$LINENO

post_tests $0
