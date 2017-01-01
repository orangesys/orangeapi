#!/bin/bash

cd ${GOPATH}/src/${ORG_PATH}/orangeapi

make deps
make build

_tag=$1

if [ -z "${_tag}" ]; then
    source _VERSION
    _tag=${_VERSION}
fi

docker build --tag "orangesys/alpine-orangeapi:${_tag}" .
