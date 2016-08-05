#!/usr/bin/env bash

javac -cp 'lib/*' Transform.java PreProcess.java Tfidf.java || exit 1

jar cf pre.jar Transform.class PreProcess.class  Tfidf.class || exit 1

rm -r example.war

curl -X POST \
--form pojo=@gbm_cf6fdeef_cad1_4e85_b644_6358166076ca.java \
--form jar=@lib/h2o-genmodel.jar \
--form prejar=@pre.jar \
--form preclass=PreProcess \
localhost:55000/makewar > example.war

if [ -s example.war ]
then
  echo "Created example.war"
  echo "Run with run-example.sh"
else
  echo "Failed to build example.war"
  exit 1
fi

