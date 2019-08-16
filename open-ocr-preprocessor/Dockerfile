# Builder image
FROM ubuntu:18.04 AS builder

# Prevent interactive steps in apt-get
ARG DEBIAN_FRONTEND=noninteractive

# Install build dependencies
RUN apt-get update && apt-get install -y \
    cmake \
    gcc \
    git \
    golang \
    libboost1.62-all-dev \
    libleptonica-dev \
    libopencv-calib3d-dev \
    libopencv-calib3d3.2 \
    libopencv-contrib-dev \
    libopencv-contrib3.2 \
    libopencv-core-dev \
    libopencv-core3.2 \
    libopencv-dev \
    libopencv-features2d-dev \
    libopencv-features2d3.2 \
    libopencv-flann-dev \
    libopencv-flann3.2 \
    libopencv-highgui-dev \
    libopencv-highgui3.2 \
    libopencv-imgproc-dev \
    libopencv-imgproc3.2 \
    libopencv-ml3.2 libopencv-ml-dev \
    libopencv-objdetect-dev \
    libopencv-objdetect3.2 \
    libopencv-photo-dev \
    libopencv-photo3.2 \
    libopencv-video-dev \
    libopencv-video3.2 \
    libtesseract-dev \
    libtesseract4 \
    tesseract-ocr-eng \
    unzip \
    wget

# Build open-ocr-preprocessor binary with golang
ENV GOPATH /opt/go
RUN mkdir -p $GOPATH \
    && go get -u -v github.com/tleyden/open-ocr \
    && cd $GOPATH/src/github.com/tleyden/open-ocr/cli-preprocessor \
    && go build -v -o open-ocr-preprocessor

# Build DetectText
RUN cd /opt \
    && git clone https://github.com/aperrault/DetectText.git \
    && cd ./DetectText \
    && mkdir build \
    && cd build \
    && cmake .. \
    && make

# Production image
FROM ubuntu:18.04

# Install runtime dependencies  
RUN apt-get -q -y update && apt-get install -q -y \
    ghostscript\
    libopencv-imgcodecs3.2

# Copy DetectText and open-ocr-preprocessor binaries from builder stage
COPY --from=builder /opt/DetectText/build/DetectText /usr/bin/
COPY --from=builder /opt/go/src/github.com/tleyden/open-ocr/cli-preprocessor/open-ocr-preprocessor /usr/bin/
