#!/usr/bin/env bash

# Start H2O and create a start log + store pid
java -jar ../var/temp/h2o-$H2OVERSION/h2o.jar > start.log 2>&1 &
H2OPID=$!
echo "Started h2o with pid ${H2OPID}"

# Wait for cluster to start
sleep 5

# Search log for 'Cloud of size' to determine port of cluster
ADDRESS=$(sed -n '/Cloud of size/s/.* \[\/\([0-9]*\.[0-9]*\.[0-9]*\.[0-9]*:[0-9]*\)]/\1/p' start.log)


# Start model generating script
echo "Preparing models in cluster"
python h2o-setup.py $ADDRESS

# Setup scoring service builder
java -jar ../var/master/assets/jetty-runner.jar ../var/master/assets/ROOT.war > compile.log 2>&1 &
SSBPID=$?
echo "Started scoring service builder with pid ${SSBPID}"

# Run GO tests here
go test ../master/web --working-directory="../../" --cluster-address="${ADDRESS}" -v -coverprofile=masterweb.cov
t1=$?

kill -9 $H2OPID
kill -9 $SSBPID

exit $t1
