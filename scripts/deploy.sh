#!/bin/bash

set -ex
_v=$(git describe --always --tags)
version=${_v#*v}

docker tag "orangesys/alpine-orangeapi:${version}" "asia.gcr.io/saas-orangesys-io/alpine-orangeapi:${version}"
sudo /opt/google-cloud-sdk/bin/gcloud docker -- push asia.gcr.io/saas-orangesys-io/alpine-orangeapi:${version}
