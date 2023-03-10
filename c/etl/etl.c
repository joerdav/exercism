#include "etl.h"
#include <stdlib.h>
#include <ctype.h>

static int comp (const void * a, const void * b) {
   return *((char*)a)-*((char*)b);
}
int convert(const legacy_map *input, const size_t input_len, new_map **output) {
   *output = calloc(26, sizeof(new_map));
   size_t len = 0;
   for (size_t i = 0; i < input_len; i++) {
      for (char *key = (char*)input[i].keys; *key; key++) {
         char ch = tolower(*key);
         (*output)[len++] = (new_map){ch, input[i].value};
      }
   }
   qsort(*output, len, sizeof(new_map), comp);
   return len;
}
