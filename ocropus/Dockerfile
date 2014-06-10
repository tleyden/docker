FROM ubuntu:14.04

ENV OCROPUSPATH /opt/ocropus/ocropy

RUN apt-get update
RUN apt-get -q -y install mercurial
RUN cd /opt && hg clone -r ocropus-0.7 https://code.google.com/p/ocropus
RUN cd $OCROPUSPATH && sudo apt-get -q -y install $(cat PACKAGES)
RUN cd $OCROPUSPATH && python setup.py download_models
RUN cd $OCROPUSPATH && sudo python setup.py install