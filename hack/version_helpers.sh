set -euo pipefail

export VERSION_FILE="cmd/ptr_to_new/version.txt"

function validate_main_branch() {
  BRANCH="$(git rev-parse --abbrev-ref HEAD)"
  if [[ "${BRANCH}" != "main" ]]; then
    echo "Error: must be on the main branch (currently on '${BRANCH}')." >&2
    return 1
  fi

  git fetch origin main
  if [[ "$(git rev-parse HEAD)" != "$(git rev-parse origin/main)" ]]; then
    echo "Error: local main differs from origin/main. Pull or push first." >&2
    return 1
  fi

  if [[ -n "$(git status --porcelain -u no)" ]]; then
    echo "Error: working tree is not clean. Commit or stash changes first." >&2
    return 1
  fi
}

function get_version() {
  if [[ ! -f "${VERSION_FILE}" ]]; then
    echo "Error: ${VERSION_FILE} not found." >&2
    echo "Run this script from the repository root." >&2
    return 1
  fi

  VERSION=$(cat "${VERSION_FILE}")

  echo "${VERSION}"
}
