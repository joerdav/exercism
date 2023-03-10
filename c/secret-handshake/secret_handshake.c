#include "secret_handshake.h"

#include <stdlib.h>
#include <stdint.h>
#include <string.h>

#define CHECK_BIT(bitmap, idx) ((bitmap & (1 << idx)) == (1 << idx))

static void append(char** arr, char* str, size_t pos) {
   arr[pos] = malloc(15*sizeof(char));
   strcpy(arr[pos], str);
}
const char **commands(size_t number) {
   char **res = calloc(4, sizeof(char*));
   size_t pos = 0;
   if (CHECK_BIT(number, 4)) {
      if (CHECK_BIT(number, 3)) append(res, "jump", pos++);
      if (CHECK_BIT(number, 2)) append(res, "close your eyes", pos++);
      if (CHECK_BIT(number, 1)) append(res, "double blink", pos++);
      if (CHECK_BIT(number, 0)) append(res, "wink", pos++);
   } else {
      if (CHECK_BIT(number, 0)) append(res, "wink", pos++);
      if (CHECK_BIT(number, 1)) append(res, "double blink", pos++);
      if (CHECK_BIT(number, 2)) append(res, "close your eyes", pos++);
      if (CHECK_BIT(number, 3)) append(res, "jump", pos++);
   }
   return (const char**) res;
}
