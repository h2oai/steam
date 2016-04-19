#!/usr/bin/env bash

# Installing JDK and Java
yum install -y java-1.7.0-openjdk
yum install -y java-1.7.0-openjdk-devel

# Installing Hadoop
echo "Using Hadoop Version $VERSION"
if [ -f `/vagrant/resources/hadoop-${VERSION}.tar.gz`]; then
	echo "Importing Hadoop locally"
	cp /vagrant/resources/hadoop-${VERSION}.tar.gz .
else
	echo "Downlading Hadoop remotely"
	wget http://apache.claz.org/hadoop/common/hadoop-${VERSION}/hadoop-${VERSION}.tar.gz 
fi
tar -xzf hadoop-$VERSION.tar.gz
rm hadoop-$VERSION.tar.gz
mv hadoop-${VERSION} /usr/local/hadoop
cp -v /vagrant/config/* /usr/local/hadoop/etc/hadoop/

# # Setting up ssh
ssh-keygen -t dsa -P '' -f ~/.ssh/id_dsa
cat ~/.ssh/id_dsa.pub >> ~/.ssh/authorized_keys
chmod 0600 ~/.ssh/authorized_keys
for HOST in "localhost" "0.0.0.0"
do
	ssh-keyscan -t rsa $HOST >> ~/.ssh/known_hosts
done

# Setting up and starting hdfs
cd /usr/local/hadoop
bin/hdfs namenode -format -nonInteractive
sbin/start-dfs.sh
bin/hdfs dfs -mkdir /user
bin/hdfs dfs -mkdir /user/vagrant

# # Start YARN
sbin/start-yarn.sh

/etc/init.d/iptables save
/etc/init.d/iptables stop

echo export PATH:$PATH:/usr/local/hadoop/bin >> /home/vagrant/.bashrc
