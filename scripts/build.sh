#!/bin/bash

set -xe

# Get rid of existing binaries
rm -f dist/orangesys*

# Check if VERSION variable set and not empty, otherwise set to default value
if [ -z "$VERSION" ]; then
  VERSION=$(git describe --always --tags)
fi
echo "Building application version $VERSION"

# Build linux amd64 binaries
OS_PLATFORM_ARG=(linux darwin)
OS_ARCH_ARG=(amd64)
for OS in ${OS_PLATFORM_ARG[@]}; do
  for ARCH in ${OS_ARCH_ARG[@]}; do
    echo "Building binary for $OS/$ARCH..."
    GOARCH=$ARCH GOOS=$OS CGO_ENABLED=0 go build -ldflags "-s -w" -ldflags "-X main.version=${VERSION}" -o "dist/orangeapi_$OS-$ARCH" $PKG_SRC
  done
done
