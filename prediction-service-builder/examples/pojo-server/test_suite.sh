#!/usr/bin/env bash

source ../bash_unit_tests.sh

pre_tests $0

# start server
nc -z localhost 55001 && echo "Something already running on port 55001" && exit 1

echo "Starting server"
java -jar ../jetty-runner-8.1.14.v20131031.jar --port 55001 example.war & 
PID=$!
sleep 2 

assert_equals_string \
"GET JFK" \
'curl -s "http://localhost:55001/predict?Dest=JFK"' \
'{"labelIndex":1,"label":"Y","classProbabilities":[0.026513747179178093,0.9734862528208219]}' \
$0:$LINENO

assert_equals_string \
"GET SFO" \
'curl -s "http://localhost:55001/predict?Dest=SFO"' \
'{"labelIndex":1,"label":"Y","classProbabilities":[0.008905417583984554,0.9910945824160154]}' \
$0:$LINENO


assert_equals_string \
"POST JFK" \
'curl -s -X POST --data {"Dest":"JFK"} http://localhost:55001/predict' \
'{"labelIndex":1,"label":"Y","classProbabilities":[0.026513747179178093,0.9734862528208219]}' \
$0:$LINENO

assert_equals_string \
"POST SFO" \
'curl -s -X POST --data {"Dest":"SFO"} http://localhost:55001/predict' \
'{"labelIndex":1,"label":"Y","classProbabilities":[0.008905417583984554,0.9910945824160154]}' \
$0:$LINENO

# for now removing ending newlines in batch result
assert_equals_string \
"Batch POST" \
'curl -s -X POST --data-binary @jsonlines.txt  http://localhost:55001/predict | tr -d "\n"' \
'{"labelIndex":1,"label":"Y","classProbabilities":[0.026513747179178093,0.9734862528208219]}{"labelIndex":1,"label":"Y","classProbabilities":[0.008905417583984554,0.9910945824160154]}{"labelIndex":1,"label":"Y","classProbabilities":[0.026513747179178093,0.9734862528208219]}{"labelIndex":1,"label":"Y","classProbabilities":[0.008905417583984554,0.9910945824160154]}' \
$0:$LINENO


# kill process
echo "Stopping server"
kill $PID
sleep 2

post_tests $0



