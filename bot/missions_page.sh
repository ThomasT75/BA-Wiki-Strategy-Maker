#!/usr/bin/env bash

begin="$1"
amount="$2"
c=0
a=0
output=""
for area in $(seq 1 29); do
  for i in $(seq 1 5); do
    if [[ $begin -le $c ]]; then
      output="${output}|Missions/${area}-${i}"
      ((a=a+1))
      if [[ $a -ge $amount ]]; then
        break 2
      fi
    fi
    ((c=c+1))
  done
  for i in $(seq 1 3); do
    if [[ $begin -le $c ]]; then
      output="${output}|Missions/${area}-${i}H"
      ((a=a+1))
      if [[ $a -ge $amount ]]; then
        break 2
      fi
    fi
    ((c=c+1))
  done
done 

echo "${output#\|}"
