#!/usr/bin/env bash

# Get path to script
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

bawsm="${SCRIPT_DIR}/../bawsm3"
input_folder="${SCRIPT_DIR}/input/"

# Regex defs
regextext='^[0-9]{1,3}-[0-9]{1,2}(N|H)\.[0-9]\.[^.]*\.txt$'
regexarg='^[0-9]{1,3}-[0-9]{1,2}(N|H)$'

# Check if mission to tab is valid
if [[ ! "$1" =~ $regexarg ]]; then 
  >&2 printf "Invalid Argument:\n  got: \"$1\"\n  regex: \"$regexarg\"\n"
  exit 2
fi

# Tabber text variable
output=""

output+='<tabber>' 
# How many strategies
i=0
for item in "${input_folder}/"$1*.txt; do
  # Does file exist
  if [[ ! -f "$item" ]]; then 
    continue
  fi
  i+=1
  filename="${item##*/}"
  # Check if filename is valid 
  if [[ ! "$filename" =~ $regextext ]]; then 
    >&2 printf "Not A Valid Filename:\n  got: \"$filename\"\n  regex: \"$regextext\"\n"
    exit 2
  fi
  title=$(echo "$filename" | grep -oP '(?<=\.\d\.).*(?=\.)')
  # Cosmic Rays Check
  if [[ "$title" == "" ]]; then 
    >&2 printf "Unable to extract Title from filename:\n  got: \"$filename\"\n  regex: \"$regextext\"\n"
    exit 2
  fi
  # Parse the Input File
  # Also capture stdout and stderr into different variables while perserving error code
  {
    IFS=$'\n' read -r -d '' CAPTURED_STDERR;
    IFS=$'\n' read -r -d '' CAPTURED_STDOUT;
    (IFS=$'\n' read -r -d '' _ERRNO_; exit ${_ERRNO_});
  } < <((printf '\0%s\0%d\0' "$("${bawsm}" < "$item")" "${?}" 1>&2) 2>&1)
  # output error msg
  if [[ ! $? == 0 ]]; then 
    >&2 echo "${filename}:${CAPTURED_STDERR}"
    exit 2
  fi
  output+="\n"
  output+='|-| '
  output+="$title ="
  output+="\n\n"
  output+="${CAPTURED_STDOUT}"
  output+="\n"
done
output+='</tabber>'
if [[ $i == 0 ]]; then 
  >&2 printf "No Strategy For: \"$1\"\n"
  exit 1
fi
# Output Tabber text
printf "$output"
