FROM ubuntu

ENV GOPATH /opt/go

RUN apt-get update && apt-get install -y \
	autoconf \
	automake \
	libtool \
	autoconf-archive \
	pkg-config \
	libpng12-dev \
	libjpeg8-dev \
	libtiff5-dev \
	zlib1g-dev \ 
	libicu-dev \
	libpango1.0-dev \
	libcairo2-dev \
	git \
	golang \
	gcc \
	curl && \
	rm -rf /var/lib/apt/lists/*

RUN curl http://www.leptonica.org/source/leptonica-1.74.1.tar.gz -o leptonica-1.74.1.tar.gz && \
	tar -zxvf leptonica-1.74.1.tar.gz && \
	cd leptonica-1.74.1 && ./configure && make && make install && \
	cd .. && rm -rf leptonica*

RUN git clone --depth 1 https://github.com/tesseract-ocr/tesseract.git && \
	cd tesseract && \
	./autogen.sh && \
	./configure && \
	LDFLAGS="-L/usr/local/lib" CFLAGS="-I/usr/local/include" make && \
	make install && \
	ldconfig && \
	make training && \
	make training-install && \
	cd .. && rm -rf tesseract

# Get trainnedata
RUN git clone https://github.com/tesseract-ocr/tessdata && \
	mv  -v tessdata/* /usr/local/share/tessdata/ && \
	rm -rf tessadata

RUN mkdir -p $GOPATH
	
# go get open-ocr
RUN go get -u -v -t github.com/tleyden/open-ocr

# build open-ocr-httpd binary and copy it to /usr/bin
RUN cd $GOPATH/src/github.com/tleyden/open-ocr/cli-httpd && go build -v -o open-ocr-httpd && cp open-ocr-httpd /usr/bin

# build open-ocr-worker binary and copy it to /usr/bin
RUN cd $GOPATH/src/github.com/tleyden/open-ocr/cli-worker && go build -v -o open-ocr-worker && cp open-ocr-worker /usr/bin