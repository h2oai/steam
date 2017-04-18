#!/bin/bash


cd $GOPATH/src/github.com/h2oai/steam/
STEAM_DIR=$GOPATH/src/github.com/h2oai/steam
TESTS_DIR=$STEAM_DIR/tests/uitests
make clean all
sudo -v
rm -f var/master/db/steam.db

H2O_DIR=~/documents/h2o-3/build


export H2O_PATH=$H2O_DIR/h2o.jar
export STEAM_PATH=$STEAM_DIR/steam

sudo ./steam serve master --admin-name=patrick --admin-password=superuser --compilation-service-address=":55000" > \
	$TESTS_DIR/steam.log 2>&1 &

STEAM_PID=$!
disown

java -jar $H2O_DIR/h2o.jar --port 54535 --name steamtest > $TESTS_DIR/h2o.log 2>&1 &
H2O_PID=$!
disown
java -jar $H2O_DIR/h2o.jar --port 54321 --name pjr > /dev/null 2>&1 &
H2O2_PID=$!
disown


export JETTY_PATH=`pwd`/var/master/assets/jetty-runner.jar
java -jar var/master/assets/jetty-runner.jar --port 55000 \
	var/master/assets/ROOT.war > $TESTS_DIR/scoring-service.log 2>&1 &
JETTY_PID=$!
disown

sleep 5
$STEAM_DIR/steam login localhost:9000 --username=patrick --password=superuser > /dev/null
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
unset H2O_PATH
unset STEAM_PATH

sudo rm -r $STEAM_DIR/var
sudo killall steam
kill $H2O_PID
kill $H2O2_PID
kill $JETTY_PID

cat .failures
rm .failures

