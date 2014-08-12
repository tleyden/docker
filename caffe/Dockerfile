
FROM ubuntu:14.04

# Get latest ubuntu packages
RUN apt-get update

# Get dependencies
RUN apt-get install -y libprotobuf-dev libleveldb-dev libsnappy-dev libopencv-dev libboost-all-dev libhdf5-serial-dev protobuf-compiler gcc-4.6 g++-4.6 gcc-4.6-multilib g++-4.6-multilib gfortran libjpeg62 libfreeimage-dev libatlas-base-dev git python-dev python-pip bc wget curl unzip cmake liblmdb-dev pkgconf

# Use gcc 4.6
RUN update-alternatives --install /usr/bin/cc cc /usr/bin/gcc-4.6 30 && update-alternatives --install /usr/bin/c++ c++ /usr/bin/g++-4.6 30 && update-alternatives --install /usr/bin/gcc gcc /usr/bin/gcc-4.6 30 && update-alternatives --install /usr/bin/g++ g++ /usr/bin/g++-4.6 30

# Clone the Caffe repo 
RUN cd /opt && git clone https://github.com/BVLC/caffe.git

# Glog 
RUN cd /opt && wget https://google-glog.googlecode.com/files/glog-0.3.3.tar.gz
RUN cd /opt && tar zxvf glog-0.3.3.tar.gz 
RUN cd /opt/glog-0.3.3 && ./configure && make && make install

# Gflags
RUN cd /opt && wget https://github.com/schuhschuh/gflags/archive/master.zip
RUN cd /opt && unzip master.zip
RUN cd /opt/gflags-master && mkdir build 
RUN cd /opt/gflags-master/build && export CXXFLAGS="-fPIC" && cmake .. && make VERBOSE=1 && make && make install

# Build Caffe core
RUN cd /opt/caffe && cp Makefile.config.example Makefile.config
RUN cd /opt/caffe && echo "CPU_ONLY := 1" >> Makefile.config 
RUN cd /opt/caffe && echo "CXX := /usr/bin/g++-4.6" >> Makefile.config 
RUN cd /opt/caffe && sed -i 's/CXX :=/CXX ?=/' Makefile
RUN cd /opt/caffe && make all

# Install python deps
RUN cd /opt/caffe && (pip install -r python/requirements.txt; easy_install numpy; pip install -r python/requirements.txt)
RUN easy_install pillow

# Numpy include path hack - github.com/BVLC/caffe/wiki/Setting-up-Caffe-on-Ubuntu-14.04
RUN ln -s /usr/local/lib/python2.7/dist-packages/numpy-1.8.2-py2.7-linux-x86_64.egg/numpy/core/include/numpy /usr/include/python2.7/numpy

# Build Caffe python bindings
RUN cd /opt/caffe && make pycaffe

# Run tests
RUN cd /opt/caffe && make test && make runtest