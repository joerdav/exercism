#include "pangram.h"

#define ALL 67108863

bool is_pangram(const char *sentence) {
   if (sentence == NULL || strlen(sentence) == 0) return false;
   int32_t hash = 0;
   for (char *curr = (char*)sentence; *curr && *curr != '\0'; curr++) {
      if (!isalpha(*curr)) continue;
      hash |= 1 << (tolower(*curr)-'a');
      if (hash == ALL) return true;
   }
   return false;
}
