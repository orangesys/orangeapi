#!/bin/sh
helm init --client-only
if ! helm version; then exit 99;fi
helm repo remove local
helm repo remove stable
helm repo add or-charts https://orangesys.github.io/charts
helm repo update

exec "/orangeapi_linux-amd64"
