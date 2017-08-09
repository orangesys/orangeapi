#!/bin/bash

set -ex

_v=$(git describe --always --tags)
version=${_v#*v}
docker build --tag "orangesys/alpine-orangeapi:${version}" .
docker images
mkdir -p /caches
docker save -o /caches/alpine-orangeapi.tar "orangesys/alpine-orangeapi:${version}"
