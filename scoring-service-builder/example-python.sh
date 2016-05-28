curl -X POST --form pojo=@examples/example-spam-detection/GBM_model_python_1463864606917_1.java --form jar=@examples/example-spam-detection/h2o-genmodel.jar --form python=@examples/example-spam-detection/score.py  localhost:55000/makepythonwar > example-python.war

echo "Created example-python.war"
echo "Run with run-example-pyhton.sh"
