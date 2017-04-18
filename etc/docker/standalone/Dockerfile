FROM ubuntu:latest

MAINTAINER H2O.ai version: 1.2.0

RUN mkdir -p /go/src/github.com/h2oai/steam
ADD . /go/src/github.com/h2oai/steam
WORKDIR /go/src/github.com/h2oai/steam

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

# Build Steam
WORKDIR /go/src/github.com/h2oai/steam
RUN make clean reset
RUN make

# Run Prediction Service Builder
RUN nohup java -jar var/master/assets/jetty-runner.jar var/master/assets/ROOT.war &

EXPOSE 8080
EXPOSE 9000
EXPOSE 9001
EXPOSE 9999

CMD ./steam serve master --admin-name docker --admin-password admin012