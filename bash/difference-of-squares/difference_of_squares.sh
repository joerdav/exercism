#!/usr/bin/env bash

SQ_O_S="SQUARE_OF_SUM"
S_O_SQ="SUM_OF_SQUARES"
F=${1^^}
I=$2


function sqos {
  ((SUM=(($1+1)*$1)/2))
  ((SQUARE=$SUM*$SUM))
  echo $SQUARE
}

function sosq {
  SUM=0
  for (( i=1; i<=$1; i++ ))
  do
    ((SUM=$SUM+($i*$i)))
  done
  echo $SUM
}

function diff {
  SQOS=$(sqos $1)
  SOQS=$(sosq $1)
  D=$((SQOS-SOQS))
  echo $D
}

[[ $F == $SQ_O_S ]] && echo $(sqos $I) && exit 0
[[ $F == $S_O_SQ ]] && echo $(sosq $I) && exit 0
diff $I
