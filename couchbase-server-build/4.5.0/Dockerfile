FROM ubuntu:trusty

# Based on instructions from https://github.com/couchbase/tlm
# It turns out this container re-invents this wheel: https://github.com/couchbase/build/blob/master/docker/buildslaves/ubuntu-16.04/jenkins/Dockerfile

# Install packages
RUN apt-get update && \
    apt-get install -y \
      build-essential \
      cmake \
      curl \
      git \
      libssl-dev \
      openssl \
      python \
      wget && \
    apt-get clean

# Download and install Go
RUN wget http://golang.org/dl/go1.5.linux-amd64.tar.gz && \
    tar -C /usr/local -xzf go1.5.linux-amd64.tar.gz && \
    rm go1.5.linux-amd64.tar.gz 

# Add Go to the PATH
ENV PATH $PATH:/usr/local/go/bin

# Configure git
RUN git config --global user.email "you@example.com" && \
    git config --global user.name "Your Name" && \
    git config --global color.ui false 

# Install repo
RUN curl https://storage.googleapis.com/git-repo-downloads/repo > /usr/bin/repo && \
    chmod a+x /usr/bin/repo 

# Create source dir
RUN mkdir -p /opt/couchbase-build/source

# Set working dir to source dir
WORKDIR /opt/couchbase-build/source

# Checkout and build code
RUN repo init -u git://github.com/couchbase/manifest -m released/4.0.0.xml && \
    repo sync

# Build
RUN make


