echo "This is an example of not spam"
echo "http://localhost:55001/pypredict?Sorry  din lock my keypad"
curl "http://localhost:55001/pypredict?Sorry%20%20din%20lock%20my%20keypad"
echo
echo

echo "This is an example of spam"
echo "You are a winner you have been specially selected to receive 1000 cash or a 2000 award. Speak to a live operator to claim call 087123002209am-7pm. Cost 10p"
curl "http://localhost:55001/pypredict?You%20are%20a%20winner%20you%20have%20been%20specially%20selected%20to%20receive%201000%20cash%20or%20a%202000%20award.%20Speak%20to%20a%20live%20operator%20to%20claim%20call%20087123002209am-7pm.%20Cost%2010p"
echo
