FROM openjdk:jre-alpine

# docker build -f Dockerfile-service -t predservice .
# docker run --rm -p 55001:55001 -v $PWD:/host predservice example.war

LABEL maintainer="H2o.ai <ops@h2o.ai>"

COPY jetty-runner-8.1.14.v20131031.jar /

WORKDIR "/host"

ENTRYPOINT [ "java", "-jar", "jetty-runner-8.1.14.v20131031.jar", "--port", "55001" ]

EXPOSE 55001
