#include "perfect_numbers.h"

static int64_t aliquot_sum(int64_t number) {
  int64_t sum = 0;
  for (int64_t i = 1; i < number; i++) {
    if (number % i == 0)
      sum += i;
  }
  return sum;
}

kind classify_number(int64_t number) {
  if (number < 1)
    return ERROR;
  int64_t aliquot = aliquot_sum(number);
  if (aliquot == number)
    return PERFECT_NUMBER;
  if (aliquot > number)
    return ABUNDANT_NUMBER;
  return DEFICIENT_NUMBER;
}
