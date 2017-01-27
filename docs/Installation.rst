Installing and Starting Steam
=============================

This section applies to Administrators and describes how to install and start Steam in a Hadoop environment and make it accessible to a set of users. The process includes uploading an H2O driver and adding users via an LDAP config file. 

Refer to the following topics. Be sure to follow the instructions for your platform:

- `Installing on Ubuntu`_
- `Installing on RHEL`_
- `Starting Steam`_
- `Adding an H2O Driver`_
- `Configure LDAP Connection Settings`_

**Notes**: 

 - Admins should verify whether their Hadoop environment requires sudo. If it does, then users must have a root password/root access.

 - This installation creates a SQLite database. 

Installing on Ubuntu 
--------------------

This section describes how to install Steam on Ubuntu. 

Requirements for Steam with Ubuntu
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

- Ubuntu 12.04 or greater
- Steam .deb file. This is available via S3 download.
- JDK 1.7 or greater
- Chrome browser with an Internet connection. Note that Chrome is currently the only supported browser.
- H2O bleeding edge driver for your version of Hadoop. This is available from the `H2O Download page <http://h2o.ai/download>`__. Click the **Install on Hadoop** tab, and select the correct version for your environment.
- HAProxy 1.6 or greater. For Ubuntu, this is available from `haproxy.debian.net <https://haproxy.debian.net>`__


Install HAProxy for Ubuntu
~~~~~~~~~~~~~~~~~~~~~~~~~~

This section describes how to install HAProxy 1.6. You can skip this section if your environment already has HAProxy 1.6.or greater.

1. In your browser, go to `https://haproxy.debian.net <https://haproxy.debian.net>`__.
2. Select the system and version that you are running, then select an HAProxy version of 1.6-stable or greater. 
3. Open a Terminal window and run the commands that are listed (using ``sudo`` if required). The example below shows the commands to use with Ubuntu version Trusty (14.04 LTS) and HAProxy version 1.7-stable. 
   
   .. figure:: images/haproxy_ubuntu.png
      :alt: HAProxy for Ubuntu

Install Steam on Ubuntu
~~~~~~~~~~~~~~~~~~~~~~~

1. Open a terminal window and ssh to your Hadoop edge node.

  ::
  
    ssh <user>@<hadoop_edge_node>

2. Download the Steam .deb file. For example:

  ::

    wget https://s3.amazonaws.com/steam-release/steam_1.1.0.93_amd64.deb

3. Unpackage the Steam .deb file.

  ::
    
    sudo dpkg -i steam_1.1.0.69_amd64.deb

4. Set the administrator username and password.

  ::

    sudo service steam set-admin
    username: administrator
    password: ***********

  **Note**: If you forget your local administrator username or password, you can rerun this command to reset the values.

5. Install the certificate and private key for the Steam server using one of the following methods:

  ::

    sudo service steam create-self-signed-cert
  
 Or 
   
   Add a certificate in **/etc/steam/private_key.pem**, **/etc/steam/cert.pem**.


6. (Optional) Change the service port numbers in **/etc/steam/steam.conf**.

At this point, you are ready to `Start Steam <installation.html#starting-steam>`__.

Installing on RHEL
------------------

This section describes how to install Steam on RHEL <version???>

Requirements for Steam with RHEL
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

- RHEL 6.6
- Steam .deb file. This is available via S3 download.
- JDK 1.7 or greater
- Chrome browser with an Internet connection. Note that Chrome is currently the only supported browser.
- H2O bleeding edge driver for your version of Hadoop. This is available from the `H2O Download page <http://h2o.ai/download>`__. Click the **Install on Hadoop** tab, and select the correct version for your environment.
- HAProxy 1.6 or greater.

Install HAProxy for RHEL
~~~~~~~~~~~~~~~~~~~~~~~~

<Steps tbd>

Install Steam on RHEL
~~~~~~~~~~~~~~~~~~~~~

1. Open a terminal window and ssh to your Hadoop edge node.

  ::
  
    ssh <user>@<hadoop_edge_node>

2. Download the Steam .deb file. For example:

  ::

    wget https://s3.amazonaws.com/steam-release/steam_1.1.0.93_amd64.deb


<more steps tbd>

At this point, you are ready to `Start Steam <installation.html#starting-steam>`__.

Starting Steam
--------------

1. Start Steam by running the following command on your YARN edge node. Refer to the `Steam Start Flags`_ section for additional flags that can be specified when starting Steam.

  ::
    
    sudo service steam start
   
2. (Optional) Check the log file to verify that Steam starts correctly:

  ::

    sudo cat /var/log/steam.log


At this point, you can open a Chrome browser and navigate to your Hadoop edge node (where Steam is currently running). For example, ``https://<hadoop-edge-node>:9000``. Note that in your browser, you may be required to authenticate using the Administrator username and password that you created during the installation process. 

Steam Start Flags
~~~~~~~~~~~~~~~~~

The following table lists the options/flags that can be added to the ``service steam start`` command when starting Steam. Use ``./steam serve master --help`` or ``./steam serve master -h`` for the most up-to-date list of available options.

+-------------------------------------------+-----------------------------------------+
| Flag                                      | Description                             |
+===========================================+=========================================+
| ``--authentication-config=``              | Specify a configuration file to use     |
|                                           | for authentication.                     |
+-------------------------------------------+-----------------------------------------+ 
| ``--authentication-provider=``            | Specify either ``basic`` or ``digest``  |
|                                           | as the authentication mechanism for     |
|                                           | client logins.                          |
+-------------------------------------------+-----------------------------------------+ 
| ``--cluster-proxy-address=``              | Specify a proxy address. For example:   |
|                                           | ``<ip>:<port>`` or ``:<port>``.         |
+-------------------------------------------+-----------------------------------------+
| ``--compilation-service-address=``        | Specify an address to use for the       |
|                                           | compilation service. For example:       |
|                                           | ``<ip>:<port>`` or ``:<port>``.         |
+-------------------------------------------+-----------------------------------------+
| ``--prediction-service-host=``            | Specify the hostname to use for the     |
|                                           | prediction service.                     |
+-------------------------------------------+-----------------------------------------+
| ``--prediction-service-port-range=``      | Specify a range of ports to create      |
|                                           | prediction services on. For example:    |
|                                           | ``<from_port>:<to_port>``.              |
+-------------------------------------------+-----------------------------------------+
| ``--profile=``                            | Specify ``true`` to enable the Go       |
|                                           | profiler.                               |
+-------------------------------------------+-----------------------------------------+
| ``--admin-name=``                         | Set the admin username. This is         |
|                                           | required at first-time-use only.        |
+-------------------------------------------+-----------------------------------------+
| ``--admin-password=``                     | Set the admin password. This is         |
|                                           | required at first-time-use only.        |
+-------------------------------------------+-----------------------------------------+
| ``--web-address=``                        | Specify the web server address. For     |
|                                           | example: ``<ip>:<port>`` or ``:<port>``.|
+-------------------------------------------+-----------------------------------------+
| ``--web-tls-cert-path=``                  | Specify the web server TLS certificate  |
|                                           | path.                                   |
+-------------------------------------------+-----------------------------------------+
| ``--web-tls-key-path=``                   | Specify the web server TLK key file     |
|                                           | path.                                   |
+-------------------------------------------+-----------------------------------------+
| ``--working-directory=``                  | Specify the working directory for       |
|                                           | application files.                      |
+-------------------------------------------+-----------------------------------------+
| ``--yarn-enable-kerberos=``               | Specify whether to enable Kerberos      |
|                                           | authentication. This requires a username|
|                                           | and keytab.                             |
+-------------------------------------------+-----------------------------------------+


Adding an H2O Driver
--------------------

**Note**: Currently, only the H2O bleeding edge release is supported on Steam. 

1. On your local machine, download the h2odriver from the `H2O Download page <http://h2o.ai/download>`__. Be sure to select your version of Hadoop. For example:

  ::

    wget http://h2o-release.s3.amazonaws.com/h2o/master/3756/h2o-3.11.0.3756-hdp2.2.zip

2. In the Steam UI, navigate to the **Clusters** page and select **Launch New Cluster**.

3. In the H2O Version section of the **Launch New Cluster page**, click the **Choose File** button and browse to the H2O driver that you just downloaded.

4. Click **Upload Engine** to add the egine to Steam.

   .. figure:: images/upload_engine.png
      :alt: Upload Engine
    
A message will display indicating that the engine was successfully uploaded.

Configure LDAP Connection Settings
----------------------------------

1. Navigate to the **Users** page and select the **Authentication** tab. 

 **Note**: Only admins have access to the Steam Users page.

2. Select LDAP in the **User DB Type** drop down menu, then configure the LDAP connection settings. (Refer to the table below and the image that follows.)

 +---------------------------+------------------------------+------------------------------------------+
 | Field                     | Description                  | Example                                  |
 +===========================+==============================+==========================================+
 | Host                      | The LDAP host server address | ldap.0xdata.loc                          |
 +---------------------------+------------------------------+------------------------------------------+
 | Port                      | The LDAP server port         | 389                                      |
 +---------------------------+------------------------------+------------------------------------------+
 | SSL-Enabled               | Enable this if your LDAP     |                                          |
 |                           | supports SSL.                |                                          |
 +---------------------------+------------------------------+------------------------------------------+
 | Bind DN                   | The Distinguished Name used  | cn=admin,dc=0xdata,dc=loc                |
 |                           | by the LDAP server if        |                                          |
 |                           | extended access is required. |                                          |
 |                           | This can be left blank if    |                                          |
 |                           | anonymous bind is sufficient.|                                          |
 +---------------------------+------------------------------+------------------------------------------+
 | Bind DN Password          | The password for the Bind DN | h2o                                      |
 |                           | user                         |                                          |
 +---------------------------+------------------------------+------------------------------------------+
 | User Base DN              | The location of the LDAP     | ou=users,dc=0xdata,dc=loc                |
 |                           | users, specified by the DN of|                                          |
 |                           | your user subtree            |                                          |
 +---------------------------+------------------------------+------------------------------------------+
 | User Base Filter          | The LDAP search filter used  | department=IT                            |
 |                           | to filter users              |                                          |
 +---------------------------+------------------------------+------------------------------------------+
 | User Name Attribute       | The User Attribute that      | uid                                      |
 |                           | contains the username        |                                          |
 +---------------------------+------------------------------+------------------------------------------+
 | Group DN                  | The Distinguished Name used  | cn=jettygroup,ou=groups,dc=0xdata,dc=loc |
 |                           | for group synch              |                                          |
 +---------------------------+------------------------------+------------------------------------------+
 | Static Member Attribute   | The attribute for static     | memberUid                                |
 |                           | group entries                |                                          |
 +---------------------------+------------------------------+------------------------------------------+
 | Search Request Size Limit | Limit the size of search     |                                          |
 |                           | results. 0 indicates         |                                          |
 |                           | unlimited.                   |                                          |
 +---------------------------+------------------------------+------------------------------------------+
 | Search Request Time Limit | Limit the time allotted for  | 0                                        |
 |                           | completing search results. 0 |                                          |
 |                           | indicates unlimited.         |                                          |
 +---------------------------+------------------------------+------------------------------------------+

 .. figure:: images/ldap_authentication_config.png
    :alt: LDAP Configuration
  
3. Click **Test Config** when you are done. A valid response message indicates that the configuration was successful.
4. Click **Save Config**.
5. In order for the configuration to take effect, you must log out and restart (or stop then start) Steam.

  ::
    
    sudo service steam restart

After LDAP is configured, users can log in to Steam using their LDAP username and password. 
