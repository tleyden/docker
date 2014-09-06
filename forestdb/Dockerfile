FROM ubuntu:14.04

RUN apt-get update

RUN apt-get -q -y install git libsnappy-dev cmake build-essential

# NOTE: this requires username and password (until repo is made public)
RUN cd /opt && git clone https://github.com/couchbaselabs/forestdb

# Build it
RUN cd /opt/forestdb && mkdir build
RUN cd /opt/forestdb/build && cmake ../
RUN cd /opt/forestdb/build && make all && make install

# Install header files
RUN mkdir /usr/local/include/forestdb
RUN cd /opt/forestdb && cp include/libforestdb/* /usr/local/include/forestdb

# Install go
ENV GOPATH /opt/go
RUN apt-get -q -y install golang
RUN mkdir -p $GOPATH

# Without this, it won't find the /usr/local/lib/libforestdb.so
RUN ldconfig

# Get go-forestdb
RUN go get -u -v -t github.com/couchbaselabs/goforestdb


