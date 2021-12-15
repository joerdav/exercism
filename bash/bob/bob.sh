#!/usr/bin/env bash
INP="$(echo -e "$1" | tr -d '[:space:]')"
ALPHA='[A-Za-z]+'
SHOUT='^[^a-z]+$'
QUESTION='^.*\?$'
WS=$'^[ \t\r\n\f]*$'

[[ $INP =~ $WS ]] && echo "Fine. Be that way!" && exit 0
[[ $INP =~ $ALPHA ]] && [[ $INP =~ $SHOUT ]] && [[ $INP =~ $QUESTION ]] && echo "Calm down, I know what I'm doing!" && exit 0
[[ $INP =~ $ALPHA ]] && [[ $INP =~ $SHOUT ]] && echo "Whoa, chill out!" && exit 0
[[ $INP =~ $QUESTION ]] && echo "Sure." && exit 0

echo "Whatever."
