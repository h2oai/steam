#!/usr/bin/env bash

rm -f example.war

curl -X POST \
--form mojo=@mojo.zip \
--form jar=@h2o-genmodel.jar \
--form deepwater=@deepwater-all.jar \
localhost:55000/makewar > example.war

if [ -s example.war ]
then
  echo "Created example.war"
  echo "Run with run-example.sh"
else
  echo "Failed to build example.war"
  exit 1
fi
