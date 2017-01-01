#!/bin/bash

_tag=$1

if [ -z "${_tag}" ]; then
    source _VERSION
    _tag=${_VERSION}
fi

docker tag "orangesys/alpine-orangeapi:${_tag}" "asia.gcr.io/saas-orangesys-io/alpine-orangeapi:${_tag}"
sudo /opt/google-cloud-sdk/bin/gcloud docker -- push asia.gcr.io/saas-orangesys-io/alpine-orangeapi:${tag}
