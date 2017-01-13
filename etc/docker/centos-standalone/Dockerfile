FROM centos:6.8

MAINTAINER H2O.ai version: 1.2.0

RUN mkdir -p /go/src/github.com/h2oai/steam
ADD . /go/src/github.com/h2oai/steam
WORKDIR /go/src/github.com/h2oai/steam

# Install NodeJS
RUN curl --silent --location https://rpm.nodesource.com/setup_6.x | bash -
RUN yum install -y nodejs-6.8.1

# Install GCC for SQLite dependency
RUN yum install -y gcc

# Install Go for Steam backend
RUN curl -o go.tar.gz https://storage.googleapis.com/golang/go1.7.3.linux-amd64.tar.gz
RUN tar -C /usr/local -xzf go.tar.gz
ENV PATH $PATH:/usr/local/go/bin
ENV GOPATH /steam

# Install Java for Prediction Service
RUN yum install -y java-1.7.0-openjdk
RUN yum install -y java-1.7.0-openjdk-devel

# Install Typings
WORKDIR /go/src/github.com/h2oai/steam/gui
RUN npm install typings -g
RUN typings install

# Build Steam
WORKDIR /go/src/github.com/h2oai/steam
RUN ls
RUN make

# Run Prediction Service Builder
RUN nohup java -jar var/master/assets/jetty-runner.jar var/master/assets/ROOT.war &

EXPOSE 8080
EXPOSE 9000
EXPOSE 9001

# Run Steam
CMD ./steam serve master --admin-name root --admin-password admin012

