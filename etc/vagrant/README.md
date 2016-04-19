YARN VAGRANT SETUP
====================

This will start a CentOS6.4 single node hadoop/yarn cluster.

(OPTIONAL) Installing Hadoop
-----------------

Install hadoop in the resources directory to speed up start time.
```
mkdir resources
cd /resources/
wget http://apache.claz.org/hadoop/common/hadoop-2.6.4/hadoop-2.6.4.tar.gz 
```


Starting Vagrant (and Steam)
----------------------------

Start vagrant
`vagrant up`

SSH into vagrant
`vagrant ssh`

Get the steam tar
```
cp /steam/steam--linux--amd64.tar.gz .
tar -xzvf steam--linux--amd64.tar.gz
```

Download the most recent h2odriver for hadoop 2.2
```
wget http://download.h2o.ai/download/h2o-3.8.2.2-hdp2.2?id=6187c90d-c016-1dc6-d382-3425ad1848f3 -O h2o-3.8.2.2-hdp2.2.zip
unzip h2o-3.8.2.2-hdp2.2.zip
cp h2o-3.8.2.2-hdp2.2/h2odriver steam--linux--amd64.tar.gz
```
