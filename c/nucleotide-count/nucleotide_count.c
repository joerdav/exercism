#include "nucleotide_count.h"

#include <stdio.h>
#include <stdlib.h>

char *count(const char *dna_strand) {
   char *result = calloc(50, sizeof(char));
   int a = 0;
   int c = 0;
   int g = 0;
   int t = 0;
   for (;*dna_strand;dna_strand++) {
      switch (*dna_strand) {
         case 'A':
            a++;
            break;
         case 'C':
            c++;
            break;
         case 'G':
            g++;
            break;
         case 'T':
            t++;
            break;
         default:
            return result;
      }
   }
   sprintf(result, "A:%d C:%d G:%d T:%d", a, c, g, t);
   return result;
}
