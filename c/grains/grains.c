#include "grains.h"

uint64_t square(uint8_t index) {
   if (index < 1 || index > 64) return 0;
   return 1ull << (index-1ull);
}
uint64_t total(void) {
   uint64_t total = 0ull;
   for (uint8_t i = 1; i <= 64; i++) {
      total += square(i);
   }
   return total;
}
