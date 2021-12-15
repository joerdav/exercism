#include "isogram.h"

#include <ctype.h>
#include <stdlib.h>

bool is_in_array(char used[], char find) {
   for (unsigned long i = 0; i < sizeof(*used) / sizeof(char); i++) {
      if (used[i] == find) {
         return true;
      }
   }
   return false;
}

bool is_isogram(const char phrase[]) 
{
   if (phrase == NULL) return false;
   char used[26];
   int usedlen = 0;
   for (unsigned long i = 0; i < sizeof(*phrase) / sizeof(char); i++) {
      char l = tolower(phrase[i]);
      if (is_in_array(used, l)) {
         return false;
      }
      used[usedlen] = l;
      usedlen++;
   }
   return true;
}
