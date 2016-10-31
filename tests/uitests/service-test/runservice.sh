#!/usr/bin/env bash

if [ -z "$JETTY_PATH" ] ;
then
	echo "JETTY_PATH is not set"
	exit 255
fi

if [ -e ".running" ] ;
then
	echo "A service is already running"
	exit 255
fi


curl -s http://localhost:54535/3/Models.java/$1 > $1.java
curl -s http://localhost:54535/3/h2o-genmodel.jar > h2o-genmodel.jar

curl -s -X POST --form pojo=@$1.java \
	--form jar=@h2o-genmodel.jar \
	localhost:55000/makewar > model.war


java -jar $JETTY_PATH --port 55001 model.war > /dev/null 2>&1 &
PID=$!

echo $PID > .running


