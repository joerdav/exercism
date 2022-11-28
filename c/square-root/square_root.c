#include "square_root.h"

uint16_t square_root(uint16_t radicand) {
   for (uint16_t i = 0; i <= radicand; i++)
      if (i*i==radicand) return i;
   return -1;
}
