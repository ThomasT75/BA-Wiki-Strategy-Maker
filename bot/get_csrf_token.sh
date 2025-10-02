#!/usr/bin/env bash

# Strict
set -euo pipefail

# Get path to script
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

. "${SCRIPT_DIR}/globals.sh"


# Get Token For Edits
RESULT=$(curl --compressed -fsSL -X GET \
  -c "${COOKIE_JAR}" \
  -b "${COOKIE_JAR}" \
  -A "${WIKI_AGENT}" \
  "${WIKI_URL}${API_URL}?action=query&format=json&assert=user&meta=tokens&formatversion=2")

ERROR=$(echo "${RESULT}" | jq -e -r '.error // ""')
if [[ ! $? == 0 || ! "$ERROR" == "" ]]; then
  >&2 printf "Failed to get csrf token.\n  Error: %s" "${ERROR}"
  exit 1
fi
CSRF_TOKEN=$(echo "${RESULT}" | jq -r '.query.tokens.csrftoken')

>&2 echo "$CSRF_TOKEN"

echo -n "$CSRF_TOKEN" > "./csrf.txt"

