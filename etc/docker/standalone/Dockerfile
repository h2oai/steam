FROM ubuntu:latest

MAINTAINER H2O.ai version: 1.1.6

WORKDIR /steam

RUN apt-get update
RUN apt-get install -y curl sqlite3 libsqlite3-dev

# Install NodeJS
RUN curl -sL https://deb.nodesource.com/setup_6.x | bash -
RUN apt-get install -y nodejs build-essential

# Clone Steam repository
RUN apt-get install -y git
RUN mkdir -p src/github.com/h2oai
WORKDIR /steam/src/github.com/h2oai
RUN git clone https://github.com/h2oai/steam.git

# Install Typings
WORKDIR /steam/src/github.com/h2oai/steam/gui
RUN npm install typings -g
RUN typings install

# Install Go for Steam backend
RUN curl -o go.tar.gz https://storage.googleapis.com/golang/go1.7.3.linux-amd64.tar.gz
RUN tar -C /usr/local -xzf go.tar.gz
ENV PATH $PATH:/usr/local/go/bin
ENV GOPATH /steam

# Install Java for Prediction Service
RUN apt-get install -y software-properties-common
RUN add-apt-repository ppa:openjdk-r/ppa
RUN apt-get update
RUN apt-get install -y openjdk-7-jdk

# Build Steam
WORKDIR /steam/src/github.com/h2oai/steam
RUN make
RUN make db

# Run Prediction Service Builder
RUN nohup java -jar var/master/assets/jetty-runner.jar var/master/assets/ROOT.war &

EXPOSE 8080
EXPOSE 9000
EXPOSE 9001

CMD ./steam serve master --superuser-name docker --superuser-password superuser