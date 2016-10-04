#!/usr/bin/env bash

source ../bash_unit_tests.sh

pre_tests $0

# start server
echo "Starting server"
nc -z localhost 55001 && echo "Something already running on port 55001" && exit 1

java -jar ../jetty-runner-8.1.14.v20131031.jar --port 55001 example.war & 
PID=$!
sleep 2 

assert_equals_string \
"GET ham" \
'curl -s "http://localhost:55001/predict?Sorry%20%20din%20lock%20my%20keypad"' \
'{"labelIndex":0,"label":"ham","classProbabilities":[0.7579880962756615,0.24201190372433845]}' \
$0:$LINENO

assert_equals_string \
"GET spam" \
'curl -s "http://localhost:55001/predict?You%20are%20a%20winner%20you%20have%20been%20specially%20selected%20to%20receive%201000%20cash%20or%20a%202000%20award.%20Speak%20to%20a%20live%20operator%20to%20claim%20call%20087123002209am-7pm.%20Cost%2010p"' \
'{"labelIndex":0,"label":"ham","classProbabilities":[0.7579880962756615,0.24201190372433845]}' \
$0:$LINENO

assert_equals_string \
"POST ham" \
'curl -s -X Post --data "Sorry  din lock my keypad" http://localhost:55001/predict' \
'{"labelIndex":0,"label":"ham","classProbabilities":[0.7579880962756615,0.24201190372433845]}' \
$0:$LINENO

assert_equals_string \
"POST spam" \
'curl -s -X Post --data "You are a winner you have been specially selected to receive 1000 cash or a 2000 award. Speak to a live operator \
to claim call 087123002209am-7pm. Cost 10p" http://localhost:55001/predict' \
'{"labelIndex":0,"label":"ham","classProbabilities":[0.7579880962756615,0.24201190372433845]}' \
$0:$LINENO


echo "Stopping server"
kill $PID
sleep 2

post_tests $0



