#!/bin/bash

set -euo pipefail

source hack/version_helpers.sh

validate_main_branch
VERSION=$(get_version)

BUMP="${1:-patch}"

case "${BUMP}" in
  major|minor|patch) ;;
  *)
    echo "Usage: $0 [major|minor|patch]" >&2
    echo "Default: patch" >&2
    exit 1
    ;;
esac

CURRENT=$(echo "${VERSION}" | sed -n 's/.*v\([0-9]*\.[0-9]*\.[0-9]*\).*/\1/p')
if [[ -z "${CURRENT}" ]]; then
  echo "Error: could not parse version from ${VERSION_FILE}" >&2
  exit 1
fi

IFS='.' read -r MAJOR MINOR PATCH <<< "${CURRENT}"

case "${BUMP}" in
  major)
    MAJOR=$((MAJOR + 1))
    MINOR=0
    PATCH=0
    ;;
  minor)
    MINOR=$((MINOR + 1))
    PATCH=0
    ;;
  patch)
    PATCH=$((PATCH + 1))
    ;;
esac

NEW_VERSION="v${MAJOR}.${MINOR}.${PATCH}"

sed -i "s/v${CURRENT}/${NEW_VERSION}/" "${VERSION_FILE}"

echo "Version bumped: v${CURRENT} -> ${NEW_VERSION}"

git checkout -b "bump_version_to_${NEW_VERSION}"
git add "${VERSION_FILE}"
git commit -s -m "Bump version to ${NEW_VERSION}"
