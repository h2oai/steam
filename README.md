# Installing, Starting, and Using Steam and the H2O Scoring Service

This document is meant for H2O developers and describes how to install, start, and use Steam on an in-house YARN cluster. External users should review the [README.md](docs/README.md) file within the **/docs** folder.

## Requirements
- Web browser and an Internet connection
- Go (available from <a href="https://golang.org">golang.org</a>) 
- Access to this steamY repository
- SSH access to a Jetty server running YARN
- Typescript
- Node.js
- JDK 1.7 or greater
- H2O AutoML for Apache HDP2.2 or CDH 5.5.3 (internal only)

## Building the H2O Scoring Service

The H2O Scoring Service Builder is an application that allows you to perform the following through either a web UI or command line:

- Compile a POJO, and then build a Jar file from a POJO and a gen-model file
- Compile the POJO, and then build a War file that is a service from a POJO, gen-model. You can then run the war file in Jetty, Tomcat, etc.
- Build a War file with a POJO predictor and Python pre-preprocessing

Perform the following steps to build the H2O Scoring Service Builder:

1. In a terminal window, navigate to the **steamY/scoring-service-builder** folder.

2. Run `./gradlew build` to build the service.

3. You will see a **BUILD SUCCESSFUL** message upon completion. Run 	`./gradlew jettyRunWar` to run the builder service.

4. Open a browser and navigate to localhost:55000 to begin using the H2O Scoring Service Builder web UI. 

### Testing the Scoring Service Builder

**Using the Web UI**

When the Builder Service is running, you can make a War file using the Web UI.

The following screenshot shows how to make a War file using a POJO file and a Jar file. Note that these files are included in the  **steamY/scoring-service-builder/examples/example-pojo** folder. 

![Make War](scoring-service-builder/images/make_war.png)


**Using the CLI**

Note that when the Builder Service is running, you can also make a war file using command line arguments. For example:

		curl curl -X POST --form pojo=@examples/example-pojo/gbm_3f258f27_f0ad_4520_b6a5_3d2bb4a9b0ff.java --form jar=@examples/example-pojo/h2o-genmodel.jar localhost:55000/makewar > example.war

 where:
 
 - `gbm_3f258f27_f0ad_4520_b6a5_3d2bb4a9b0ff.java` is the POJO file from H2O
 - `h2o-genmodel.jar` is the corresponding Jar file from your version of H2O

Both of the above files are included in the **steamY/scoring-service-builder/examples/example-pojo** folder.

#### Importing a War File

When the H2O Scoring Service Builder is up and running, open another terminal window, navigate to the **steamY/scoring-service-builder** folder, and run the following command to import the war file into the H2O Prediction service. The example below uses the **example.war** file that was built using the CLI in the previous section.

		java -jar jetty-runner-9.3.9.M1.jar --port 55001 example.war

This starts the H2O Prediction Service at localhost:55001. You can view this web service at http://localhost:55001.

![Example Service](scoring-service-builder/images/example_service.png)

Close the browser and the terminal window when the testing is complete.


## Steam Installation
Perform the following steps to install Steam.

1. Open a Terminal window and install Go (available from <a href="https://golang.org" target="_blank">golang.org</a>).

2. Export the GOPATH in the location where Steam will reside. This can be added to your bash profile or specified each time you run Steam. For example:

	`export GOPATH=$HOME/Desktop` 

3. Create the directory where Steam will reside. This must be in GOPATH location. For example:

	`mkdir -p $GOPATH/src/github.com/h2oai`
	
	>**Note**: The path MUST include `/src/github.com/h2o.ai`. 

4. Change directories to the new **h2oai** folder, and clone the repository in this folder. Enter your git username and password when prompted. 

		cd Desktop/src/github.com/h2oai
		git clone https://github.com/h2oai/steamY
	
5. Change directories to **steamY** and then run `make linux`.

		cd steamY
		make
		make linux

	>***Note***: If you are building this on your own local machine, then you can run `make` instead of `make linux`. This will only work if your local machine has YARN. 

You will see a `BUILD SUCCESSFUL` message when the installation is complete. At this point, you are ready to start using Steam. 


## Starting Steam
When Steam is installed, perform the following steps to start and use Steam. Note that two terminal windows will remain open: one for the Jetty server and one for Steam.

1. Open a terminal window. In the **steamY** directory, copy the **automl-hdp2.2.jar** file to your local machine. This engine is required in order to run any AutoML jobs in Steam. 

		scp <user>@<domain>:./automl-hdp2.2.jar .

2. ssh to the machine running YARN, changing `<user>` and `<domain>` below with appropriate values. Specify your password when prompted. 

	`ssh <user>@<domain>`
	
	The machine will include the following folder and files:
	
	- steam--linux-amd64.tar.gz
	- automl-hdp2.2.jar	
	
3. On the YARN machine, change directories to the Steam directory, then run the H2O Java Model Compilation Service. This service is required for Steam to be able to compile and build models for deployment purposes:

		cd steam--linux-amd64/steam
		java -jar var/master/assets/jetty-runner.jar --port 8811 var/master/assets/ROOT.war

5. Open another terminal window and again, ssh to the machine running YARN, changing `<user>` and `<domain>` below with appropriate values. Specify your password when prompted. 

	`ssh <user>@<domain>` 
	
6. Untar the **steam--linux-amd64.tar.gz** package.

	`tar -xzf steam--linux-amd64.tar.gz`

6. Change directories to the new **steam--linux-amd64** folder, then start Steam on the master server (the compilation server that will run the POJOs). For example, the following commands will start Steam on localhost.

		cd steam--linux-amd64
		./steam serve master --compilation-service-address="localhost:8080"
		
	>**Note**: You can view all available options for starting Steam on the master using `./steam help serve master`

	You will see a message similar to the following when Steam starts successfully.

		2016/04/28 13:34:56 steam v build 2016-04-28T20:15:00+0000
		2016/04/28 13:34:56 Working directory: /home/seb/steam--linux-amd64/var/master
		2016/04/28 13:34:56 WWW root: /home/seb/steam--linux-amd64/var/master/www
		2016/04/28 13:34:57 Priming datastore for first time use...
		2016/04/28 13:34:57 Datastore location: /home/seb/steam--linux-amd64/var/master/db/steam.db
		2016/04/28 13:34:57 Web server listening at 172.16.2.182:9000
		2016/04/28 13:34:57 Point your web browser to http://172.16.2.182:9000/

## Using Steam

In a Web browser, navigate to the Steam Web server (for example, http://172.16.2.182:9000).

### Adding an Engine
An empty Steam UI will display. Before performing any tasks, you must first add an Asset. 

1. Click the **Assets** icon (<img src="docs/images/icon_assets.png" alt="Thumbnail" style="width: 25px;" />) on the left navigation panel, then select **Add Engine**. 

	![](docs/images/add_engine.png)

2. Browse to the **automl-hdp2.2.jar** file on your local machine, then click **Upload**. 

	>***Note***: This file was added during the "Install Steam" steps. 

### Starting a Cloud

Clouds can be configured after the engine asset was successfully added. 

1.  Click the **Clouds** icon (<img src="docs/images/icon_clouds.png" alt="Thumbnail" style="width: 25px;" />) on the left navigation panel, then select **Start a Cloud**. 

2. Enter/specify the following information to set up your cloud:

	a. A name for the cloud

	b. The version of H2O that will run on the cloud.

	c. The number of nodes on the cloud.
	
	d. The amount of memory available on each node. Be sure to include the unit ("m" or "g").
	
	![](docs/images/add_cloud.png)
	
3. Click **Start Cloud** when you are finished.

>***Note***: You can view a stream of the cloud creation log in the terminal window that is running Steam. In the UI, Steam will respond with an error if the cloud configuration is incorrect (for example, if you specify more nodes than available on the cluster). 

The Cloud Details page opens upon successful completion. This page shows the cloud configuration information and includes a link to the H2O Flow URL. From this page, you can begin building model. 

![](docs/images/cloud_details.png)

### Adding a Model
Models are created from the Cloud Details page. When building a model, you will need to provide the location of the dataset that you will use as well as the response column. 

1. Click the **Build a Model** button on the bottom of the Cloud Details page.

2. Enter a path for the dataset that you want to use to build the model. 

>***Note***: If you choose to use a local dataset, then that dataset must reside in the same folder/path on each node in the cluster.

3. Specify the column that will be used as the response column in the model. 

4. Specify the maximum run time in seconds. H2O will return an error if the model build stalls after this threshold is reached.

5. Click **Start Building** when you are finished. 

	![](docs/images/build_model.png)

### Viewing Models	

Click the **Models** icon (<img src="docs/images/icon_models.png" alt="Thumbnail" style="width: 25px;" />) on the left navigation panel to view models that were successfully created. 

These models are processed using H2O's AutoML algorithm, which determines the best method to use to build the model. The model name includes this method. So, for example, if Steam returns a model named "DRF_model...", then this indicates that DRF was the algorithm that provided the best result.

### Deploying Models

After a model is built, the next step is to deploy the model in order to make/view predictions.

1. Click the **Models** icon (<img src="docs/images/icon_models.png" alt="Thumbnail" style="width: 25px;" />) on the left navigation panel.

2. Select the model that you want to use, then click the **Deploy this Model** button on the bottom of the page.

3. Specify the port to use for the scoring service. 

	>***Note***: Steam will return an error if you specify a port that is already being used.

4. Click **Deploy** when you are done.

### Making Predictions

Successfully deployed models can be viewed on the **Services** page. On this page, click the **Endpoint** link to open the H2O Prediction Service and begin making predictions.  

## Using Steam with Flow

As with other H2O products, Flow can be used alongside Steam when performing machine learning tasks.

On the Cloud Details page, click the Address link to open H2O Flow in a new tab. 

![](docs/images/h2o_flow.png)

>***Note***: Refer to the H2O Flow documentation for information on how to use Flow. 

## Stopping Steam	

When you are finished, use the following process to safely shut down Steam:

1. On the Services page, stop all running services.

2. Stop all running clouds.
