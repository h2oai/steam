# Steam Standard Standalone Demo


This demo describes how to use Steam without the need for a local running instance of YARN. This demo will walk through the following procedures:

- Installing and starting Steam, the Compilation Service, and H2O
- Building a simple model in Python (Optional for users who don't have an existing demo)
- Deploying the model using Steam

During this demo, three terminal windows will remain open for the Steam, Scoring, and H2O services. A fourth terminal window will be used to run H2O commands in a Python example. 

Finally, these steps were created using H2O version 3.8.2.8, and that version resides in a Downloads folder. Wherever used, this version number and path should be adjusted to match your version and path.

## Requirements

- Web browser with an Internet connection
- JDK 1.7 or greater
- Steam binary for Linux or OS X 
	- `s3cmd get s3://steam-release/steamY-master-linux-amd64.tar.gz`
	- `s3cmd get s3://steam-release/steamY-master-darwin-amd64.tar.gz`
- H2O jar file
	- available from the <a href="http://h2o-release.s3.amazonaws.com/h2o/rel-turchin/8/index.html"> H2O Download</a> page
- A dataset that will be used to generate a model. This demo uses the well-known iris.csv dataset, and the dataset is saved onto the desktop. 
- Python 2.7 (for building the model)

## Installing and Starting Steam, the Compilation Service, and H2O

1. Open a terminal window and download the steamY binary. Note that the command below downloads the OS X binary. Replace `darwin` with `linux` in the steps that follow to build on Linux.

	`user$ s3cmd get s3://steam-release/steamY-master-darwin-amd64.tar.gz`

2. Untar the steamY binary.

	`user$ tar xvf steamY-master-darwin-amd64.tar.gz`

3. Change directories to your Steam directory, and start the compilation service.

		user$ cd steam-master-darwin-amd64
		user$ java -jar var/master/assets/jetty-runner.jar var/master/assets/ROOT.war
		
4. Open another terminal window. Navigate to the Steam folder and start Steam. Replace the IP address with the IP address of your local machine.

		user$ cd steam-master-darwin-amd64/
		user$ ./steam serve master --compilation-service-address=192.16.2.119:8080 --scoring-service-address=192.16.2.119

	>***Note***: Use `./steam help serve master` or `./steam serve master -h` to view additional options

5. Open another terminal window. Navigate to the folder with your H2O jar file and start H2O. This will create a one-node cluster on your local machine.

		user$ cd ~/Downloads/h2o-3.8.2.8
		user$ java -jar h2o.jar 
		
	Note the IP address and port in the output. You will enter this information in Steam when connecting to your local cluster. For example:
	
		...
		06-07 15:36:36.642 192.16.2.119:54321    15931  main      INFO: Cloud of size 1 formed [/192.16.2.119:54321] 
		...

6. Point you browser to the Steam url, for example, http://192.16.2.119:9000/.
 
7. In left pane, select the **Clusters** tab, then click the **Connect To Cluster** button to setup Steam with H2O. Specify the IP address and port of the cluster currently running H2O (for example, 192.16.2.119:54321), then click **Register Cluster**. 

	![](images/register_cluster.png)

You are now ready to build a model on this cluster in Python. 

## Building a Model in Python

>**Note**: This section can be skipped if you already have demo steps that you use in R, Python, or Flow. If you use another demo, be sure that you initialize H2O on your local cluster so that the data will be available in Steam.

The steps below show how to build model using the Iris dataset and the GBM algorithm. The steps will be run using H2O in Python. Once created, the model can be deployed in Steam. 


1. Open a terminal window. Change directories to the H2O folder, and start Python. Import the modules that will be used for this demo. 

		$ cd ~/Downloads/h2o-3.8.2.8
		$ python
		>>> import h2o
		>>> from h2o.estimators.gbm import H2OGradientBoostingEstimator

2. Initialize H2O using your local IP address and port. 

		>>> h2o.init(ip=“192.16.2.119”, port=54321)
		------------------------------  -------------------------------------
		H2O cluster uptime:             2 minutes 37 seconds 168 milliseconds
		H2O cluster version:            3.8.2.8
		H2O cluster name:               user
		H2O cluster total nodes:        1
		H2O cluster total free memory:  3.35 GB
		H2O cluster total cores:        8
		H2O cluster allowed cores:      8
		H2O cluster healthy:            True
		H2O Connection ip:              192.16.2.119
		H2O Connection port:            54321
		H2O Connection proxy:
		Python Version:                 2.7.9
		------------------------------  -------------------------------------

3. Upload the Iris dataset. Note that in this example, Python is running from the Downloads folder, and the Iris dataset is on the Desktop:

		>>> df=h2o.upload_file("../../Desktop/iris.csv")

4. Specify the configuration options to use when building a GBM model.

		>>> gbm_regressor = H2OGradientBoostingEstimator(distribution="gaussian", ntrees=10, max_depth=3, min_rows=2, learn_rate="0.2")

5. Train the model using the Iris dataset (`df` object) and the GBM configuration options. 

		>>> gbm_regressor.train(x=range(1, df.ncol), y=0, training_frame=df)

6. Optionally view the model details.

		>>> gbm_regressor

Once created, the model will be visible in the Steam UI. 

## Deploying a Model in Steam

1. In the Steam UI, Select **Cluster** > **Models**. Select a model from your demo, and then click **Import Model to Steam**. This pulls the model into Steam. Once imported, the model can then be deployed to the scoring service.

  ![](images/import_model.png)

2. Select the **Models** tab in the left pane. You should see the model that you just imported. Select this model, and then click **Deploy This Model** to create scoring services for the model.

	![](images/deploy_model.png)

3. Specify the port number for the scoring service (defaults to 8000), then click **Deploy**.

4. Select the **Services** tab in the left pane.

5. Select a service (in this case, the model you just deployed), and click the link in the **Endpoint** field to reach the scoring service.

	![](images/select_service.png)

6. Make predictions using one of the following methods:
    
    - Specify input values based on column data from the original dataset, and then press the first **Predict** button.
    - Enter a query string using the format `field1=value1&field2=value2` (for example, `C2=2.5&C3=3.4`), and then press the second **Predict** button.
    
    ![](images/prediction_service.png)
    
    >***Notes***:
    
	- Click the **URL** link below a prediction result. The predict variables will display in the URL (for example, http://192.16.2.119:8000/predict?C2=2.5&C3=3.4) 

	- 	Use the **Clear** buttons to clear all entries and begin a new prediction.
     
	- You can view additional statistics about the scoring service by clicking the **More statistics at** link.
