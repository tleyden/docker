

FROM ubuntu:14.04

# Get latest ubuntu packages
RUN apt-get update

# Dependencies
RUN apt-get install -y build-essential gfortran g++ libopenblas-dev git curl

# Python
RUN apt-get install -y python python-numpy python-scipy python-dev python-pip python-nose python-yaml python-imaging python-matplotlib

# install Theano (development version)
RUN cd /opt && git clone git://github.com/Theano/Theano.git
RUN cd /opt/Theano && python setup.py develop

# install Pylearn2 (development version)
RUN cd /opt && git clone git://github.com/lisa-lab/pylearn2.git
RUN cd /opt/pylearn2 && python setup.py develop

# set Data path
RUN bash -c 'echo "export PYLEARN2_DATA_PATH=/opt/data" >> /.bashrc'

# set Scripts path
RUN bash -c 'echo "export PATH=/opt/pylearn2/pylearn2/scripts:$PATH" >> /.bashrc'