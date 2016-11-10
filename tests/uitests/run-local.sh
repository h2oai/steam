#!/bin/bash


cd $GOPATH/src/github.com/h2oai/steam/
STEAM_DIR=$GOPATH/src/github.com/h2oai/steam
TESTS_DIR=$STEAM_DIR/tests/uitests
make clean
make all db

curl -o h2o.zip http://h2o-release.s3.amazonaws.com/h2o/rel-turing/8/h2o-3.10.0.8.zip
unzip h2o.zip
H2O_DIR=$STEAM_DIR/h2o-3.10.0.8
rm h2o.zip

./steam serve master --superuser-name=superuser --superuser-password=superuser --compilation-service-address=":55000" > \
	$TESTS_DIR/steam.log 2>&1 &

STEAM_PID=$!
disown

java -jar $H2O_DIR/h2o.jar --port 54535 --name steamtest > $TESTS_DIR/h2o.log 2>&1 &
H2O_PID=$!
disown

export JETTY_PATH=`pwd`/var/master/assets/jetty-runner.jar
java -jar var/master/assets/jetty-runner.jar --port 55000 \
	var/master/assets/ROOT.war > $TESTS_DIR/scoring-service.log 2>&1 &
JETTY_PID=$!
disown

sleep 5
$STEAM_DIR/steam login localhost:9000 --username=superuser --password=superuser > /dev/null
$STEAM_DIR/steam register cluster --address="localhost:54535"

cd $TESTS_DIR

python init_h2o.py


echo > .failures

export STEAM_PATH=$STEAM_DIR/steam

for dir in `ls -d *-test`; do
	cp testutil.py $dir/
	sleep 1
	TEST_FIREFOX=1 python $dir/run.py > .testmp
	pass=$?
	if [ $pass -ne 0 ] 
		then
		echo $dir >> .failures
		cat .testmp >> .failures
		failcount=`expr $failcount + $pass`
		echo -e "\n\n" >> .failures
		i=`expr $i + 1`
		echo "TEST FAILED"
	fi	
	rm $dir/testutil.py*	
done

unset JETTY_PATH

kill $STEAM_PID
kill $H2O_PID
kill $JETTY_PID

cat .failures
rm .failures

rm -rf $H2O_DIR


