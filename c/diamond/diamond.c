#include "diamond.h"
#include <stdlib.h>
#include <stdint.h>
#include <stdio.h>

//width- (C)*2-1
/*1-2"  A  ",*/
/*2-1" B B ",*/
/*3-0"C   C",*/

/*2" B B ",*/
/*1"  A  "*/
char **make_diamond(const char letter) {
   size_t letter_idx = (letter - 'A') + 1;
   char **result = malloc(letter_idx*2*sizeof(char*));
   size_t len = 0;
   for (size_t i = 0; i < letter_idx; i++) {
      result[len] = calloc(letter_idx*2-1, sizeof(char));
      if (i == 0) {
         result[len][letter_idx-1] = 'A';
         len++;
         continue;
      }
      result[len][(letter_idx-1)-i] = 'A'+(letter_idx-1);
      result[len][letter_idx+i] = 'A'+(letter_idx-1);
      len++;
   }
   if (letter == 'A') return result;
   for (size_t i = 0; i < letter_idx-1; i++) {
      size_t l = letter_idx-i;
      result[len] = calloc(letter_idx*2-1, sizeof(char));
      if (l == 2) {
         result[len][letter_idx-1] = 'A';
         len++;
         continue;
      }
      printf("%zu\n", l);
      result[len][(letter_idx-1)-l] = 'A'+(letter_idx-1);
      result[len][letter_idx+l] = 'A'+(letter_idx-1);
      len++;
   }
   return result;
}


void free_diamond(char **diamond) {
   free(diamond);
}
