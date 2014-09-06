FROM ubuntu:14.04

ENV GOPATH /opt/go
RUN apt-get update

RUN apt-get -q -y install libleptonica-dev
RUN apt-get -q -y install libtesseract3 libtesseract-dev
RUN apt-get -q -y install tesseract-ocr
RUN apt-get -q -y install tesseract-ocr-eng
RUN apt-get -q -y install tesseract-ocr-ara    
RUN apt-get -q -y install tesseract-ocr-bel    
RUN apt-get -q -y install tesseract-ocr-ben    
RUN apt-get -q -y install tesseract-ocr-bul    
RUN apt-get -q -y install tesseract-ocr-ces    
RUN apt-get -q -y install tesseract-ocr-dan    
RUN apt-get -q -y install tesseract-ocr-deu    
RUN apt-get -q -y install tesseract-ocr-ell    
RUN apt-get -q -y install tesseract-ocr-fin    
RUN apt-get -q -y install tesseract-ocr-fra    
RUN apt-get -q -y install tesseract-ocr-heb    
RUN apt-get -q -y install tesseract-ocr-hin    
RUN apt-get -q -y install tesseract-ocr-ind    
RUN apt-get -q -y install tesseract-ocr-isl    
RUN apt-get -q -y install tesseract-ocr-ita    
RUN apt-get -q -y install tesseract-ocr-jpn    
RUN apt-get -q -y install tesseract-ocr-kor    
RUN apt-get -q -y install tesseract-ocr-nld    
RUN apt-get -q -y install tesseract-ocr-nor    
RUN apt-get -q -y install tesseract-ocr-pol    
RUN apt-get -q -y install tesseract-ocr-por    
RUN apt-get -q -y install tesseract-ocr-ron    
RUN apt-get -q -y install tesseract-ocr-rus    
RUN apt-get -q -y install tesseract-ocr-spa    
RUN apt-get -q -y install tesseract-ocr-swe    
RUN apt-get -q -y install tesseract-ocr-tha    
RUN apt-get -q -y install tesseract-ocr-tur    
RUN apt-get -q -y install tesseract-ocr-ukr    
RUN apt-get -q -y install tesseract-ocr-vie    
RUN apt-get -q -y install tesseract-ocr-chi-sim
RUN apt-get -q -y install tesseract-ocr-chi-tra

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
