YARN VAGRANT SETUP
====================

Installing Hadoop
-----------------

Install hadoop in the vagrant directory.
```
wget http://apache.claz.org/hadoop/common/hadoop-2.6.4/hadoop-2.6.4.tar.gz 
tar xzf hadoop-2.6.4.tar.gz 
ln -s hadoop-2.6.4.tar.gz hadoop
```

Configuring Hadoop
------------------

`cd hadoop` edit These files

`etc/hadoop/core-site.xml`
```
<configuration>
    <property>
        <name>fs.defaultFS</name>
        <value>hdfs://localhost:9000</value>
    </property>
</configuration>
```

`etc/hadoop/hdfs-site.xml`
```
<configuration>
    <property>
        <name>dfs.replication</name>
        <value>1</value>
    </property>
</configuration>
```

`etc/hadoop/mapred-site.xml`
```
<configuration>
    <property>
        <name>mapreduce.framework.name</name>
        <value>yarn</value>
    </property>
</configuration>
```

`etc/hadoop/yarn-site.xml`
```
<configuration>
    <property>
        <name>yarn.nodemanager.aux-services</name>
        <value>mapreduce_shuffle</value>
    </property>
</configuration>
```

Starting Vagrant
----------------

Move the the HDPYarn directory
`cd ../HDPYarn`

Start vagrant
`vagrant up`

SSH into vagrant
`vagrant ssh`

The Yarn Application manager:
`localhost:8088`
