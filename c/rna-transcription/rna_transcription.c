#include "rna_transcription.h"
#include <stdlib.h>
#include <stdint.h>

#define MAX_RESULT_LEN 13
char *to_rna(const char *dna) {
   char *res = malloc(MAX_RESULT_LEN*sizeof(char));
   size_t curr = 0; 
   for (;*dna;dna++) {
      if (curr >= MAX_RESULT_LEN-1) break;
      if (*dna == 'C') res[curr++] = 'G';
      else if (*dna == 'G') res[curr++] = 'C';
      else if (*dna == 'T') res[curr++] = 'A';
      else if (*dna == 'A') res[curr++] = 'U';
   }
   res[curr] = '\0';
   return res;
}
