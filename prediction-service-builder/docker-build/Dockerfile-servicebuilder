FROM openjdk:jdk-alpine

# docker build -f Dockerfile-servicebuilder -t predservicebuilder .
# docker run --rm -p 55000:55000 predservicevbuilder

LABEL maintainer="H2o.ai <ops@h2o.ai>"

COPY jetty-runner-8.1.14.v20131031.jar ROOT.war /

ENTRYPOINT [ "java", "-jar", "jetty-runner-8.1.14.v20131031.jar", "--port", "55000", "/ROOT.war" ]

EXPOSE 55000

