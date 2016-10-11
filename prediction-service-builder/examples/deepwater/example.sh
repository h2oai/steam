#!/usr/bin/env bash

rm -f example.war

curl -X POST \
--form mojo=@DeepWater_model_python_1476140406268_1.zip \
--form jar=@h2o-genmodel.jar \
localhost:55000/makewar > example.war

if [ -s example.war ]
then
  echo "Created example.war"
  echo "Run with run-example.sh"
else
  echo "Failed to build example.war"
  exit 1
fi
