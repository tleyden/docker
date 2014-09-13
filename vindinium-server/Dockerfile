FROM ubuntu:14.04

RUN apt-get update

RUN apt-get -q -y install git wget mongodb
RUN apt-get -q -y install --no-install-recommends openjdk-7-jdk

RUN mkdir /opt

RUN cd /opt && wget http://dl.bintray.com/sbt/debian/sbt-0.13.5.deb

# Messy .. but this is what I actually ended up doing

RUN cd /opt && dpkg -i sbt-0.13.5.deb
RUN apt-get -q -y -f install
RUN cd /opt && dpkg -i sbt-0.13.5.deb

RUN cd /opt && git clone https://github.com/ornicar/vindinium.git

RUN cd /opt/vindinium && sbt compile


# Trying to get "TV" running

RUN apt-get -q -y install npm
RUN cd /opt/vindinium/client && npm install
RUN cd /opt/vindinium/client && npm install -g grunt-cli

RUN ln -s /usr/bin/nodejs /usr/bin/node
RUN cd /opt/vindinium/client && grunt





 




