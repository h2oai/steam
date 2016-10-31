#!/usr/bin/env bash

if ! [ -e ".running" ] ;
then
	echo "No service is currently running"
	exit 255
fi

kill `cat .running`
rm .running model.war h2o-genmodel.jar *.java

