#!/usr/bin/env bash

# Start H2O and create a start log + store pid
java -jar ../../var/temp/h2o-$H2OVERSION/h2o.jar > start.log 2>&1 &
PID=$!
echo "Started h2o with pid ${PID}"

# Wait for cluster to start
sleep 5

# Search log for 'Cloud of size' to determine port of cluster
ADDRESS=$(sed -n '/Cloud of size/s/.* \[\/\([0-9]*\.[0-9]*\.[0-9]*\.[0-9]*:[0-9]*\)]/\1/p' start.log)


# Start model generating script
echo "Preparing models in cluster"
python h2o-setup.py $ADDRESS

# Run GO tests here
go test ../../master/data -v -coverprofile=masterdata.cov
go test ../../master/web -v -coverprofile=masterweb.cov

kill -9 $PID
