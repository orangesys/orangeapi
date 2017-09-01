#!/bin/bash

set -ex

version=$(git describe --always --tags)
docker build --tag "orangesys/alpine-orangeapi:${version}" .
docker images
mkdir -p /caches
docker save -o /caches/alpine-orangeapi.tar "orangesys/alpine-orangeapi:${version}"
