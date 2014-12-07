#!/usr/bin/env bash

./generate_dockerfiles.sh

cp README.md cpu/develop/
cp README.md cpu/master/
cp README.md gpu/develop/
cp README.md gpu/master/
