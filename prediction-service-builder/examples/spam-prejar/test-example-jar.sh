
#for i in {1..100}
#do 


echo "This is an example of NOT SPAM"
text="Sorry  din lock my keypad"
echo \"$text\"
echo
#echo "http://localhost:55001/predict?$text"
#curl "http://localhost:55001/predict?$text"
#echo
#echo
echo "Post"
echo "curl -X Post --data \"$text\" http://localhost:55001/predict"
curl -X Post --data "$text" http://localhost:55001/predict
echo
echo

echo "This is an example of SPAM"
text="No 1 POLYPHONIC tone 4 ur mob every week! Just txt PT2 to 87575. 1st Tone FREE ! so get txtin now and tell ur friends. 150p/tone. 16 reply HL 4info"
echo "\"$text\""
echo
#echo "http://localhost:55001/predict?UR GOING 2 BAHAMAS! CallFREEFONE 08081560665 and speak to a live operator to claim either Bahamas cruise of2000 CASH 18+only. To opt out txt X to 07786200117"
#echo

echo
echo "Post"
echo "curl -X Post --data \"$text\" http://localhost:55001/predict"
curl -X Post --data "$text" http://localhost:55001/predict
echo

#done


