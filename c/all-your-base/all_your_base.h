#ifndef ALL_YOUR_BASE_H
#define ALL_YOUR_BASE_H

#define DIGITS_ARRAY_SIZE 64

#include <math.h>
#include <stdint.h>
#include <stdlib.h>

size_t rebase(int8_t digits[DIGITS_ARRAY_SIZE], int16_t input, int16_t output,
              size_t length);

#endif
