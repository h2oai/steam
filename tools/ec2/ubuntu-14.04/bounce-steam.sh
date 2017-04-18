#!/usr/bin/env bash

set -x

DIR=/home/ubuntu

# Kill server
pkill steam
pkill jetty-runner

set -e

cd $DIR
# Backup log
mkdir -p $DIR/steam-logs
if [ -f $DIR/steam/steam.log ]; then
        mv $DIR/steam/steam.log "$DIR/steam-logs/steam-$(date +%Y%m%d-%H%M%S).log"
fi

# Delete application directory
if [ -d $DIR/steam ]; then
        rm -rf $DIR/steam
fi

# Fetch latest package
cd $DIR
s3cmd get s3://steam-release/steam-master-linux-amd64.tar.gz
tar xvf steam-master-linux-amd64.tar.gz
mv steam-master-linux-amd64 steam
rm steam-master-linux-amd64.tar.gz

# Reset database
cd $DIR/steam/var/master/scripts && sudo -u postgres ./drop-database.sh
cd $DIR/steam/var/master/scripts && sudo -u postgres ./create-database.sh

# Start server
cd $DIR/steam
nohup java -jar \
		./var/master/assets/jetty-runner.jar \
		./var/master/assets/ROOT.war \
		>> comp.log 2>&1 &

nohup ./steam serve master \
        --admin-name=steamer \
        --admin-password=terrabella \
        --web-tls-cert-path=/etc/ssl/star_h2o_ai.pem \
        --web-tls-key-path=/etc/ssl/star_h2o_ai.key \
        >> steam.log 2>&1 &

