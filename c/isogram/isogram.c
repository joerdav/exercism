#include "isogram.h"
#include <string.h>
#include <ctype.h>
#include <stdio.h>
#include <stdlib.h>

bool is_isogram(const char phrase[]) {
   if (phrase == NULL) return false;
   bool found[ALPHABET_LEN]={0};
   for (size_t i = 0; phrase[i]; i++) {
      unsigned char letter = phrase[i];
      if (!isalpha(letter)) {
         continue;
      }
      int lower = tolower(letter)-'a';
      if (found[lower]) {
         return false;
      }
      found[lower] = true;
   }
   return true;
}
