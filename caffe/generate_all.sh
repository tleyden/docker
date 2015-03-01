#!/usr/bin/env bash

./generate_dockerfiles.sh

cp README.md cpu/master/
cp README.md gpu/master/

cp caffe-ld-so.conf cpu/master/
cp caffe-ld-so.conf gpu/master/
