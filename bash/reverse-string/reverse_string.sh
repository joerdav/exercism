#!/usr/bin/env bash
set -o noglob
RESULT=""
for (( i=0; i<${#1}; i++ )); do
  RESULT=${1:$i:1}$RESULT
done

echo "${RESULT}" 
