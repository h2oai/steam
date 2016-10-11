#### H2O Prediction Service Builder
#Spam Detection with Python Preprocessing

In this example we'll build a spam detector for text messages. The model is built from a data file. 
Feature preprocessing is done in Python. 
After that we create a Prediction Web Service that includes the Python preprocessing.

This example is in ```examples/spam-detection-python``` in the ```steam/prediction_service_builder``` directory.

## Train and preprocess data (Optional)

This step is optional.

Python and the H2O package must be installed on the machine. Numpy, Sklearn and textblob must also be installed.

The data file in ```data/smsData.txt```. This file contains messages with a label if 
it's spam or not, "ham".

```
ham    Ok... But they said i've got wisdom teeth hidden inside n mayb need 2 remove.
ham    thk of wat to eat tonight.
ham    I dunno until when... Lets go learn pilates...
spam   Someonone you know is trying to contact you via our dating service! To find out who it could be call from your mobile or landline 09064015307 BOX334SK38ch 
ham	   Ok c  then.
spam   URGENT! We are trying to contact U. Todays draw shows that you have won a 800 prize GUARANTEED. Call 09050003091 from land line. Claim C52. Valid12hrs only
ham	   staff.science.nus.edu.sg/~phyhcmk/teaching/pc1323
ham	   Thank god they are in bed!
ham	   Hey tmr meet at bugis 930 ?
spam   Todays Voda numbers ending with 7634 are selected to receive a 350 reward. If you have a match please call 08712300220 quoting claim code 7684 standard rates apply.
```

```train.py``` is the program we use to preprocess the data. It reads the data file and converts it to TFiDF (term frequency, inverse document frequency) encoded data. 
This transformed data has one column per word in the text messages and each column has a floating point value indicating its weight.


To train run 
```
$ python train.py
/Users/magnus/anaconda/lib/python2.7/site-packages/numpy/lib/npyio.py:1769: ConversionWarning: Some errors were detected !
    Line #15 (got 3 columns instead of 2)
    Line #102 (got 3 columns instead of 2)
    Line #107 (got 3 columns instead of 2)
    Line #166 (got 3 columns instead of 2)
    Line #212 (got 3 columns instead of 2)
    Line #269 (got 3 columns instead of 2)
    Line #288 (got 3 columns instead of 2)
    Line #294 (got 3 columns instead of 2)
    Line #370 (got 3 columns instead of 2)
    Line #432 (got 3 columns instead of 2)
    Line #442 (got 3 columns instead of 2)
    Line #465 (got 3 columns instead of 2)
    Line #476 (got 3 columns instead of 2)
    Line #485 (got 3 columns instead of 2)
    Line #584 (got 3 columns instead of 2)
    Line #870 (got 3 columns instead of 2)
    Line #871 (got 3 columns instead of 2)
    Line #917 (got 3 columns instead of 2)
    Line #1038 (got 3 columns instead of 2)
    Line #1287 (got 3 columns instead of 2)
  warnings.warn(errmsg, ConversionWarning)


No instance found at ip and port: localhost:54321. Trying to start local jar...


JVM stdout: /var/folders/fq/lm9wfzqj5ksg4tkhm1ml1fcc0000gn/T/tmpXxAIAM/h2o_magnus_started_from_python.out
JVM stderr: /var/folders/fq/lm9wfzqj5ksg4tkhm1ml1fcc0000gn/T/tmpwqS6Lv/h2o_magnus_started_from_python.err
Using ice_root: /var/folders/fq/lm9wfzqj5ksg4tkhm1ml1fcc0000gn/T/tmpZmo8oL


Java Version: java version "1.8.0_77"
Java(TM) SE Runtime Environment (build 1.8.0_77-b03)
Java HotSpot(TM) 64-Bit Server VM (build 25.77-b03, mixed mode)


Starting H2O JVM and connecting: ............ Connection successful!
------------------------------  -------------------------------------
H2O cluster uptime:             1 seconds 388 milliseconds
H2O cluster version:            3.8.2.8
H2O cluster name:               H2O_started_from_python_magnus_rpr813
H2O cluster total nodes:        1
H2O cluster total free memory:  3.56 GB
H2O cluster total cores:        8
H2O cluster allowed cores:      8
H2O cluster healthy:            True
H2O Connection ip:              127.0.0.1
H2O Connection port:            54321
H2O Connection proxy:
Python Version:                 2.7.11
------------------------------  -------------------------------------

Parse Progress: [##################################################] 100%

gbm Model Build Progress: [##################################################] 100%
Filepath: ./lib//GBM_model_python_1470258640703_1.java
Are you sure you want to shutdown the H2O instance running at localhost:54321 (Y/N)? y
$
```

You now have a couple of new files in the ```lib``` directory. A POJO file ```GBM_model_python_1470258640703_1.java```, 
```h2o-genmodel.jar```, a companion library, and ```vectorizer.pickle``` which is the preprocessing code.

## Build the web service
 
The Prediction Service Builder must be running on port 55000. A Java JDK must be installed.
 
Build the example web service with 
```
$ sh example-python.sh 
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100 1861k  100 1447k  100  413k   947k   270k  0:00:01  0:00:01 --:--:--  947k
Created example-python.war
Run with run-example-pyhton.sh
$
```
 which creates the service as ```example-python.war```
 
 An alternative to this is to build the service using the web UI at ```http://localhost:55000```. 
 You then select the files that are needed in the UI. 
 Select ```GBM_model_python_1463864606917_1.java``` and ```h2o-genmodel.jar``` in the Java part.
 Select ```score.py``` as the Python file and as additional files, select: 
 ```vectorizer.pickle```, ```lib/modelling.py```, and ```lib/__init__.py``` .
 
 **NOTE** You may get errors if Java JDK is not installed on the machine, in which case
 you want to install that.
 
## Run the web service

Python must be installed on the machine. Numpy, Sklearn and textblob must also be installed. 

The service is run with
```
$ sh run-example-python.sh
 starting prediction service on port 55001
 2016-08-03 14:16:38.157:INFO:omjr.Runner:Runner
 2016-08-03 14:16:38.157:WARN:omjr.Runner:No tx manager found
 2016-08-03 14:16:38.190:INFO:omjr.Runner:Deploying file:/Users/magnus/Git/steam/prediction-service-builder/examples/spam-detection-python/example-python.war @ /
 2016-08-03 14:16:38.214:INFO:oejs.Server:jetty-8.y.z-SNAPSHOT
 2016-08-03 14:16:38.235:INFO:oejw.WebInfConfiguration:Extract jar:file:/Users/magnus/Git/steam/prediction-service-builder/examples/spam-detection-python/example-python.war!/ to /private/var/folders/fq/lm9wfzqj5ksg4tkhm1ml1fcc0000gn/T/jetty-0.0.0.0-55001-example-python.war-_-any-/webapp
 2016-08-03 14:16:38.453:INFO:oejpw.PlusConfiguration:No Transaction manager found - if your webapp requires one, please configure one.
 2016-08-03 14:16:39.019 -0700 [main] INFO PredictPythonServlet - Python started
 2016-08-03 14:16:39.019:WARN:oejsh.RequestLogHandler:!RequestLog
 2016-08-03 14:16:39.039:INFO:oejs.AbstractConnector:Started SelectChannelConnector@0.0.0.0:55001
 ^C2016-08-03 14:20:33.501 -0700 [Thread-0] INFO PredictPythonServlet - Python destroyed
 2016-08-03 14:20:33.505:INFO:oejsl.ELContextCleaner:javax.el.BeanELResolver purged
 2016-08-03 14:20:33.505:INFO:oejsh.ContextHandler:stopped o.e.j.w.WebAppContext{/,file:/private/var/folders/fq/lm9wfzqj5ksg4tkhm1ml1fcc0000gn/T/jetty-0.0.0.0-55001-example-python.war-_-any-/webapp/},file:/Users/magnus/Git/steam/prediction-service-builder/examples/spam-detection-python/example-python.war
```
It is now started at ```http://localhost:55001```
and looks like this


![Builder Service](images/spam-detection-python localhost 55001.png)


**NOTE** You may get Python errors if you don't have the correct Python packages install.
A good way to test this is to run ```score.py```. If it works fine it looks like this. 
```
$ python score.py
python ready

```
If not, it will tell you what Python packages you need to install. Typically you need 
```textblob``` which is installed with ```pip install textblob```.
