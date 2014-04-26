FROM debian:jessie

ENV GOPATH /opt/go
RUN apt-get update

RUN apt-get -q -y install libleptonica-dev
RUN apt-get -q -y install -t testing libtesseract3 libtesseract-dev
RUN apt-get install -q -y -t testing tesseract-ocr-eng
RUN apt-get -q -y install git
RUN apt-get -q -y install golang
RUN apt-get -q -y install gcc

# In theory, should be able to set export TESSDATA_PREFIX=/usr/share/tesseract-ocr/, 
# but when I tried I still got error: Error opening data file /usr/local/share/tessdata/eng.traineddata
# Workaround: just copy the data to where it expects
RUN mkdir -p /usr/local/share/tessdata/
RUN cp -R /usr/share/tesseract-ocr/tessdata/* /usr/local/share/tessdata/

RUN mkdir -p $GOPATH

RUN go get -u -v github.com/GeertJohan/go.tesseract
RUN go get -u -v github.com/davecgh/go-spew/spew