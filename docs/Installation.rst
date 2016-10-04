Installing and Starting Steam
=============================

This section describes how administrators can install and start Steam. Refer to the following topics:

-  `Requirements`_
-  `Installing and Starting Steam on YARN`_
-  `Installing and Starting Steam on a Local Machine`_
-  `Next Steps`_

Requirements
------------

-  Chrome browser with an Internet connection. Note that Chrome is currently the only supported browser.
-  Steam tar file

   -  available from the `H2O Download <http://h2o.ai/download>`__ site

-  JDK 1.7 or greater
-  H2O jar file for version 3.10.0.7 or greater

   -  available from the H2O Download page
   -  If necessary, follow the instructions on the
      http://www.h2o.ai/download/h2o/python or
      http://www.h2o.ai/download/h2o/r page to upgrade H2O for Python or
      R.

Installing and Starting Steam on YARN
-------------------------------------

A Steam YARN installation provides a method for data scientists and developers to work collaboratively when trainging and deploying models. 

Installation
~~~~~~~~~~~~

Perform the following steps to install Steam on a YARN edge node. 

1. Go to the `H2O Download <http://h2o.ai/download>`__ site to download Steam. 

2. Change directories to the Steam download file, copy the file to your edge node, then untar the Steam file. Be sure to enter the correct password when prompted. For example:

   ::

       cd ~/Downloads/steam-1.0.0-darwin-amd64
       scp -r steam-1.0.0-darwin-amd64 <user>@<yarn_edge_node>:~
       tar -xzvf steam-1.0.0-darwin-amd64.tar.gz 


Start Steam on YARN
~~~~~~~~~~~~~~~~~~~

After Steam is installed on the YARN edge node, the next step is to provide the designated Steam superuser with
the URL of the edge node and the superuser login credentials. The superuser can then start Steam and begin creating roles, workgroups, and users.

1. SSH into the YARN edge node where the Steam package was copied. Note that this step requires superuser privileges. 

 ::

  ssh <user>@<yarn_edge_node>

2. Open another terminal window and start the Jetty server from within the Steam folder. This allows you to deploy models and run the Steam Prediction Service.

  ::

    java -jar var/master/assets/jetty-runner.jar var/master/assets/ROOT.war

3. Open another terminal window and run the following command to start Steam. Be sure to include the ``--superuser-name=superuser`` and ``--superuser-password=superuser`` flags. (Or provide a more secure password.) This starts Steam on the edge node at port 9000 and creates a Steam superuser. The Steam superuser is responsible for creating roles, workgroups, and users and maintains the H2O cluster.

 ::

  sudo ./steam serve master --superuser-name=superuser --superuser-password=superuser

 **Note**: Use ``./steam serve master --help`` or ``./steam serve master -h`` for information on how to start the compilation and/or prediction service on a different location.

4. Open a Chrome browser and navigate to the YARN edge node.

Installing and Starting Steam on a Local Machine
------------------------------------------------

Users can download and install steam directly on a local machine without the need for a running instance of YARN. 

Installation
~~~~~~~~~~~~

1. Go to the `H2O Download <http://h2o.ai/download>`__ site and download Steam. 

2. Change directories to the Steam download file and untar the file.

 ::
    
    cd ~/Downloads/steam-1.0.0-darwin-amd64
    tar -xzvf steam-1.0.0-darwin-amd64.tar.gz 

That's it! You are now ready to start Steam.

Start Steam Locally
~~~~~~~~~~~~~~~~~~~

After Steam is installed, the following steps describe how to start Steam.

1. Navigate to the untarred Steam folder. 

 ::

  cd steam--darwin-amd64

2. Open another terminal window and start the Jetty server from within the Steam folder. This allows you to deploy models and run the Steam Prediction Service.

  ::

    java -jar var/master/assets/jetty-runner.jar var/master/assets/ROOT.war

3. Open another terminal window and start Steam. Be sure to include the ``--superuser-name=superuser`` and
   ``--superuser-password=superuser`` flags. (Or provide a more secure password.) This creates Steam superuser. A Steam superuser is responsible for creating roles,workgroups, and users. This also starts the Steam web service on ``localhost:9000``, the compilation service on ``localhost:8080`` (same as the Jetty server), and the prediction service on the external IP address of ``localhost``. You can change these using ``--compilation-service-address=<ip_address:port>`` and ``--prediction-service-address=<ip_address>``. Use ``./steam serve master --help`` or ``./steam serve master -h`` to view additional options.

 ::

  ./steam serve master --superuser-name=superuser --superuser-password=superuser

 **Note**: If you are demoing Steam and do not have an Internet connection, you can set the prediction service to point to localhost using ``--prediction-service-address=localhost``. 

4. Open a Chrome browser and navigate to http://localhost:9000.

Next Steps
----------

Now that Steam is up and running, you can log in to the machine that is
running Steam and use the CLI to create additional roles, workgroups,
and users. Refer to the `User Management <UserManagement.html>`__ section.
