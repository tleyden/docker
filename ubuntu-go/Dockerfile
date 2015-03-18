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
  emacs \
  git \
  mercurial \
  wget && \
  apt-get clean

# clone go source
RUN cd /usr/local && \
    git clone https://go.googlesource.com/go && \
    git checkout go1.4.2 

# build it
RUN cd /usr/local/go/src; \
    ./all.bash; \
    echo "Done building"

# install go tools
RUN go get golang.org/x/tools/cmd/...

# install go packages needed for emacs config
RUN go get github.com/tools/godep && \
    go get github.com/nsf/gocode && \
    go get code.google.com/p/go.tools/cmd/goimports && \
    go get github.com/golang/lint/golint && \
    go get code.google.com/p/rog-go/exp/cmd/godef

# clone emacs conf
RUN git clone https://github.com/fgimenez/.emacs.d.git /root/.emacs.d && \
    cd /root/.emacs.d && \
    git checkout origin/go


