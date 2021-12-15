#!/usr/bin/env bash

if [ $# != 2 ]; then
    echo "Usage: hamming.sh <string1> <string2>"
    exit 1
fi

A="$1"
B="$2"

if [ ${#A} -ne ${#B} ]; then
    echo "left and right strands must be of equal length"
    exit 1
fi

if [[ ${#A} -eq 0 ]] && [[ ${#B} -eq 0 ]]; then
	echo 0
	exit 0
fi

if [[ $A+$B == *([^AGCT]) ]]; then
    exit 1
fi

DIFF=0

for (( i=0; i<${#A}; i++)); do
    if [ "${A:$i:1}" != "${B:$i:1}" ]; then
        ((DIFF=DIFF+1))
    fi
done

echo "$DIFF"

exit 0
