
FROM tleyden5iwx/go-tesseract

# go get open-ocr
RUN go get -u -v github.com/tleyden/open-ocr

# build open-ocr-httpd binary and copy it to /usr/bin
RUN cd $GOPATH/src/github.com/tleyden/open-ocr/cli-httpd && go build -v -o open-ocr-httpd && cp open-ocr-httpd /usr/bin

# build open-ocr-worker binary and copy it to /usr/bin
RUN cd $GOPATH/src/github.com/tleyden/open-ocr/cli-worker && go build -v -o open-ocr-worker && cp open-ocr-worker /usr/bin

