#include "isogram.h"

#include <ctype.h>
#include <stdlib.h>
#include <stdio.h>
#include <string.h>

bool is_in_array(char *used, char find) {
   for (unsigned long i = 0; i < strlen(used); i++) {
      if (used[i] == find) {
         return true;
      }
   }
   return false;
}

bool is_isogram(const char *phrase) 
{
   if (phrase == NULL) return false;
   char used[26];
   memset (used, 0 , sizeof(used));
   int usedlen = 0;
   for (unsigned long i = 0; i < strlen(phrase); i++) {
      char l = tolower(phrase[i]);
      if (!isalpha(l)) {
         continue;
      }
      if (is_in_array(used, l)) {
         return false;
      }
      used[usedlen++] = l;
   }
   return true;
}
