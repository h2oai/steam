FROM ubuntu:precise
MAINTAINER H2O.ai version: 1.2.0

ENV DEBIAN_FRONTEND noninteractive

RUN mkdir -p /go/src/github.com/h2oai/steam
WORKDIR /go/src/github.com/h2oai/steam

RUN \
    apt-get -q -y update && \
    apt-get -y install software-properties-common && \
    apt-get -y install less curl sudo rsync && \
    apt-get clean 
            
RUN \
    curl --silent --location https://deb.nodesource.com/setup_6.x | bash - && \
    apt-get -y install nodejs && \
    npm install typings -g

RUN \
    curl -o go.tar.gz https://storage.googleapis.com/golang/go1.7.3.linux-amd64.tar.gz && \
    tar -C /usr/local -xzf go.tar.gz
ENV PATH $PATH:/usr/local/go/bin
ENV GOPATH /go

RUN \
    apt-get install -y --no-install-recommends openjdk-7-jdk && \
    apt-get clean

RUN \
    apt-get install -y --no-install-recommends \
        g++ \
        gcc \
        libc6-dev \
        make \
        pkg-config \
    && apt-get clean \
    && rm -rf /var/cache/apt/* \
    && rm -rf /var/lib/apt/lists/*
