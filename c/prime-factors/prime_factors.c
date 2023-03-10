#include "prime_factors.h"
#include <stdint.h>

size_t find_factors(uint64_t n, uint64_t *factors) {
  for (uint64_t i = 2; i <= n; i++)
    if (n % i == 0) {
      *(factors++) = i;
      return 1 + find_factors(n / i, factors);
    }
  return 0;
}
