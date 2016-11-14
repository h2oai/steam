#!/usr/bin/env bash

rm -f example-python.war

curl -X POST \
--form pojo=@GBM_model_python_1463864606917_1.java \
--form jar=@h2o-genmodel.jar \
--form python=@score.py \
--form pythonextra=@vectorizer.pickle \
--form pythonextra=@lib/modelling.py \
--form pythonextra=@lib/__init__.py \
--form envfile=@sp2.yaml \
localhost:55000/makepythonwar > example-python.war

if [ -s example-python.war ]
then
  echo "Created example-python.war"
  echo "Run with run-example-python.sh"
else
  echo "Failed to build example.war"
  exit 1
fi

