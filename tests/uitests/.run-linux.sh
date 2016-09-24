#!/bin/bash

WD=`pwd`
touch .failtmp
H2O_PATH=~/Documents/h2o/h2o.jar

rm -rf ./steam*-master-linux-amd64*
rm -rf ./steam*-master-linux-amd64*
s3cmd get s3://steam-release/steamY-master-linux-amd64.tar.gz -f

tar xvf steamY-master-linux-amd64.tar.gz
cp steam-master-linux-amd64/var/master/scripts/database/create-schema.sql steam-master-linux-amd64/var/master/db


java -jar $H2O_PATH -port 54535 -name steamtest > h2o.log 2>&1 &
H2O_PID=$!
disown
sleep 5

python init_h2o.py

echo > steam.log

java -jar steam-master-linux-amd64/var/master/assets/jetty-runner.jar \
	steam-master-linux-amd64/var/master/assets/ROOT.war > scoring-service.log 2>&1 &
JETTY_PID=$!

sleep 1

i=0
failcount=0

echo > $WD/.failures

for dir in `ls -d *-test`; do
	cd steam-master-linux-amd64
	sleep 1
	echo "Resetting database"
	cd var/master/db
	rm steam.db
	sqlite3 steam.db < create-schema.sql
	cd ../../..
	rm -rf var/master/model/*
	./steam login localhost:9000 --username=superuser --password=superuser > /dev/null
	./steam serve master --superuser-name superuser --superuser-password superuser >> ../steam.log  2>&1 &
	STEAM_PID=$!
	disown
	cd ..
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
	kill -9 $STEAM_PID > /dev/null 2>&1
	rm $dir/testutil.py*
done


echo "$i test(s) failed"
cat $WD/.failures
rm $WD/.failtmp $WD/.failures $WD/.testmp
rm -rf $WD/steamY-master-linux-amd64.tar.gz $WD/steam-master-linux-amd64




kill -9 $H2O_PID > /dev/null 2>&1
kill -9 $JETTY_PID > /dev/null 2>&1

