#!/bin/bash

set -o errexit
set -o nounset
set -o pipefail

GIT_TAG=$(git describe --abbrev=0 --tags 2>/dev/null || echo $TAG_NAME)

cat << EOF
STABLE_BUILD_GIT_TAG ${GIT_TAG-}
EOF
