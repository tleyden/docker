
FROM ubuntu:14.04

# Get latest ubuntu packages
RUN apt-get update

# Caffe requires BLAS as the backend of its matrix and vector computations
RUN apt-get -q -y install libatlas-base-dev

# Clone the Caffe repo 
RUN apt-get -q -y install git wget curl unzip
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
RUN apt-get -q -y install libprotobuf-dev libleveldb-dev libsnappy-dev libopencv-dev libboost-all-dev libhdf5-serial-dev liblmdb-dev protobuf-compiler

# Glog 
RUN cd /opt && wget https://google-glog.googlecode.com/files/glog-0.3.3.tar.gz
RUN cd /opt && tar zxvf glog-0.3.3.tar.gz 
RUN cd /opt/glog-0.3.3 && ./configure && make && make install

# Gflags
RUN apt-get -q -y install cmake
RUN cd /opt && wget https://github.com/schuhschuh/gflags/archive/master.zip
RUN cd /opt && unzip master.zip
RUN cd /opt/gflags-master && mkdir build 
RUN cd /opt/gflags-master/build && export CXXFLAGS="-fPIC" && cmake .. && make VERBOSE=1 && make && make install

# Compile core
RUN cd /opt/caffe && cp Makefile.config.example Makefile.config
RUN cd /opt/caffe && echo "CPU_ONLY := 1" >> Makefile.config
RUN cd /opt/caffe && make all && make test && make runtest

# Compile python wrappers
# TODO: edit Python paths in Makefile.config (needed?)
RUN cd /opt/caffe && make pycaffe



