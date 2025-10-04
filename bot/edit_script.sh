#!/usr/bin/env bash

jq_file="$(jq -e -r '.' "$1")"
jq_pages_len="$(echo "$jq_file" | jq -e -r '.pages | length - 1')"
jq_output="$jq_file"
for item in $(seq 0 "$jq_pages_len"); do
  jq_page="$(echo "$jq_file" | jq -e -r --argjson i "$item" '.pages[$i]')"

  page_text="$(echo "$jq_page" | jq -e -r '.content')"
  page_title="$(echo "$jq_page" | jq -e -r '.title')"
  if [[ ! $? == 0 ]]; then
    >&2 echo "No title field for page: ${item}"
    continue
  fi

  # EDIT ZONE BEING
  mission_code="$(echo "$page_title" | sed -E 's/[Mm]issions\///' )"

  regextitlehard='^[0-9]{1,3}-[0-9]{1,2}H$'
  regextitlenormal='^[0-9]{1,3}-[0-9]{1,2}$'

  if [[ "$mission_code" =~ $regextitlehard ]]; then
    # HARD
    mission_code="${mission_code}"
  elif [[ "$mission_code" =~ $regextitlenormal ]]; then
    # NORMAL
    mission_code="${mission_code}N"
  else
    >&2 echo "Failed get mission code from title: ${page_title}"
    continue
  fi

  category="$(echo "$page_text" | grep -oE '\[\[[Cc]ategory:[Mm]issions\|[0-9]*\]\]')"
  page_edit="$(echo "$page_text" | sed -z -E 's/(== *Strategy *==[^\[]*)?(\[\[[Cc]ategory:[Mm]issions\|[0-9]*\]\])//')"
  strategy="$(../strategies/get_tabbed_strategy.sh "${mission_code}")" # | sed -e 's/$/\\n/' | tr -d '\n')"

  if [[ ! $? == 0 ]]; then
    continue
  fi
  
  # Remember to set these
  edit="$(printf '%s\n\n==Strategy==\n%s\n\n%s' "${page_edit}" "${strategy}" "${category}")"
  summary="Bot: Add/Reformat Strategy Section"

  # EDIT ZONE END
  # Remember to set edit and summary variables in edit zone
  jq_edit_output_check="$(echo "$jq_output" | jq -e -r --argjson i "$item" --arg edit "$edit" '.pages[$i].edit = $edit')"
  if [[ $? == 0 ]]; then
    jq_output="$jq_edit_output_check"
    jq_edit_output_check="$(echo "$jq_output" | jq -e -r --argjson i "$item" --arg summary "$summary" '.pages[$i].editsummary = $summary')"
    if [[ $? == 0 ]]; then
      jq_output="$jq_edit_output_check"
    else 
      >&2 echo "Failed to set summary for page: ${page_title}"
    fi
  else 
    >&2 echo "Failed to set edit for page: ${page_title}"
  fi


done

echo "$jq_output"
