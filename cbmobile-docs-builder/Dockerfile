FROM centos:centos7

MAINTAINER Traun Leyden <tleyden@couchbase.com>

# Install yum dependencies
RUN yum -y update && \
    yum install -y \
    git \
    java-1.7.0-openjdk \
    sudo \
    tar \
    wget 

# Install setuptools + pip
RUN cd /tmp && \
    wget --no-check-certificate https://pypi.python.org/packages/source/s/setuptools/setuptools-1.4.2.tar.gz && \
    tar -xvf setuptools-1.4.2.tar.gz && \
    cd setuptools-1.4.2 && \
    python2.7 setup.py install && \
    curl https://raw.githubusercontent.com/pypa/pip/master/contrib/get-pip.py | python2.7 - && \
    pip install virtualenv

# Install aws cli
RUN sudo pip install awscli

# Install s3cmd (mime type detection doesn't seem to work too well)
RUN cd /tmp/ && \
    wget https://github.com/s3tools/s3cmd/releases/download/v1.5.2/s3cmd-1.5.2.tar.gz && \
    tar xvfz s3cmd-1.5.2.tar.gz && \
    cd s3cmd-1.5.2 && \
    python setup.py install

