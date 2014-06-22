
FROM ubuntu:14.04

RUN apt-get -q -y update

RUN apt-get -q -y install wget
RUN apt-get -q -y install unzip
RUN apt-get -q -y install git

RUN apt-get -q -y install libopencv-core2.4
RUN apt-get -q -y install libopencv-core-dev
RUN apt-get -q -y install libboost1.55-all-dev
RUN apt-get -q -y install libopencv-flann2.4 libopencv-flann-dev
RUN apt-get -q -y install libopencv-imgproc2.4 libopencv-imgproc-dev
RUN apt-get -q -y install libopencv-photo2.4 libopencv-photo-dev
RUN apt-get -q -y install libopencv-video2.4 libopencv-video-dev
RUN apt-get -q -y install libopencv-features2d2.4 libopencv-features2d-dev
RUN apt-get -q -y install libopencv-objdetect2.4 libopencv-objdetect-dev
RUN apt-get -q -y install libopencv-calib3d2.4 libopencv-calib3d-dev
RUN apt-get -q -y install libopencv-ml2.4 libopencv-ml-dev
RUN apt-get -q -y install libopencv-contrib2.4 libopencv-contrib-dev
RUN apt-get -q -y install libopencv-highgui2.4 libopencv-highgui-dev

RUN mkdir -p /opt
RUN cd /opt && git clone https://github.com/tleyden/DetectText.git
RUN cd /opt/DetectText && g++ -o DetectText TextDetection.cpp FeaturesMain.cpp -lopencv_core -lopencv_highgui -lopencv_imgproc -I/opt/DetectText

RUN cd /opt/DetectText && cp DetectText /usr/local/bin



