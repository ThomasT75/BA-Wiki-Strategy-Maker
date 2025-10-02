#!/usr/bin/env bash

# get path to script
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

# Regex defs
regextext='^[0-9]{1,3}-[0-9]{1,2}(N|H)\.[0-9]\.[^.]*\.txt$'

bawsm="${SCRIPT_DIR}/../bawsm3"
input_folder="${SCRIPT_DIR}/input/"


for item in "$input_folder/"*.txt; do
  # Does file exist
  if [[ ! -f "$item" ]]; then 
    continue
  fi
  # Set filename
  filename="${item##*/}"
  # Check if filename is valid 
  if [[ ! "$filename" =~ $regextext ]]; then 
    >&2 printf "Not A Valid Filename:\n  got: \"$filename\"\n  regex: \"$regextext\"\n"
    exit 2
  fi
  # Parse the Input File
  # Also capture stdout and stderr into different variables while perserving error code
  {
    IFS=$'\n' read -r -d '' CAPTURED_STDERR;
    IFS=$'\n' read -r -d '' CAPTURED_STDOUT;
    (IFS=$'\n' read -r -d '' _ERRNO_; exit ${_ERRNO_});
  } < <((printf '\0%s\0%d\0' "$("${bawsm}" < "$item")" "${?}" 1>&2) 2>&1)
  # Save file or output error msg
  if [[ ! $? == 0 ]]; then 
    >&2 echo "${filename}:${CAPTURED_STDERR}"
  fi
done
