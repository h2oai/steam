echo "Airlines data, predicting delay"
echo

echo "Destination JFK"
echo "http://localhost:55001/predict?Dest=JFK"
curl "http://localhost:55001/predict?Dest=JFK"
echo
echo

echo "Destination SFO"
echo "http://localhost:55001/predict?Dest=SFO"
curl "http://localhost:55001/predict?Dest=SFO"
echo
