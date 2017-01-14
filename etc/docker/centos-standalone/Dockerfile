FROM alanfranz/fwd-centos-6:latest

MAINTAINER H2O.ai version: 1.1.6

WORKDIR /steam

# Install NodeJS
RUN curl --silent --location https://rpm.nodesource.com/setup_6.x | bash -
RUN yum install -y nodejs-6.8.1

# Add jenkins and SSH
RUN mkdir /var/run/sshd
RUN adduser jenkins
RUN yum install -y openssh-server
RUN /usr/sbin/sshd

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

# Clone Steam repository
RUN yum install -y git
RUN mkdir -p src/github.com/h2oai
WORKDIR /steam/src/github.com/h2oai
RUN git clone https://github.com/h2oai/steam.git

# Install Typings
WORKDIR /steam/src/github.com/h2oai/steam/gui
RUN npm install typings -g
RUN typings install

# Build Steam
WORKDIR /steam/src/github.com/h2oai/steam
RUN make
RUN make db

# Run Prediction Service Builder
RUN nohup java -jar var/master/assets/jetty-runner.jar var/master/assets/ROOT.war &

EXPOSE 8080
EXPOSE 9000
EXPOSE 9001

# Run Steam
CMD ./steam serve master --superuser-name root --superuser-password superuser

