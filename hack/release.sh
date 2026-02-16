#!/bin/bash

set -euo pipefail

source hack/version_helpers.sh

validate_main_branch
VERSION=$(get_version)
if git rev-parse "${VERSION}" >/dev/null 2>&1; then
  echo "Error: tag ${VERSION} already exists." >&2
  exit 1
fi

git tag -a "${VERSION}" -m "Release ${VERSION}"
echo "Tagged ${VERSION}"
