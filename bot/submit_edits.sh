#!/usr/bin/env bash

# Strict
set -euo pipefail

# Get path to script
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

. "${SCRIPT_DIR}/globals.sh"

jq_file="$(jq -e -r '.' "$1")"
if [[ ! $? == 0 ]]; then
  >&2 echo "Fail with provided file"
  exit 2
fi
jq_output="$jq_file"
jq_pages_len="$(echo "$jq_file" | jq -e -r '.pages | length - 1')"
if [[ ! $? == 0 ]]; then
  >&2 echo "Failed to retrive amount of pages"
  exit 2
fi

csrf="$(cat "${SCRIPT_DIR}/csrf.txt")"
if [[ ! $? == 0 ]]; then
  >&2 echo "Invalid csrf.txt file???"
  exit 2
fi

curtimestamp="$(echo "$jq_file" | jq -e -r '.curtimestamp')"
if [[ ! $? == 0 ]]; then
  >&2 echo "No curtimestamp field in provided file"
  exit 2
fi
for item in $(seq 0 "$jq_pages_len"); do
  jq_page="$(echo "$jq_file" | jq -e -r --argjson i "$item" '.pages[$i]')"
  if [[ ! $? == 0 ]]; then
    >&2 echo "No page field in provided file"
    exit 2
  fi

  # Get Page Title Early for Debug Print
  page_title="$(echo "$jq_page" | jq -e -r '.title')"
  if [[ ! $? == 0 ]]; then
    >&2 echo "No title field for page: ${item}"
    continue
  fi

  # Check If The Edit Was Approved
  is_manual_approved="$(echo "$jq_page" | jq -e -r '.approved')"
  if [[ ! $? == 0 ]]; then
    >&2 echo "No approved field for page: ${page_title}"
    continue
  fi
  if [[ ! "$is_manual_approved" == "Yes" ]]; then
    continue
  fi

  # Get Edit Content
  page_edit="$(echo "$jq_page" | jq -e -r '.edit')"
  if [[ ! $? == 0 ]]; then
    >&2 echo "No edit field for page: ${page_title}"
    continue
  fi

  # Get Edit Summary
  page_edit_summary="$(echo "$jq_page" | jq -e -r '.editsummary')"
  if [[ ! $? == 0 ]]; then
    >&2 echo "No editsummary field for page: ${page_title}"
    continue
  fi

  # Get Revision Timestamp
  page_revtimestamp="$(echo "$jq_page" | jq -e -r '.revtimestamp')"
  if [[ ! $? == 0 ]]; then
    >&2 echo "No revtimestamp field for page: ${page_title}"
    continue
  fi

  sleep 5s

  RESULT=$(curl --compressed -fsSL -X POST \
    -c "${COOKIE_JAR}" \
    -b "${COOKIE_JAR}" \
    -A "${WIKI_AGENT}" \
    --url-query "format=json" \
    --url-query "assert=user" \
    --url-query "formatversion=2" \
    --url-query "nocreate=1" \
    --url-query "bot=1" \
    --url-query "basetimestamp=${page_revtimestamp}" \
    --url-query "starttimestamp=${curtimestamp}" \
    --url-query "title=${page_title}" \
    --url-query "summary=${page_edit_summary}" \
    --url-query "text=${page_edit}" \
    --data-urlencode "token=${csrf}" \
    "${WIKI_URL}${API_URL}?action=edit")

  ERROR=$(echo "${RESULT}" | jq -e -r '.error // ""')
  if [[ ! $? == 0 || ! "$ERROR" == "" ]]; then
    >&2 printf "Failed to submit edit.\n  Error: %s\n Continuing and Not Retrying in 1s." "${ERROR}"
    jq_output_check="$(echo "$jq_output" | jq -e -r --argjson i "$item" --arg output "Failed" '.pages[$i].approved = $output')"
    if [[ $? == 0 ]]; then
      jq_output="$jq_output_check"
      echo "$jq_output" > "$1"
    else
      >&2 echo "Failed to set approved to \"Failed\" for page: ${page_title}"
    fi
    sleep 1s
  else
    jq_output_check="$(echo "$jq_output" | jq -e -r --argjson i "$item" --arg output "Submitted" '.pages[$i].approved = $output')"
    if [[ $? == 0 ]]; then
      jq_output="$jq_output_check"
      echo "$jq_output" > "$1"
    else
      >&2 echo "Failed to set approved to \"Submitted\" for page: ${page_title}"
    fi
  fi

done
