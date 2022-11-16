#include "hamming.h"
#include <string.h>

int compute(const char *lhs, const char *rhs) {
   if (!lhs || !rhs) return -1;
   int diff = 0;
   char l, r;
   for (l = *lhs, r = *rhs; l && r; l=*++lhs, r=*++rhs) {
      if (l!=r) diff++;
   }
   return l || r ? -1 : diff;
}
