FROM tleyden5iwx/stroke-width-transform

RUN apt-get -q -y update

ENV GOPATH /opt/go
RUN mkdir -p $GOPATH

RUN apt-get -q -y install git
RUN apt-get -q -y install golang

RUN go get -u -v github.com/tleyden/open-ocr

# build open-ocr-preprocessor binary and copy it to /usr/bin
RUN cd $GOPATH/src/github.com/tleyden/open-ocr/cli-preprocessor && go build -v -o open-ocr-preprocessor && cp open-ocr-preprocessor /usr/bin

