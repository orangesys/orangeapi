#!/bin/bash
set -ex

version=$(git describe --always --tags|sed 's/^v//')

docker push "orangesys/alpine-orangeapi:${version}"


docker tag "orangesys/alpine-orangeapi:${version}" "asia.gcr.io/${PROJECT_NAME}/alpine-orangeapi:${version}"
gcloud docker -- push asia.gcr.io/${PROJECT_NAME}/alpine-orangeapi:${version}