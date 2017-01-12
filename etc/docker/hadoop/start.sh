#!/bin/bash

(nohup steam serve master --admin-name=root --admin-password=admin012 --web-address=:9002 & nohup java -jar var/master/assets/jetty-runner.jar var/master/assets/ROOT.war &) && /etc/bootstrap.sh -d
