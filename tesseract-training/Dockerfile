FROM ubuntu:14.04

RUN apt-get update

RUN apt-get -q -y install subversion emacs24-nox wget
RUN apt-get -q -y install liblept4 libleptonica-dev
RUN apt-get -q -y install autoconf automake libtool
RUN apt-get -q -y install libpng12-dev
RUN apt-get -q -y install libjpeg62-dev
RUN apt-get -q -y install libtiff4-dev
RUN apt-get -q -y install zlib1g-dev	
RUN apt-get -q -y install libicu-dev
RUN apt-get -q -y install libpango1.0-dev
RUN apt-get -q -y install libcairo2-dev


RUN cd /opt && svn checkout http://tesseract-ocr.googlecode.com/svn/trunk/ tesseract-ocr-read-only

RUN cd /opt/tesseract-ocr-read-only && ./autogen.sh && ./configure && make && sudo make install && sudo ldconfig 

RUN cd /opt/tesseract-ocr-read-only && make training && sudo make training-install

RUN cd /opt/tesseract-ocr-read-only && wget http://tesseract-ocr.googlecode.com/files/tesseract-ocr-3.01.eng.tar.gz

RUN cd /opt/tesseract-ocr-read-only && tar xf tesseract-ocr-3.01.eng.tar.gz

RUN cd /opt/tesseract-ocr-read-only
