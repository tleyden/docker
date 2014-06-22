FROM tleyden5iwx/stroke-width-transform

RUN apt-get -q -y update

ENV GOPATH /opt/go
RUN mkdir -p $GOPATH

RUN apt-get -q -y install git
RUN apt-get -q -y install golang
RUN apt-get -q -y install gcc

# since we are gonna install open-ocr, we need all the tesseract / leptonica deps
RUN apt-get -q -y install libleptonica-dev
RUN apt-get -q -y install libtesseract3 libtesseract-dev
RUN apt-get install -q -y tesseract-ocr-eng

RUN go get -u -v github.com/tleyden/open-ocr

# build open-ocr-preprocessor binary and copy it to /usr/bin
RUN cd $GOPATH/src/github.com/tleyden/open-ocr/cli-preprocessor && go build -v -o open-ocr-preprocessor && cp open-ocr-preprocessor /usr/bin

