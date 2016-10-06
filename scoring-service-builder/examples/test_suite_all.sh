#!/usr/bin/env bash

source bash_unit_tests.sh

(cd .. && ./gradlew build) || exit 1

# start server
echo "Starting building server"

nc -z localhost 55000 && echo "Something already running on port 55000" && exit 1


java -jar ../jetty-runner-8.1.14.v20131031.jar --port 55000 ../build/libs/ROOT.war &
BUILD_PID=$!
sleep 2 

function killproc {
    kill $BUILD_PID
    echo "Exiting"
    echo
    echo
    echo
    echo
    exit 1
}

#trap killproc EXIT
trap killproc 1

pre_tests $0

echo "Building and testing pojo-server prediction service"
(cd pojo-server && bash example.sh) || exit 1
(cd pojo-server && bash test_suite.sh) || exit 1

echo "Building and testing rawmodel prediction service"
(cd raw-model && bash example.sh) || exit 1
(cd raw-model && bash test_suite.sh) || exit 1

echo "Building and testing spam-detection-python prediction service"
(cd spam-detection-python && bash example-python.sh) || exit 1
(cd spam-detection-python && bash test_suite.sh) || exit 1

echo "Building and testing spam-prejar prediction service"
(cd spam-prejar  && bash build.sh) || exit 1
(cd spam-prejar  && bash test_suite.sh) || exit 1

echo "Building and testing compile to jar"
(cd compile  && bash example.sh) || exit 1
(cd compile  && bash test_suite.sh) || exit 1

# kill process
echo "Stopping building server"
kill $BUILD_PID
sleep 2

post_tests $0

