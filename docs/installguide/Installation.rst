Installing Steam
================

Steam is supported on Ubuntu and Red Hat Enterprise Linux. Be sure to follow the instructions for your platform:

- `Ubuntu Installation`_
- `RHEL Installation`_

**Notes**: 

 - Admins should verify whether their Hadoop environment requires sudo. If it does, then users must have a root password/root access.

 - This installation creates a SQLite database. 

Ubuntu Installation
-------------------

This section describes how to install Steam on Ubuntu. 

Requirements for Steam with Ubuntu
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

- Ubuntu 12.04 or greater
- Steam .deb file. This is available via the `H2O download site for Steam <http://h2o.ai/download/steam>`__.
- JDK 1.7 or greater
- Chrome browser with an Internet connection. Note that Chrome is currently the only supported browser.
- H2O bleeding edge driver for your version of Hadoop. This is available from the `H2O Download page <http://h2o.ai/download>`__. Click the **Install on Hadoop** tab, and select the correct version for your environment.
- HAProxy 1.5 or greater. For Ubuntu, this is available from `haproxy.debian.net <https://haproxy.debian.net>`__

Install HAProxy for Ubuntu
~~~~~~~~~~~~~~~~~~~~~~~~~~

This section describes how to install HAProxy 1.5. You can skip this section if your environment already has HAProxy 1.5.or greater.

1. In your browser, go to `https://haproxy.debian.net <https://haproxy.debian.net>`__.
2. Select the system and version that you are running, then select an HAProxy version of 1.5-stable or greater. 
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

    wget http://h2o.ai/download/steam/steam_2.0.0_amd64.deb

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

At this point, you are ready to `Start Steam <Start_Steam.html>`__.

RHEL Installation
-----------------

This section describes how to install Steam on Red Hat Enterprise Linux.

Requirements for Steam with RHEL
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

- RHEL 6.8 or greater. Note that HAProxy is already included with this version of Red Hat. 
- Steam .rpm file. This is available via the `H2O download site for Steam <http://h2o.ai/download/steam>`__.
- JDK 1.7 or greater
- Chrome browser with an Internet connection. Note that Chrome is currently the only supported browser.
- H2O bleeding edge driver for your version of Hadoop. This is available from the `H2O Download page <http://h2o.ai/download>`__. Click the **Install on Hadoop** tab, and select the correct version for your environment.

Install HAProxy on RHEL
~~~~~~~~~~~~~~~~~~~~~~~

1. Note that RHEL 6.8 or greater includes HAProxy. Perform the following steps if you have not already installed HAProxy.

  :: 

    sudo yum haproxy

Install Steam on RHEL
~~~~~~~~~~~~~~~~~~~~~

1. Download the Steam .rpm file from the `H2O download site <http://h2o.ai/download/steam>`__. or via wget. For example:

  ::

    wget http://h2o.ai/download/steam/steam_2.0.0_amd64.rpm

2. Open a terminal window and ssh to your Hadoop edge node.

  ::
  
    ssh <user>@<hadoop_edge_node>

3. Copy the Steam .rpm file to your edge node.

  ::

    scp ... 

4. Run the Steam .rpm file.

  ::

    sudo rpm -i <steam_rpm_package>

5. Set the administrator username and password.

  ::

    sudo service steam set-admin
    username: administrator
    password: ***********

  **Note**: If you forget your local administrator username or password, you can rerun this command to reset the values.

6. Install the certificate and private key for the Steam server using one of the following methods:

  ::

    sudo service steam create-self-signed-cert
  
 Or 
   
   Add a certificate in **/etc/steam/private_key.pem**, **/etc/steam/cert.pem**.

7. (Optional) Change the service port numbers in **/etc/steam/steam.conf**.

At this point, you are ready to `Start Steam <Start_Steam.html>`__.
