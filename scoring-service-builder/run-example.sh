#!/usr/bin/env bash

echo "Starting prediction service on port 55001"
echo ""

java -jar jetty-runner-8.1.14.v20131031.jar --port 55001 example.war 

