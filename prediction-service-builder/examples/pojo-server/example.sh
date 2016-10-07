#!/usr/bin/env bash

rm -f example.war

curl -X POST \
--form pojo=@gbm_3f258f27_f0ad_4520_b6a5_3d2bb4a9b0ff.java \
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



