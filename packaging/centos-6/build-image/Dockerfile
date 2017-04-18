FROM alanfranz/fwd-centos-6:latest
MAINTAINER H2O.ai version: 1.2.0

RUN mkdir -p /go/src/github.com/h2oai/steam
WORKDIR /go/src/github.com/h2oai/steam

RUN curl --silent --location https://rpm.nodesource.com/setup_6.x | bash - && yum install -y nodejs-6.8.1

RUN yum install -y gcc

RUN curl -o go.tar.gz https://storage.googleapis.com/golang/go1.7.3.linux-amd64.tar.gz && tar -C /usr/local -xzf go.tar.gz
ENV PATH $PATH:/usr/local/go/bin
ENV GOPATH /go

RUN yum install -y java-1.7.0-openjdk-devel
ENV JAVA_HOME /usr/lib/jvm/java

