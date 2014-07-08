
FROM ubuntu:14.04

MAINTAINER Traun Leyden tleyden@couchbase.com

ENV GOPATH /opt/go
ENV SGROOT /opt/sync_gateway

RUN apt-get update

RUN apt-get -q -y install git bc
RUN apt-get -q -y install golang

RUN mkdir -p $GOPATH

RUN mkdir -p /opt

RUN cd /opt && git clone https://github.com/couchbase/sync_gateway.git
RUN cd $SGROOT && git submodule init && git submodule update

RUN cd $SGROOT && ./build.sh

RUN cd $SGROOT && cp bin/sync_gateway /usr/local/bin

RUN mkdir -p $SGROOT/data
