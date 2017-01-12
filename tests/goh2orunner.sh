#!/usr/bin/env bash

# Start H2O and create a start log + store pid
java -jar ../var/temp/h2o-3.10.2.1/h2o.jar --name steam-$(cat /dev/urandom | env LC_CTYPE=C tr -dc 'a-zA-Z0-9' | fold -w 32 | head -n 1) > start.log 2>&1 &
H2OPID=$!
echo "Started h2o with pid ${H2OPID}"

# Wait for cluster to start
sleep 10

# Search log for 'Cloud of size' to determine port of cluster
ADDRESS=$(sed -n '/Cloud of size/s/.* \[\/\([0-9]*\.[0-9]*\.[0-9]*\.[0-9]*:[0-9]*\)]/\1/p' start.log)


# Start model generating script
echo "Preparing models in cluster"
source bin/activate
python h2o-setup.py $ADDRESS

# Setup scoring service builder
java -jar ../var/master/assets/jetty-runner.jar --port 8181 ../var/master/assets/ROOT.war > compile.log 2>&1 &
SSBPID=$!
echo "Started scoring service builder with pid ${SSBPID}"

# Run GO tests here
go test ../master/web -h2o-address="${ADDRESS}" -compilation-service-address=":8181" -v -coverprofile=masterweb.cov
t1=$?

kill $H2OPID
kill $SSBPID

exit $t1
