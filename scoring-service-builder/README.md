# H2O Scoring Service Builder
This is a service that can
1. Compile the Pojo and build a Jar file from a Pojo and a gen-model file
2. Compile the Pojo and build a War file that is a service from a Pojo, gen-model and
another Jar file. You can run the war file in Jetty, Tomcat, etc.
## How to Build

`./gradlew build
`
## War web service
`./gradlew jettyRunWar
`

starts a web server on localhost:55000

This web server can build a WAR file from an H2O Pojo,
h2o-genmodel.jar file and a an extra jar file. Example files are in
directory example-pojo.

`curl -X POST --form pojo=@example-pojo/gbm_3f258f27_f0ad_4520_b6a5_3d2bb4a9b0ff.java --form jar=@example-pojo/h2o-genmodel.jar --form extra=@makewar-extra.jar localhost:8080/makewar > gbm.war
`

gbm_3f258f27_f0ad_4520_b6a5_3d2bb4a9b0ff.java is the Pojo from
H2O. h2o-genmodel.jar is the corresponding jar file from the version
of H2O. makewar-extra.jar is a jar file that contains additional files
needed to build the war file.

If you go to http://localhost:55000 you get a web page with a form that
can be used to create war files.

## Prediction server

The result of the service above is a war file that can be run with

`java -jar jetty-runner-9.3.9.M1.jar --port 8081 gbm.war
`

This in turn starts a web service at localhost:8081 . 

This service can be used for prediction

`curl "localhost:8081/predict?Dest=JFK"`

which returns a JSON result

`{"labelIndex":1,"label":"Y","classProbabilities":[0.026513747179178093,0.9734862528208219]}
`

The predictor has two classes. "Y" was predicted with probability 0.97.

There is also a web page for the predictor at http://localhost:8081 .

## Jetty runner versions

If you use an older Java version, you need to use an older
jetty-runner. Jetty 9.3 requires Java 1.8. Jetty 9.0-9.2 requires Java
1.7. Jetty 8 requires Java 1.6. 

Testing has been done on Java 1.8 and 1.7. For Java 1.8 you can use
all jetty runners while on Java 1.7 you can use all except the 9.3
version.

## To Do

The web pages provided are just examples of what can be done and need
to be improved to look good.

We could also improve the handling of the parameters of the predictor
by getting the right types and variable names from the predictor.
