#!/usr/bin/env bash

function yes_or_no {
    while true; do
        read -p "$* [y/n]: " yn
        case $yn in
            [Yy]*) return 0 ;;  
            [Nn]*) return 1 ;;
        esac
    done
}

jq_file="$(jq -e -r '.' "$1")"
jq_pages_len="$(echo "$jq_file" | jq -e -r '.pages | length - 1')"
jq_output="$jq_file"
for item in $(seq 0 "$jq_pages_len"); do
  jq_page="$(echo "$jq_file" | jq -e -r --argjson i "$item" '.pages[$i]')"

  page_title="$(echo "$jq_page" | jq -e -r '.title')"
  if [[ ! $? == 0 ]]; then
    >&2 echo "No title field for page: ${item}"
    continue
  fi

  page_diff="$(echo "$jq_page" | jq -e -r '.diff')"
  if [[ ! $? == 0 ]]; then
    >&2 echo "No diff for page: ${page_title}"
    continue
  fi

  is_manual_approved="$(echo "$jq_page" | jq -e -r '.approved')"
  if [[ ! $? == 0 ]]; then
    >&2 echo "No approved field for page: ${page_title}"
    continue
  fi
  if [[ ! "$is_manual_approved" == "Manual" ]]; then
    >&2 echo "Page: ${page_title}, auto disapproved Reason: ${is_manual_approved}"
    continue
  fi

  # Show user the diff
  echo "$page_diff" | >&2 less -R 

  approved="No"
  if yes_or_no "Approve Edit? for page: ${page_title}"; then
    approved="Yes"
  else
    approved="No"
  fi

  jq_edit_output_check="$(echo "$jq_output" | jq -e -r --argjson i "$item" --arg output "$approved" '.pages[$i].approved = $output')"
  if [[ $? == 0 ]]; then
    jq_output="$jq_edit_output_check"
  else
    >&2 echo "Failed to set approved to \"$approved\" for page: ${page_title}"
  fi

done

echo "$jq_output"

