#!/bin/bash

set -euo pipefail

BUMP="${1:-patch}"

VERSION_FILE="myversion/version.go"

if [[ ! -f "${VERSION_FILE}" ]]; then
  echo "Error: ${VERSION_FILE} not found." >&2
  echo "Run this script from the repository root." >&2
  exit 1
fi

case "${BUMP}" in
  major|minor|patch) ;;
  *)
    echo "Usage: $0 [major|minor|patch]" >&2
    echo "Default: patch" >&2
    exit 1
    ;;
esac

CURRENT=$(grep -oP 'Version = "v\K[0-9]+\.[0-9]+\.[0-9]+' "${VERSION_FILE}")
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

sed -i "s/Version = \"v${CURRENT}\"/Version = \"${NEW_VERSION}\"/" "${VERSION_FILE}"

echo "Version bumped: v${CURRENT} -> ${NEW_VERSION}"

git add "${VERSION_FILE}"
git commit -s -m "Bump version to ${NEW_VERSION}"
git tag -a "${NEW_VERSION}" -m "Release ${NEW_VERSION}"

echo "Tagged ${NEW_VERSION}"