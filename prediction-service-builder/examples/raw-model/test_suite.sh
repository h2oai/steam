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
"Get no model id" \
'curl -s "http://localhost:55001/predict?C2=1&C6=2"' \
'{"value":54.003414884892194}' \
$0:$LINENO

assert_equals_string \
"Get with model id" \
'curl -s "http://localhost:55001/predict/GBM_model_python_1473313897851_6?C2=1&C6=1"' \
'{"value":54.003414884892194}' \
$0:$LINENO

assert_equals_string \
"Post no model id" \
'curl -s -X POST --data \{"C2":"1","C6":"2"\} http://localhost:55001/predict' \
'{"value":54.003414884892194}' \
$0:$LINENO

assert_equals_string \
"Post no model id" \
'curl -s -X POST --data \{"C2":"1","C6":"2"\} http://localhost:55001/predict/GBM_model_python_1473313897851_6' \
'{"value":54.003414884892194}' \
$0:$LINENO

# kill process
echo "Stopping server"
kill $PID
sleep 2

post_tests $0



