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

The following table lists the options/flags that can be added to the ``service steam start`` command when starting Steam. Use ``service steam start --help`` or ``service steam start -h`` for the most up-to-date list of available options.

+-------------------------------------------+-----------------------------------------+
| Flag                                      | Description                             |
+===========================================+=========================================+
| ``--admin-name=``                         | Set the admin username. This is         |
|                                           | required at first-time-use only.        |
+-------------------------------------------+-----------------------------------------+
| ``--admin-password=``                     | Set the admin password. This is         |
|                                           | required at first-time-use only.        |
+-------------------------------------------+-----------------------------------------+
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
