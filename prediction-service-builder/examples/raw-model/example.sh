#!/usr/bin/env bash

rm -f example.war

curl -X POST \
--form mojo=@GBM_model_python_1473313897851_6.zip \
--form jar=@h2o-genmodel-all.jar \
localhost:55000/makewar > example.war

if [ -s example.war ]
then
  echo "Created example.war"
  echo "Run with run-example.sh"
else
  echo "Failed to build example.war"
  exit 1
fi
