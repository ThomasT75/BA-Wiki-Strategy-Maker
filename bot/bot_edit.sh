#!/usr/bin/env bash

folder="./responses/"

mkdir -p "$folder"
if [[ ! $? == 0 ]]; then
  exit 1
fi

for i in $(seq 0 $(( (29 * 8) / 48 )) ); do
  filename="${folder}/mission.api.${i}.json"

  jsonfile="$(./get_pages.sh "$(./missions_page.sh $(( 0 + ($i*48) )) 48)")"
  if [[ ! $? == 0 ]]; then
    exit 1
  fi
  echo "$jsonfile" > "$filename"

  jsonfile="$(./edit_script.sh "$filename")"
  if [[ ! $? == 0 ]]; then
    exit 1
  fi
  echo "$jsonfile" > "$filename"

  jsonfile="$(./parse_edits.sh "$filename")"
  if [[ ! $? == 0 ]]; then
    exit 1
  fi
  echo "$jsonfile" > "$filename"

  jsonfile="$(./approve_edits.sh "$filename")"
  if [[ ! $? == 0 ]]; then
    exit 1
  fi
  echo "$jsonfile" > "$filename"
done

>&2 echo "run bot_submit.sh when ready" 
