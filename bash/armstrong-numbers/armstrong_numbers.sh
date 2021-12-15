#!/usr/bin/env bash

inp=$1
len=${#inp}
total=0
for (( i=0; i<len; i++ )); do
    ltr="${inp:$i:1}"
    ((p=$ltr**$len))
    ((total=$total+$p))
done
if [[ "${total}" == "${inp}" ]]; then
    echo "true"
else
    echo "false"
fi
