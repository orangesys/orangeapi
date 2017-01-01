#!/bin/bash

_tag=$1

if [ -z "${_tag}" ]; then
    source _VERSION
    _tag=${_VERSION}
fi

docker tag "orangesys/alpine-orangeapi:${_tag}" "asia.gcr.io/saas-orangesys-io/alpine-orangeapi:${_tag}"
docker login -e $DOCKER_EMAIL -u _json_key -p "$(cat ${HOME}/account-auth.json)" https://asia.gcr.io
docker push asia.gcr.io/saas-orangesys-io/alpine-orangeapi:${_tag}
