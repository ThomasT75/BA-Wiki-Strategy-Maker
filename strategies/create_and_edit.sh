#!/usr/bin/env bash

function yes_or_no {
    while true; do
        read -p "$* [y/n]: " yn
        case $yn in
            [Yy]*) return 0  ;;  
            [Nn]*) return  1 ;;
        esac
    done
}

regextext='^[0-9]{1,3}-[0-9]{1,2}(N|H)\.[0-9]\.[^.]*\.txt$'
regexarg='^[0-9]{1,3}-[0-9]{1,2}(N|H)$'

# Check if mission is valid
if [[ ! "$1" =~ $regexarg ]]; then 
  >&2 printf "Invalid Argument:\n  got: \"$1\"\n  regex: \"$regexarg\"\n"
  exit 2
fi
# How many strategies
i=1
for item in "./input/"$1*.txt; do
  # Does file exist
  if [[ ! -f "$item" ]]; then 
    continue
  fi
  filename="${item##*/}"
  # Check if filename is valid 
  if [[ ! "$filename" =~ $regextext ]]; then 
    continue
  fi
  ((i+=1))
done

file1="$1.$i.editing.txt"

nvim "$file1"

if [[ ! -f "$file1" ]]; then
  >&2 printf "File wasn't saved\n"
  exit 2
fi

turns=$(tr ' ' '\n' < "$file1" | grep -ci '\<turn\>')
time=$(tr ' ' '\n' < "$file1" | grep -i -A 1 '\<timed\>' | tail -n 1)

finaltitle=""

while true; do
  read -p "[a]ll clear | [g]ift | [c]hallenge | [t]imed | [v]Challenge/[b]Gift only | a[l]t clear | [d]one : " rvar
  case $rvar in
    [Aa]*) finaltitle+=$turns"-Turn Clear" ;;
    [Gg]*) finaltitle+=$turns"-Turn Gift Clear" ;;
    [Cc]*) finaltitle+=$turns"-Turn Challenge" ;;
    [Tt]*) finaltitle+=$time" Seconds Clear" ;;
    [Ll]*) finaltitle+=$turns"-Turn Alternative Clear" ;;
    [Vv]*) finaltitle+=$turns"-Turn Challenge Only" ;;
    [Bb]*) finaltitle+=$turns"-Turn Gift Only" ;;
    [Dd]*) break ;;
  esac
done

mv -i "$file1" "./input/${file1//editing/$finaltitle}"

./progress.sh
