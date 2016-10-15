FROM sequenceiq/hadoop-ubuntu:2.6.0
MAINTAINER H2O.ai

USER root

<<<<<<< HEAD
ENV STEAM_VERSION 1.0.0
=======
ENV STEAM_VERSION 1.1.2
>>>>>>> 40db9e0884586ad6edacaf171aa766624772088e
ENV PATH /steam/steam-${STEAM_VERSION}-linux-amd64:$PATH
ENV PATH $PATH:/usr/local/hadoop/bin

# Fix DNS resolution issues when nss is not installed
RUN echo 'hosts: files mdns4_minimal [NOTFOUND=return] dns mdns4' >> /etc/nsswitch.conf

WORKDIR /steam
ADD http://s3.amazonaws.com/steam-release/steam-${STEAM_VERSION}-linux-amd64.tar.gz /steam/
RUN tar xvf steam-${STEAM_VERSION}-linux-amd64.tar.gz
WORKDIR /steam/steam-${STEAM_VERSION}-linux-amd64
ADD start.sh start.sh
RUN ls -la

EXPOSE 9002
EXPOSE 9001

CMD start.sh
