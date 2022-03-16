#include "hamming.h"
#include <string.h>

int compute(const char *lhs, const char *rhs) {
   int lenl = strlen(lhs);
   int lenr = strlen(rhs);
   if (lenl != lenr) return -1;
   if (lenl == 0) return 0;
   int diff = 0;
   for (int i = 0; i < lenl; i++) {
      if (lhs[i] != rhs[i]) diff++;
   }
   return diff;
}
