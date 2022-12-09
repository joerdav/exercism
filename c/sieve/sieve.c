#include "sieve.h"
#include <stdbool.h>
#include <stdio.h>
#include <string.h>

uint32_t sieve(uint32_t limit, uint32_t *primes, size_t max_primes) {
   bool candidates[limit];
   for (size_t i = 0; i < limit; i++) candidates[i] = true;
   uint32_t count = 0;
   for (uint32_t i = 2; i <= limit && count < max_primes; i++) {
      if (!candidates[i-1]) {
         continue;
      }
      primes[count++] = i;
      for (uint32_t j = 2; j <= limit; j++) {
         if (i*j>limit) break; 
         candidates[(i*j)-1] = false;
      }
   }
   return count;
}
