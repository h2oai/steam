#!/usr/bin/env bash

#
# ubuntu/trusty64 provisioning for Steam
# 

# Update
sudo apt-get update
sudo apt-get upgrade -y

# Java PPA
sudo add-apt-repository -y ppa:webupd8team/java
echo debconf shared/accepted-oracle-license-v1-1 select true | sudo debconf-set-selections
echo debconf shared/accepted-oracle-license-v1-1 seen true | sudo debconf-set-selections

# Update
sudo apt-get update
sudo apt-get upgrade -y

# Git
sudo apt-get install -y git

# Java
sudo apt-get install -y oracle-java7-installer

# Node.js
curl -sL https://deb.nodesource.com/setup_6.x | sudo -E bash -
sudo apt-get install -y nodejs

# Postgres
sudo apt-get install -y postgresql

# Golang
wget https://storage.googleapis.com/golang/go1.6.3.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.6.3.linux-amd64.tar.gz
rm go1.6.3.linux-amd64.tar.gz

# Set up Golang env
echo "export PATH=\$PATH:/usr/local/go/bin" >> /home/vagrant/.profile
mkdir -p /home/vagrant/go
echo "export GOPATH=/home/vagrant/go" >> /home/vagrant/.profile

echo "*************************************************************************"
echo "* VM ready. Complete these steps manually:                              *"
echo "* --------------------------------------------------------------------- *"
echo "* vagrant ssh                                                           *"
echo "* Create postgres user:                                                 *"
echo "*   https://github.com/h2oai/steam/tree/develop/tools/ec2/ubuntu-14.04 *"
echo "* go get github.com/h2oai/steam                                        *"
echo "* cd ~/go/src/github.com/h2oai/steam                                   *"
echo "* make                                                                  *"
echo "*************************************************************************"
