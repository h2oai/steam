# Steam Installation and Setup

This section describes how to install and start Steam. Refer to the following topics:

- [Requirements](#requirements)
- [Linux and Mac OS X Installation](#installationlin)
- [RedHat 6.7 Installation](#installationrhat)
- [Start Steam](#startsteam)

## <a name="requirements"></a>Requirements

- Web browser with an Internet connection
- Steam tar for OS X, Linux, or RedHat 6.7
	- available from <a href="http://www.h2o.ai/steam/">h2o.ai/steam/</a>
- JDK 1.7 or greater
- PostgreSQL 9.1 or greater
	- available from <a href="https://www.postgresql.org/" target="_blank">PostgreSQL.org</a>
- H2O jar file for version 3.10.0.3 or greater
	- available from the <a href="http://www.h2o.ai/download/h2o/choose" target="_blank">H2O Download</a> page
	- If necessary, follow the instructions on the <a href="http://www.h2o.ai/download/h2o/python">http://www.h2o.ai/download/h2o/python</a> or <a href="http://www.h2o.ai/download/h2o/r">http://www.h2o.ai/download/h2o/r</a> page to upgrade H2O for Python or R.


## <a name="installationlin"></a>Linux and Mac OS X Installation
Perform the following steps to install Steam on Linux or Mac OS X. 

***Note***: This installation should only be performed on a YARN edge node.

### Download the Software to the Edge Node

1. Go to <a href="http://www.h2o.ai/steam/">h2o.ai/steam/</a> to download Steam. Be sure to accept the EULA.

2. Change directories to the Steam download folder for your OS (Linux or OS X), then copy the folder to your edge node. Enter the correct password when prompted.

		cd ~/Downloads/steam-0/steam-0.1.0-darwin-amd64
		scp -r steam-0.1.0-darwin-amd64 <user>@<yarn_edge_node>:~

You need to create the Steam superuser before you can start Steam. The next sections describe starting PostgreSQL, creating the Steam superuser, then creating the Steam database.

### Start PostgreSQL

Open a terminal window and run the following command to start PostgreSQL. This should be started from the folder where PostgreSQL was installed.

		postgres -D /usr/local/var/postgres

### Create the Steam Superuser

The Steam superuser is responsible for maintaining Steam clusters and for setting up roles, workgroups, and users. This step creates the superuser for the Steam database and then creates the database. The example below creates a Steam superuser with a password ``superuser``, and then creates the Steam database. ***If prompted, do not enter a password***.

		createuser -P steam 
		Enter password for new role:
		Enter it again:
		
### Create the Steam Database

The following commands show how to change directories to the Steam **/var/master/scripts** folder, and then create the database.

		cd steam-master-darwin-amd64/var/master/scripts
		./create-database.sh

## <a name="installationrhat"></a>RedHat 6.7 Installation
Perform the following steps to install Steam on RedHat 6.7. This is currently the only supported version of RedHat and can be downloaded using the following command:

		wget https://dl.fedoraproject.org/pub/epel/epel-release-latest-6.noarch.rpm 

***Note***: This installation should only be performed on a YARN edge node.

### Install and Start PostgreSQL

1. SSH to the Steam edge node, and then add the following line to the **[main]** section of /etc/yum/pluginconf.d/rhnplugin.conf.
		
 ```ssh <user>@<yarn_edge_node>```
	
 ```exclude=postgresql*	```	

1. Run the following commands to install PostgreSQL.

 ```sudo yum localinstall http://yum.postgresql.org/9.4/redhat/rhel-6-x86_64/pgdg-redhat94-9.4-2.noarch.rpm```
		
 ```sudo yum install postgresql94-server```

1. After the PostgreSQL server is installed, run the following command as the postgres user.

 ```/usr/pgsql-9.4/bin/initdb -D /var/lib/pgsql/9.4/data```
		
1. Run the following command as root to start PostgreSQL.

 ```sudo /etc/init.d/postgresql-9.4 start```
		
1. Run the following commands as the postgres user to create a Steam user. Note that a password is not required and can be left blank.

		sudo -i -u postgres
		createuser --interactive -P steam
		Enter password for new role:  # remember this password 
		Enter it again: 
		Shall the new role be a superuser? (y/n) n
		Shall the new role be allowed to create databases? (y/n) y
		Shall the new role be allowed to create more new roles? (y/n) n

### Create the Steam Database

Now that PostgreSQL is running and a Steam user is created, the following commands show how to change directories to the Steam **/var/master/scripts** folder, and then create the Steam database.

		cd steam-master-darwin-amd64/var/master/scripts
		./create-database.sh
		
### Set Up .pgpass for PostgreSQL

As the Steam user, set up .pgpass for postgres. This is done by editing **~/.pgpass** and appending the line ``*:*:*:steam:pa$$word``.

		chmod 600 ~/.pgpass

## <a name="startsteam"></a>Start Steam

After Steam is installed on the YARN edge node and a superuser is created, the next step is to provide the designated Steam superuser with the URL of the edge node and the superuser login credentials. The superuser can then start Steam and begin creating roles, workgroups, and users. 

1. SSH into the YARN edge node where the Steam package was copied. 

 ```ssh <user>@<yarn_edge_node>```

1. Start the Steam compilation and scoring service. Be sure to include the ``--superuser-name=superuser`` and ``--superuser-password=superuser`` flags. (Or provide a more secure password.) This starts Steam on localhost:9000 and creates a Steam superuser. The Steam superuser is responsible for creating roles, workgroups, and users and maintains the H2O cluster.

 ```./steam serve master --compilation-service-address=<yarn_edge_node>:<port> --scoring-service-address=<ip_address> --superuser-name=superuser --superuser-password=superuser```

>**Note**: This starts the Steam web service on ``<yarn_edge_node>:<port>`` and the scoring service on ``<ip_address>``. Use ``./steam help serve master`` or ``./steam serve master -h`` to view additional help information.

Now that Steam is up and running, you can log in to the machine that is running Steam and use the CLI to create additional roles, workgroups, and users. Refer to the [User Management](UserManagement.md) section.
