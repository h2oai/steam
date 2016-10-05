echo "Some random models"
echo

curl  "http://localhost:55001/predict?C2=1&C6=2"
echo

curl  "http://localhost:55001/predict/GBM_model_python_1473313897851_6?C2=1&C6=1"
echo

curl  -X POST --data '{"C2":"1", "C6":"2"}' http://localhost:55001/predict
echo

curl  -X POST --data '{"C2":"1", "C6":"2"}' http://localhost:55001/predict/GBM_model_python_1473313897851_6
echo



