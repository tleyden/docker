FROM centos:6

MAINTAINER Traun Leyden <tleyden@couchbase.com>

ENV GOPATH /opt/go
ENV GOROOT /usr/local/go

ENV PATH $PATH:$GOPATH/bin:$GOROOT/bin

# Get dependencies
RUN yum -y update && \
  yum install -y \
  curl \
  git \
  mercurial \
  tar \
  wget

# Download and install Go 1.4
RUN wget http://golang.org/dl/go1.4.1.linux-amd64.tar.gz && \
    tar -C /usr/local -xzf go1.4.1.linux-amd64.tar.gz && \
    rm go1.4.1.linux-amd64.tar.gz

# Install cbgb
RUN go get -u -v -t github.com/couchbaselabs/cbgb/...
