#!/usr/bin/env bash

match=false

function checker {
    if [[ $(expr $1 % $2) -eq 0 ]]
    then
        match=true
        echo -n "$3"
    fi
}

checker $1 3 "Pling"
checker $1 5 "Plang"
checker $1 7 "Plong"


if [[ "$match" = false ]]
then
    echo "$1"
fi
echo ""
exit 0
