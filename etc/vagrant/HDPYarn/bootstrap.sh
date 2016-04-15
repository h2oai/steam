#!/usr/bin/env bash

# Installing Java and JDK
apt-get update
apt-get install -y default-jre
apt-get install -y default-jdk

echo 'JAVA_HOME="/usr"' >> /etc/environment
