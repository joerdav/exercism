#!/usr/bin/env bash
inp=$1
len=${#inp}
function printlet {
  echo -n ${1^^}
}
for (( i=0; i<len; i++ )); do
  ltr="${inp:$i:1}"
  if [[ i -eq 0 ]]; then
    printlet $ltr
  fi
  next=(i+1)
  if [[ $ltr == @([ \-_\*]) ]] && [[ next -lt len ]] && [[ ${inp:$next:1} != @([ \-_\*]) ]]; then
    printlet ${inp:$next:1}
  fi
done

echo ""
exit
