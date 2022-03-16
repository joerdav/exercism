#ifndef RATIONAL_NUMBERS_H
#define RATIONAL_NUMBERS_H

#include <stdint.h>
#include <stdlib.h>
#include <math.h>

typedef struct {
	int16_t numerator;
	int16_t denominator;
} rational_t;

rational_t add(rational_t l, rational_t r);
rational_t subtract(rational_t l, rational_t r);
rational_t multiply(rational_t l, rational_t r);
rational_t divide(rational_t l, rational_t r);
rational_t absolute(rational_t l);
rational_t exp_rational(rational_t l, uint16_t r);
float exp_real(uint16_t l, rational_t r);
rational_t reduce(rational_t l);

#endif
