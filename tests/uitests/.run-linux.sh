#!/bin/bash

WD=`pwd`
touch .failtmp
H2O_PATH=~/documents/h2o/h2o.jar

rm -rf ./steam*-develop-linux-amd64*
rm -rf ./steam*-develop-linux-amd64*
curl -O http://s3.amazonaws.com/steam-release/steam-develop-linux-amd64.tar.gz 

tar xvf steam-develop-linux-amd64.tar.gz
cp steam-develop-linux-amd64/var/master/scripts/database/create-schema.sql steam-develop-linux-amd64/var/master/db
mv steam-develop-linux-amd64 steam

java -jar $H2O_PATH -port 54535 -name steamtest > h2o.log 2>&1 &
H2O_PID=$!
disown
sleep 5

python init_h2o.py

echo > steam.log


export JETTY_PATH=`pwd`/steam/var/master/assets/jetty-runner.jar
java -jar steam/var/master/assets/jetty-runner.jar \
	--port 55000 steam/var/master/assets/ROOT.war > scoring-service.log 2>&1 &
JETTY_PID=$!

sleep 1

i=0
failcount=0

cd steam
echo > $WD/.failures
./steam login localhost:9000 --username=superuser --password=superuser > /dev/null
./steam serve master --superuser-name superuser --superuser-password superuser --compilation-service-address=":55000" >> ../steam.log  2>&1 &
STEAM_PID=$!
disown
sleep 1
./steam register cluster --address="localhost:54535"
cd ..

for dir in `ls -d *-test`; do
	cd steam
	sleep 1
	#./steam serve master --superuser-name superuser --superuser-password superuser >> ../steam.log  2>&1 &
	cd ..
	cp testutil.py $dir/
	sleep 1
	TEST_FIREFOX=1 python $dir/run.py > $WD/.testmp
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

unset JETTY_PATH

echo "$i test(s) failed"
cat $WD/.failures
rm $WD/.failtmp $WD/.failures $WD/.testmp
rm -rf $WD/steam-develop-linux-amd64.tar.gz $WD/steam-develop-linux-amd64 $WD/steam




kill -9 $H2O_PID > /dev/null 2>&1
kill -9 $JETTY_PID > /dev/null 2>&1
kill -9 $STEAM_PID > /dev/null 2>&1
