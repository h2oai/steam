FROM sequenceiq/hadoop-ubuntu:2.6.0
MAINTAINER H2O.ai

USER root

ENV STEAM_VERSION 1.2.0
ENV PATH /go/src/github.com/h2oai/steam:$PATH
ENV PATH $PATH:/usr/local/hadoop/bin

# Fix DNS resolution issues when nss is not installed
RUN echo 'hosts: files mdns4_minimal [NOTFOUND=return] dns mdns4' >> /etc/nsswitch.conf

RUN mkdir -p /go/src/github.com/h2oai/steam
ADD . /go/src/github.com/h2oai/steam
WORKDIR /go/src/github.com/h2oai/steam
ADD etc/docker/hadoop/start.sh start.sh

RUN apt-get update
RUN apt-get install -y curl sqlite3 libsqlite3-dev

# Install NodeJS
RUN curl -sL https://deb.nodesource.com/setup_6.x | bash -
RUN apt-get install -y nodejs build-essential

# Install Typings
WORKDIR /go/src/github.com/h2oai/steam/gui
RUN npm install typings -g
RUN typings install

# Install Go for Steam backend
RUN curl -o go.tar.gz https://storage.googleapis.com/golang/go1.7.3.linux-amd64.tar.gz
RUN tar -C /usr/local -xzf go.tar.gz
ENV PATH $PATH:/usr/local/go/bin
ENV GOPATH /go

# Install Java for Prediction Service
RUN apt-get install -y software-properties-common
RUN add-apt-repository ppa:openjdk-r/ppa
RUN apt-get update
RUN apt-get install -y openjdk-7-jdk

# Install HAProxy
RUN apt-get install software-properties-common
RUN add-apt-repository ppa:vbernat/haproxy-1.7
RUN apt-get update
RUN apt-get install -y haproxy

# Build Steam
WORKDIR /go/src/github.com/h2oai/steam
RUN make clean reset
RUN make

# Create self-signed cert
#RUN openssl req -nodes -x509 -newkey rsa:4096 -keyout cert.pem -out cert.crt -days 365 -subj '/CN=localhost'
RUN openssl genrsa -out steam.key 2048
RUN openssl req -new -key steam.key -out steam.csr -subj "/C=US/ST=California/L=Mountain View/O=IT/CN=www.h2o.ai"
RUN sudo openssl x509 -req -days 365 -in steam.csr -signkey steam.key -out steam.crt
RUN cat steam.crt steam.key | tee steam.pem
RUN cp steam.pem /go/src/github.com/h2oai/steam/var/master/assets/

# Run Prediction Service Builder
RUN nohup java -jar var/master/assets/jetty-runner.jar var/master/assets/ROOT.war &

EXPOSE 8080
EXPOSE 9000
EXPOSE 9001
EXPOSE 9002
EXPOSE 9999

CMD /go/src/github.com/h2oai/steam/start.sh