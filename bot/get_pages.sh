#!/usr/bin/env bash

# Strict
set -euo pipefail

# Get path to script
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

. "${SCRIPT_DIR}/globals.sh"

# sleep 5s

# GET PAGE SOURCE + TIMESTAMPS FOR EDITS
RESULT=$(curl --compressed -fsSL -X GET \
  -c "${COOKIE_JAR}" \
  -b "${COOKIE_JAR}" \
  -A "${WIKI_AGENT}" \
  --url-query "format=json" \
  --url-query "assert=user" \
  --url-query "curtimestamp=1" \
  --url-query "prop=revisions" \
  --url-query "formatversion=2" \
  --url-query "rvprop=timestamp|content" \
  --url-query "rvslots=main" \
  --url-query "titles=${1}" \
  "${WIKI_URL}${API_URL}?action=query")

ERROR=$(echo "${RESULT}" | jq -e -r '.error // ""')
if [[ ! $? == 0 || ! "$ERROR" == "" ]]; then
  >&2 printf "Failed to get pages.\n  Error: %s\n" "${ERROR}"
  exit 1
fi

# Convert Pages JSON Into Edit Format
JQ_TRANSLATION='
{
  curtimestamp, 
  "pages": [.query.pages[] | 
    {
      "title": .title, 
      "revtimestamp": .revisions[0].timestamp,
      "content": .revisions[0].slots.main.content
    }
  ]
}
'
JQ_OUTPUT="$(echo "${RESULT}" | jq -e -r "${JQ_TRANSLATION}")"
if [[ ! $? == 0 || "$JQ_OUTPUT" == "" ]]; then
  >&2 printf "Failed to Parse Result into json: %s\n" "${RESULT}"
else
  echo "$JQ_OUTPUT"
fi


