FROM ubuntu:latest

MAINTAINER H2O.ai version: 1.1.0

ENV STEAM_VERSION 1.1.0
ENV PATH /steam/steam-${STEAM_VERSION}-linux-amd64:$PATH

WORKDIR /steam
ADD http://s3.amazonaws.com/steam-release/steam-${STEAM_VERSION}-linux-amd64.tar.gz /steam/
RUN tar xvf steam-${STEAM_VERSION}-linux-amd64.tar.gz
WORKDIR /steam/steam-${STEAM_VERSION}-linux-amd64
RUN ls

EXPOSE 9000
EXPOSE 9001

CMD ["steam", "serve", "master","--superuser-name=docker","--superuser-password=superuser"]