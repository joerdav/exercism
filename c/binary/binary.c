#include "binary.h"
#include <string.h>
#include <math.h>

int convert(const char *input) {
   int len = strlen(input);
   int result = 0;
   int pos = 1;
   for (const char *curr = input; *curr; curr++) {
      if (*curr != '0' && *curr != '1') return INVALID;
      int digit = *curr-'0';
      result += digit*pow(2,(len-pos++));
   }
   return result;
}
