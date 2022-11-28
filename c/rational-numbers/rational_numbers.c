#include "rational_numbers.h"

static int16_t greatest_common_denominator(int16_t a, int16_t b) {
    while (b != 0) {
        int16_t tmp = b;
        b = a % b;
        a = tmp;
    }
    return a;
}

rational_t add(rational_t a, rational_t b) {
   return reduce((rational_t){ 
      a.numerator*b.denominator+b.numerator*a.denominator, 
      a.denominator*b.denominator
   });
}

rational_t subtract(rational_t a, rational_t b) {
   return reduce((rational_t){ 
      a.numerator*b.denominator-b.numerator*a.denominator, 
      a.denominator*b.denominator
   });
}
rational_t multiply(rational_t a, rational_t b) {
   return reduce((rational_t){ 
      a.numerator*b.numerator, 
      a.denominator*b.denominator
   });
}
rational_t divide(rational_t a, rational_t b) {
   return reduce((rational_t){ 
      a.numerator*b.denominator, 
      b.numerator*a.denominator
   });
}
rational_t absolute(rational_t a) {
   return reduce((rational_t){ 
      abs(a.numerator), 
      abs(a.denominator)
   });
}
rational_t exp_rational(rational_t a, int16_t b) {
   if (b == 0) return (rational_t) { 1, 1};
   if (b > 0) return reduce((rational_t){ 
      pow(a.numerator, b), 
      pow(a.denominator, b)
   });
   int16_t buf = a.numerator;
   return reduce((rational_t){ 
      pow(a.denominator, -b), 
      pow(buf, -b)
   });

}
float exp_real(int16_t a, rational_t b) {
   return pow(a, (float)b.numerator/(float)b.denominator);
}
rational_t reduce(rational_t a) {
    int16_t gcd = greatest_common_denominator(a.numerator, a.denominator);
    a.numerator /= gcd;
    a.denominator /= gcd;
    if (a.denominator < 0) {
        a.numerator = -a.numerator;
        a.denominator = -a.denominator;
    }
    return (rational_t){a.numerator, a.denominator};
}
