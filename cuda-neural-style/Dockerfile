FROM kaixhin/cuda-torch:latest
MAINTAINER Traun Leyden "traun.leyden@gmail.com"

RUN apt-get install -y wget libpng-dev libprotobuf-dev protobuf-compiler
RUN git clone --depth 1 https://github.com/jcjohnson/neural-style.git
RUN /root/torch/install/bin/luarocks install loadcaffe

WORKDIR /root/torch/neural-style
VOLUME /root/torch/neural-style/models
VOLUME /root/torch/neural-style/images
VOLUME /root/torch/neural-style/outputs