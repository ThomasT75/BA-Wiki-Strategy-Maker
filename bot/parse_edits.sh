#!/usr/bin/env bash

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

  page_content="$(echo "$jq_page" | jq -e -r '.content')"
  page_edit="$(echo "$jq_page" | jq -e -r '.edit')"
  if [[ ! $? == 0 ]]; then
    >&2 echo "No Edit for page: ${page_title}"
    continue
  fi

  # Create Diff
  output="$(diff --color="always" <(echo "$page_content") <(echo "$page_edit"))"
  if [[ $? == 1 ]]; then
    # printf "$output"
    jq_edit_output_check="$(echo "$jq_output" | jq -e -r --argjson i "$item" --arg output "$output" '.pages[$i].diff = $output')"
    if [[ $? == 0 ]]; then
      jq_output="$jq_edit_output_check"
    else
      >&2 echo "Failed to set Diff for page: ${page_title}"
    fi
  fi

  # Check if there is something to change else disapprove edit
  is_there_diff="$(echo "$jq_output" | jq -r --argjson i "$item" '.pages[$i].diff // ""')"
  if [[ ! "$is_there_diff" ]]; then
    jq_edit_output_check="$(echo "$jq_output" | jq -e -r --argjson i "$item" --arg output "NoDiff" '.pages[$i].approved = $output')"
    if [[ $? == 0 ]]; then
      jq_output="$jq_edit_output_check"
    else
      >&2 echo "Failed to set approved to \"NoDiff\" for page: ${page_title}"
    fi
  fi

  # Disapprove Edits with Summary
  is_there_edit_summary="$(echo "$jq_output" | jq -r --argjson i "$item" '.pages[$i].editsummary // ""')"
  if [[ ! "$is_there_edit_summary" ]]; then
    jq_edit_output_check="$(echo "$jq_output" | jq -e -r --argjson i "$item" --arg output "NoSummary" '.pages[$i].approved = $output')"
    if [[ $? == 0 ]]; then
      jq_output="$jq_edit_output_check"
    else
      >&2 echo "Failed to set approved to \"NoSummary\" for page: ${page_title}"
    fi
  fi

  # Requires Manual Approved  
  is_there_approved="$(echo "$jq_output" | jq -r --argjson i "$item" '.pages[$i].approved // ""')"
  if [[ ! "$is_there_approved" ]]; then
    jq_edit_output_check="$(echo "$jq_output" | jq -e -r --argjson i "$item" --arg output "Manual" '.pages[$i].approved = $output')"
    if [[ $? == 0 ]]; then
      jq_output="$jq_edit_output_check"
    else
      >&2 echo "Failed to set approved to \"Manual\" for page: ${page_title}"
    fi
  fi

  # If We re edit this page we should Reapprove it
  if [[ ! "$is_there_approved" == "Yes" ]]; then
    jq_edit_output_check="$(echo "$jq_output" | jq -e -r --argjson i "$item" --arg output "Manual" '.pages[$i].approved = $output')"
    if [[ $? == 0 ]]; then
      jq_output="$jq_edit_output_check"
    else
      >&2 echo "Failed to set approved to \"Manual\" for page: ${page_title}"
    fi
  fi

done

echo "$jq_output"

