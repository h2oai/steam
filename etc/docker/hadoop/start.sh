#!/bin/bash

(nohup steam serve master --superuser-name=root --superuser-password=superuser --web-address=:9002 & nohup java -jar var/master/assets/jetty-runner.jar var/master/assets/ROOT.war &) && /etc/bootstrap.sh -d
