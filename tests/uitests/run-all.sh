#!/bin/bash

WD=`pwd`
touch .failtmp
H2O_PATH=/Users/patrickrice/Documents/h2o-3/build/h2o.jar

rm -rf ./steam*-develop-darwin-amd64*
rm -rf ./steam*-master-darwin-amd64*
s3cmd get s3://steam-release/steamY-develop-darwin-amd64.tar.gz -f

tar xvf steamY-develop-darwin-amd64.tar.gz

java -jar $H2O_PATH -port 54535 -name steamtest > h2o.log 2>&1 &
H2O_PID=$!

sleep 2

python init_h2o.py

postgres -D /usr/local/var/postgres > postgres.log 2>&1 &
POSTGRES_PID=$!
echo > steam.log

java -jar steam-develop-darwin-amd64/var/master/assets/jetty-runner.jar \
	steam-develop-darwin-amd64/var/master/assets/ROOT.war > scoring-service.log 2>&1 &
JETTY_PID=$!

sleep 1

i=0
failcount=0

echo > $WD/.failures

for dir in `ls -d *-test/`; do
	cd steam-develop-darwin-amd64
	sleep 1
	echo "Resetting database"
	(cd var/master/scripts && ./reset-database.sh > /dev/null 2>&1)
	rm -rf var/master/model/*
	./steam serve master --superuser-name superuser --superuser-password superuser >> ../steam.log  2>&1 &
	STEAM_PID=$!
	cd ../$dir
	sleep 1
	python run.py > $WD/.testmp
	pass=$?
	if [ $pass -ne 0 ] 
		then
		cat $WD/.testmp >> $WD/.failures
		failcount=`expr $failcount + $pass`
		echo $dir >> $WD/.failtmp
		i=`expr $i + 1`
		echo "TEST FAILED"
	fi	
	kill -9 $STEAM_PID > /dev/null 2>&1
	cd ..
done


echo "$i test(s) failed"
cat $WD/.failures
rm $WD/.failtmp $WD/.failures $WD/.testmp



kill -9 $POSTGRES_PID > /dev/null 2>&1
kill -9 $H2O_PID > /dev/null 2>&1
kill -9 $JETTY_PID > /dev/null 2>&1

