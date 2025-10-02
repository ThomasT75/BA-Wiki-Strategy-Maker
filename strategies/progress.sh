#!/usr/bin/env bash

lhs_done="$(find ./input/ -iname "*N.1.*.txt" -or -iname "*H.1.*.txt" | sort | wc -l)"
rhs_need="$((29*8))"

echo "${lhs_done}/${rhs_need}"

