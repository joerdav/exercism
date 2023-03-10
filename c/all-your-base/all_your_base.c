#include "all_your_base.h"
#include <stdio.h>

static int to_10(const int8_t digits[DIGITS_ARRAY_SIZE], int64_t base,
                 size_t length) {
  int result = 0;
  for (size_t i = 0; i < length; i++) {
    result += digits[i] * pow(base, i);
  }
  return result;
}

size_t rebase(int8_t digits[DIGITS_ARRAY_SIZE], int16_t input, int16_t output,
              size_t length) {
  int ten = to_10(digits, input, length);
  for (size_t i = 0; i < DIGITS_ARRAY_SIZE; i++)
    digits[i] = 0;
  size_t result_len = 0;
  while (ten != 0) {
    float rem = fmodf(ten, output);
    printf("%d,%d,%f\n", ten, output, rem);
    ten /= output;
    digits[result_len++] = rem;
  }
  return result_len;
}
