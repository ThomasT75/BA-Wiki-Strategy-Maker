#!/usr/bin/env bash

for item in './responses/mission.api.'*.json; do
  ./submit_edits.sh "$item"
  if [[ ! $? == 0 ]]; then
    >&2 echo "failed to submit edit for response: ${item}"
    exit 2
  fi
done
