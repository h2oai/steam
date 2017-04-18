#!/bin/bash

cd /home/patrick

rm -f steamY-develop-linux-amd64.tar.gz
rm -rf steam-develop-linux-amd64

curl -O https://s3.amazonaws.com/steam-release/steamY-develop-linux-amd64.tar.gz
tar xvf steamY-develop-linux-amd64.tar.gz
cd steam-develop-linux-amd64
./steam serve master --admin-name patrick --admin-password admin012

