FROM ubuntu:14.04

MAINTAINER Traun Leyden <tleyden@couchbase.com>

ENV LC_ALL C.UTF-8

# Get dependencies
RUN apt-get update && apt-get install -y \
  bundler \
  curl \
  emacs24-nox \
  git \
  ruby \
  wget && \
  apt-get clean

# Install octopress
RUN cd /root && git clone git://github.com/imathis/octopress.git && \
    cd octopress && \
    gem install bundler && \
    bundle install



