#include "series.h"
#include <stdlib.h>
#include <string.h>

slices_t slices(char *input_text, unsigned int substring_length) {
   slices_t s = {0};
   if (substring_length == 0) return s;
   s.substring = calloc(MAX_SERIES_RESULTS, sizeof(*s.substring));
   for (; *input_text; input_text++) {
      if (strlen(input_text) < substring_length) break;
      s.substring[s.substring_count] = calloc(MAX_SERIES_LENGTH, sizeof(char));
      strncpy(s.substring[s.substring_count], input_text, substring_length);
      s.substring_count++;
   }
   return s;
}
