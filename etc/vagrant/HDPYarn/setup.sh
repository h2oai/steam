#!/usr/bin/env bash

# Installing JDK and Java
yum install -y java-1.7.0-openjdk
yum install -y java-1.7.0-openjdk-devel

# Installing Hadoop
echo "Using Hadoop Version $VERSION"
FILE="/vagrant/resources/hadoop-${VERSION}.tar.gz"
if [ -f "$FILE" ]; then
	echo "$FILE exists locally"
	echo "Importing Hadoop locally"
	cp $FILE .
else
	echo "$FIlE does not exist locally"
	echo "Downlading Hadoop remotely"
	wget http://apache.claz.org/hadoop/common/hadoop-${VERSION}/hadoop-${VERSION}.tar.gz 
fi
tar -xzf hadoop-$VERSION.tar.gz
rm hadoop-$VERSION.tar.gz
mv hadoop-${VERSION} /usr/local/hadoop
cp -v /vagrant/config/* /usr/local/hadoop/etc/hadoop/

# Setting up ssh
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

# Start YARN
sbin/start-yarn.sh

/etc/init.d/iptables save
/etc/init.d/iptables stop

# Setup Go
yum groupinstall -y 'Development Tools'
wget https://storage.googleapis.com/golang/go1.7.1.linux-amd64.tar.gz
tar -C /usr/local -xzf go1.7.1.linux-amd64.tar.gz

# # Setting up postgres
# sed -n -i.bak 'H;${x;s/\[base]\n/&exclude=postgres*\n/g; s/\[updates]\n/&exclude=postgres*\n/g;p;}' /etc/yum.repos.d/CentOS-Base.repo
# yum localinstall -y http://yum.postgresql.org/9.5/redhat/rhel-6-x86_64/pgdg-centos95-9.5-2.noarch.rpm
# yum install -y postgresql95-server

# service postgresql-9.5 initdb
# # this is required to get steam to connect to postgres; changes configurations on authentication; see http://www.cyberciti.biz/faq/psql-fatal-ident-authentication-failed-for-user/
# sed -n -i.bak 'H;${x;s/peer/trust/g; s/ident/trust/g;p;}' /var/lib/pgsql/9.5/data/pg_hba.conf

# service postgresql-9.5 start
# chmod 755 /home/vagrant
# sudo -u postgres createuser steam

echo export GOPATH=/home/vagrant/Go >> /home/vagrant/.bashrc
echo export PATH=$PATH:/usr/local/hadoop/bin:/usr/local/go/bin:$GOPATH/bin >> /home/vagrant/.bashrc
