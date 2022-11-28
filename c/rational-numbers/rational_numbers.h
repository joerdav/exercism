#ifndef RATIONAL_NUMBERS_H
#define RATIONAL_NUMBERS_H

#include <stdlib.h>
#include <stdint.h>
#include <math.h>
#include <stdio.h>

typedef struct {
	int16_t numerator;
	int16_t denominator;
} rational_t;

rational_t add(rational_t a, rational_t b);
rational_t subtract(rational_t a, rational_t b);
rational_t multiply(rational_t a, rational_t b);
rational_t divide(rational_t a, rational_t b);
rational_t absolute(rational_t a);
rational_t exp_rational(rational_t a, int16_t b);
float exp_real(int16_t a, rational_t b);
rational_t reduce(rational_t a);

#endif
