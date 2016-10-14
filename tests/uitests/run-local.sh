#!/bin/bash


cd $GOPATH/src/github.com/h2oai/steam/
STEAM_PATH=$GOPATH/src/github.com/h2oai/steam
TESTS_PATH=$STEAM_PATH/tests/uitests
make clean
make all db

curl -o h2o.zip http://h2o-release.s3.amazonaws.com/h2o/rel-turing/8/h2o-3.10.0.8.zip
unzip h2o.zip
H2O_PATH=$STEAM_PATH/h2o-3.10.0.8

./steam serve master --superuser-name=superuser --superuser-password=superuser --compilation-service-address=":55000" > \
	$TEST_PATH/steam.log 2>&1 &

STEAM_PID=$!
disown

java -jar $H2O_PATH/h2o.jar --port 54535 > $TESTS_PATH/h2o.log 2>&1 &
H2O_PID=$!
disown

java -jar var/master/assets/jetty-runner.jar --port 55000 \
	var/master/assets/ROOT.war > $TESTS_PATH/scoring-service.log 2>&1 &
JETTY_PID=$!
disown

$STEAM_PATH/steam login localhost:9000 --name=superuser --password=superuser > /dev/null
$STEAM_PATH/steam register cluster --address="localhost:54535"

cd $TESTS_PATH

sleep 5
python init_h2o.py


echo > .failures

for dir in `ls -d *-test`; do
	cp testutil.py $dir/
	sleep 1
	python $dir/run.py > $WD/.testmp
	pass=$?
	if [ $pass -ne 0 ] 
		then
		echo $dir >> $WD/.failures
		cat $WD/.testmp >> $WD/.failures
		failcount=`expr $failcount + $pass`
		echo -e "\n\n" >> $WD/.failures
		i=`expr $i + 1`
		echo "TEST FAILED"
	fi	
	rm $dir/testutil.py*	
done

kill $STEAM_PID
kill $H2O_PID
kill $JETTY_PID

