for file1 in *.txt; do
  if [[ ! -f "$file1" ]]; then
    >&2 printf "File wasn't saved\n"
    exit 2
  fi

  turns=$(tr ' ' '\n' < "$file1" | grep -ci '\<turn\>')

  filename=$(echo "$file1" \
  | sed -E "s/s Time Challenge.txt/ Seconds Clear.txt/" \
  )
  # echo "$filename"
  mv -i "$file1" "$filename"
done

  # | sed -E "s/[0-9]+-Turn Clear.txt/$turns-Turn Clear.txt/" \
  # | sed -E "s/Clear without Gift.txt/Challenge.txt/" \
  # | sed -E "s/Clear with Gift.txt/Gift Clear.txt/" \
  # | sed -E "s/ Seconds Clear.txt/s Time Challenge.txt/" \
