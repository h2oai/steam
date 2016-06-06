Steam Standard Standalone Demo
==============================

This is a demo to use Steam without the need for a local running instance of YARN.

Preparation Steps
-----------------
These steps should be run prior to the actual demo. In order to run these a copy of the Steam tar file is required as well as either a running H2O cluster or a copy of the H2O jar file. Steam can be downloaded from s3 depending on your OS. 
* `s3cmd get s3://steam-release/steamY-master-linux-amd64.tar.gz`
* `s3cmd get s3://steam-release/steamY-master-darwin-amd64.tar.gz` 

1. Start H2O (optional). If not using a remote H2O cluster, start H2O using the standard `java -jar h2o.jar`
2. Untar the Steam tar file. At this point the `$STEAM` will refer to your OS version of Steam, either Linux(`steamY-master-linux-amd64`) or OSX(`steamY-master-darwin-amd64`).
    1. Untar the file: `tar xzvf $STEAM.tar.gz`
    2. Move to the directory: `cd $STEAM`
3. Start compilation service
    1. Move to the assets directory: `cd ./var/master/assets`
    2. Run jetty `java -jar jetty-runner.jar ROOT.war`
        * (Optionally) If the default port 8080 is in use, instead use: `java -jar jetty-runner.jar --port [PORT] ROOT.war`. **NOTE** If this option is used, see the note in the next step.
4. Start Steam
    1. Return to the root steam directory. `../../..` or `cd $STEAM`
    2. Issue the steam start command: `./steam serve master`
        * **NOTE** If the compilation service was started using a non-default port specify the port using `--compilation-service-address=localhost:[PORT]`
        * For additional commands use: `./steam help serve master` or `./steam serve master -h`

Steam Demo
----------
After running the inital steps, the following will walk through the presentable demo using the url recieved from the last step.

1. Point browser to Steam url.
2. In left pane, select the `Clusters` tab. (initially selected)
3. Use the `Connect To Cluster` button to setup Steam with H2O. (point to H2O cluster)
4. Run standard H2O Demo. (Flow/R/Python/etc.)
5. Select `Cluster` > `Models`. (May require user to reselect cluster)
6. Use the `Import Model To Steam` to save relevant models.
7. Select the `Models` tab in the left pane.
8. Use the `Deploy This Model...` button to create scoring services for the models.
9. Select the `Services` tab in the left pane.
10. Select a service and use the Endpoint field to reach the scoring service.
11. Make predictions using:
    - fields and dropdown menus. Press the first `Predict` button
    - query string with the format `field1=value1&field2=value2`. Press the second `Predict` button.

