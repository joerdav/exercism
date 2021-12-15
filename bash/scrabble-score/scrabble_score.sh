#!/usr/bin/env bash

INP=${1^^}
SCORE=0

declare -A SCORES
SCORES["A"]=1; SCORES["E"]=1; SCORES["I"]=1; SCORES["O"]=1; SCORES["U"]=1;
SCORES["L"]=1; SCORES["N"]=1; SCORES["R"]=1; SCORES["S"]=1; SCORES["T"]=1;
SCORES["D"]=2; SCORES["G"]=2; 
SCORES["B"]=3; SCORES["C"]=3; SCORES["M"]=3; SCORES["P"]=3;
SCORES["F"]=4; SCORES["H"]=4; SCORES["V"]=4; SCORES["W"]=4; SCORES["Y"]=4;
SCORES["K"]=5;
SCORES["J"]=8; SCORES["X"]=8;
SCORES["Q"]=10; SCORES["Z"]=10;


for (( i=0; i<${#INP}; i++ )); do
  LET="${INP:$i:1}"
  LETSCORE=${SCORES[$LET]}
  ((SCORE=SCORE+LETSCORE))
done
echo $SCORE
