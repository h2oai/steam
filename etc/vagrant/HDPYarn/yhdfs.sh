#!/usr/bin/env bash

# Allowing ssh with no passphrase for hadoop
ssh-keygen -t dsa -P '' -f ~/.ssh/id_dsa
cat ~/.ssh/id_dsa.pub >> ~/.ssh/authorized_keys
chmod 0600 ~/.ssh/authorized_keys
ssh-keyscan -t rsa localhost >> ~/.ssh/known_hosts
ssh-keyscan -t rsa 0.0.0.0 >> ~/.ssh/known_hosts

# Starting up hadoop
cd /usr/local/hadoop
bin/hdfs namenode -format -nonInteractive
sbin/start-dfs.sh
bin/hdfs dfs -mkdir /user
bin/hdfs dfs -mkdir /user/vagrant

# Starting Yarn
sbin/start-yarn.sh
