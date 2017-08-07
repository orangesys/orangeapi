#!/bin/bash

set -ex

_v=$(git describe --always --tags)
version=${_v#*v}
docker build --tag "orangesys/alpine-orangeapi:${version}" .
