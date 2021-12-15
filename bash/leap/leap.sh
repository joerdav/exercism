#!/usr/bin/env bash

[[ $# != 1 ]] && echo "Usage: leap.sh <year>" && exit 1
if ! [[ "$1" =~ ^[0-9]+$ ]]; then
  echo "Usage: leap.sh <year>" && exit 1
fi

if [[ $(($1 % 100)) != 0 ]]; then
  [[ $(($1 % 4)) == 0 ]] && echo "true" && exit 0
else
  [[ $(($1 % 400)) == 0 ]] && echo "true" && exit 0
fi

echo "false"
