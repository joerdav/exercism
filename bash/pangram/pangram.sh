#!/usr/bin/env bash

a="abcdefghijklmnopqrstuvwxyz"
shopt -s nocasematch

for (( i=0; i<${#a}; i++ )); do
    l=${a:$i:1}
    if [[ $1 != *"$l"* ]]; then
        echo "false"
        exit 0
    fi
done

echo "true"
