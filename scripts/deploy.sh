#!/bin/bash

set -xe
version=$(git describe --always --tags|sed 's/^v//')

kubectl --namespace apigateway patch deployment sys-orangeapi \
-p '{"spec":{"template":{"spec":{"containers":[{"name":"sys-orangeapi","image":"asia.gcr.io/'"$PROJECT_NAME"'/alpine-orangeapi:'"$version"'"}]}}}}'