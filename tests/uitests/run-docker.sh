#!/bin/bash

curl -O http://s3.amazonaws.com/steam-release/steam-develop-linux-amd64.tar.gz
tar xvf steam-develop-linux-amd64.tar.gz
mv steam-develop-linux-amd64 steam
rm steam-develop-linux-amd64.tar.gz
cd steam
./steam serve master --scoring-service-address "localhost:55000" --superuser-name superuser --superuser-password superuser > /dev/null 2>&1 &
STEAM_PID=$!
disown
./steam login localhost:9000 --username=superuser --password=superuser
cd ..

curl -O http://download.h2o.ai/versions/h2o-3.10.0.7.zip
unzip h2o-3.10.0.7.zip
mv h2o-3.10.0.7 h2o
rm h2o-3.10.0.7.zip
cd h2o
java -jar h2o.jar --port 54535 -name steamtest > h2o.log 2>&1 &
H2O_PID=$!
disown
sleep 5
cd ..
python init_h2o.py

java -jar steam/var/master/assets/jetty-runner.jar \
	--port 55000 steam/var/master/assets/ROOT.war &
JETTY_PID=$!
disown
sleep 2
./steam/steam register cluster --address='localhost:54535'

WD=`pwd`

i=1
rm -f .aproc
for dir in `ls -d *-test`; do
	sudo docker run --net=host -v $WD:/root/steam -v /tmp/.X11-unix:/tmp/.X11-unix -e DISPLAY=$DISPLAY steamuser $i &
	echo $! >> .aproc
	i=`expr $i + 1`
	sleep 0.5
done

for PID in `cat .aproc`; do
	wait $PID
done
rm -f .aproc

rm -rf h2o steam

kill -9 $STEAM_PID
kill -9 $H2O_PID
kill -9 $JETTY_PID

