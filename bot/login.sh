#!/usr/bin/env bash

# Strict
set -euo pipefail

# Get path to script
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

. "${SCRIPT_DIR}/globals.sh"

# GET LOGIN TOKEN
RESULT=$(curl --compressed -fsSL -X GET \
  -c "${COOKIE_JAR}" \
  -b "${COOKIE_JAR}" \
  -A "${WIKI_AGENT}" \
  "${WIKI_URL}${API_URL}?action=query&meta=tokens&type=login&format=json")

ERROR=$(echo "${RESULT}" | jq -e -r '.error // ""')
if [[ ! $? == 0 || ! "$ERROR" == "" ]]; then
  >&2 printf "Failed to get login token.\n  Error: %s\n" "${ERROR}"
  exit 1
fi

LOGIN_TOKEN="$(echo "${RESULT}" | jq -r '.query.tokens.logintoken')"
>&2 echo "$LOGIN_TOKEN"

# POST LOGIN
RESULT=$(curl --compressed -fsSL -X POST \
  -c "${COOKIE_JAR}" \
  -b "${COOKIE_JAR}" \
  -A "${WIKI_AGENT}" \
  -d lgname="$WIKI_USER" \
  -d lgpassword="$WIKI_PASSWORD" \
  --data-urlencode lgtoken="$LOGIN_TOKEN" \
  "${WIKI_URL}${API_URL}?action=login&format=json")

LOGIN_RESULT_CODE="$(echo "${RESULT}" | jq -e -r '.login.result')"

if [[ ! "$LOGIN_RESULT_CODE" == "Success" ]]; then
  >&2 printf "Failed to login.\n  Error: %s\n" "${RESULT}"
  exit 1
fi

>&2 echo "${RESULT}"
