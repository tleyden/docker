FROM ubuntu:14.04

MAINTAINER Traun Leyden <tleyden@couchbase.com>

ENV GOPATH /opt/go
ENV GOROOT /usr/local/go
ENV PATH $PATH:$GOPATH/bin:$GOROOT/bin

# Get dependencies
RUN apt-get update && apt-get install -y \
  bc \
  build-essential \
  curl \
  emacs24-nox \
  git \
  mercurial \
  wget && \
  apt-get clean

# Install Go
RUN wget http://golang.org/dl/go1.4.2.linux-amd64.tar.gz && \
    tar -C /usr/local -xzf go1.4.2.linux-amd64.tar.gz && \
    rm go1.4.2.linux-amd64.tar.gz

# Install go packages
RUN go get golang.org/x/tools/cmd/... && \
    go get github.com/tools/godep && \
    go get -u -v github.com/nsf/gocode && \
    go get -u -v golang.org/x/tools/cmd/goimports && \
    go get -u -v github.com/golang/lint/golint && \
    go get -u -v github.com/rogpeppe/godef && \
    go get -u -v github.com/smartystreets/goconvey

# clone emacs conf
RUN git clone https://github.com/fgimenez/.emacs.d.git /root/.emacs.d && \
    cd /root/.emacs.d && \
    git checkout origin/go


