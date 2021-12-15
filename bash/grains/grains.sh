#!/usr/bin/env bash

function grain {
  ((MULTIS=$1-1))
  RESULT=$(bc <<< "2^$MULTIS")
  echo $RESULT
}

O=0

if [ $1 == "total" ]; then
  for I in {1..64}
  do
    SQUARE=$(grain $I)
    O=$(bc <<< "$O+$SQUARE")
  done

else
  [ $1 -le 0 ] || [ $1 -gt 64 ] && echo "Error: invalid input" && exit 1 
  O=$(grain $1)
fi
echo "${O##*[+-]}"

