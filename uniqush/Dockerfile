
FROM ubuntu:14.04

RUN apt-get update

RUN apt-get -q -y install redis-server
RUN apt-get -q -y install wget
RUN cd /tmp && wget http://uniqush.org/downloads/uniqush-push_1.5.2_amd64.deb
RUN cd /tmp && dpkg -i uniqush-push_1.5.2_amd64.deb
RUN /etc/init.d/redis-server start

COMMAND uniqush-push

