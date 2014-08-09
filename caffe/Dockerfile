
FROM ubuntu:14.04

# Get latest ubuntu packages
RUN apt-get update

# Caffe requires BLAS as the backend of its matrix and vector computations
RUN apt-get -q -y install libatlas-base-dev

# Clone the Caffe repo 
RUN apt-get -q -y install git
RUN cd /opt && git clone https://github.com/BVLC/caffe.git

# Install python + python dependencies
RUN apt-get -q -y install python python-dev python-pip
RUN apt-get -q -y install pkg-config
RUN apt-get -q -y install libfreetype6-dev 
RUN apt-get -q -y install libgfortran3 gfortran
RUN pip install -r /opt/caffe/python/requirements.txt

# Add Caffe to python path
ENV PYTHONPATH /opt/caffe/python:$PYTHONPATH

# Other dependencies
RUN apt-get -q -y install libprotobuf-dev libleveldb-dev libsnappy-dev libopencv-dev libboost-all-dev libhdf5-serial-dev


