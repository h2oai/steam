echo "Airlines data, predicting delay"
echo

echo "Destination JFK"
echo "http://localhost:55001/predict?Dest=JFK"
curl -s "http://localhost:55001/predict?Dest=JFK"
echo
echo

echo "Destination SFO"
echo "http://localhost:55001/predict?Dest=SFO"
curl -s "http://localhost:55001/predict?Dest=SFO"
echo
echo

echo "Post"
echo 'curl -X POST --data '{"Dest":"JFK"}' http://localhost:55001/predict'
curl -s -X POST --data '{"Dest":"JFK"}' http://localhost:55001/predict
echo
echo

echo "Post"
echo 'curl -X POST --data '{"Dest":"SFO"}' http://localhost:55001/predict'
curl -s -X POST --data '{"Dest":"SFO"}' http://localhost:55001/predict
echo
echo

echo "Batch Post"
cat jsonlines.txt
echo 'curl -X POST --data-binary @jsonlines.txt  http://localhost:55001/predict'
curl -s -X POST --data-binary @jsonlines.txt  http://localhost:55001/predict
echo
echo




