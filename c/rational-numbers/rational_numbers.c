#include "rational_numbers.h"

static int16_t gcd(int16_t a, int16_t b) {
   if (b == 0)
      return a;
   return gcd(abs(b), abs(a) % abs(b));
}
rational_t reduce(rational_t r) {
  rational_t tmp = {0};
  if (r.numerator != 0) {
       int div = gcd(r.numerator, r.denominator);
       tmp.numerator = r.numerator / div;
       tmp.denominator = r.denominator / div;
  } else
       tmp.denominator = 1;
  if (tmp.denominator < 0) {
       tmp.numerator *= -1;
       tmp.denominator *= -1;
  }
  return tmp;
}

rational_t add(rational_t l, rational_t r) {
   rational_t res =  (rational_t){
      l.numerator*r.denominator+l.denominator*r.numerator, 
      l.denominator*r.denominator 
   };
   return reduce(res);
}
rational_t subtract(rational_t l, rational_t r) {
   rational_t res = (rational_t){
      l.numerator*r.denominator-l.denominator*r.numerator, 
      l.denominator*r.denominator 
   };
   return reduce(res);
}
rational_t multiply(rational_t l, rational_t r) {
   rational_t res = (rational_t){
      l.numerator*r.numerator, 
      l.denominator*r.denominator 
   };
   return reduce(res);
}
rational_t divide(rational_t l, rational_t r) {
   rational_t res = (rational_t){
      l.numerator*r.denominator, 
      l.denominator*r.numerator 
   };
   return reduce(res);
}
rational_t absolute(rational_t l) {
   rational_t res = {0};
   res.numerator = abs(l.numerator);
   res.denominator = abs(l.denominator);
   return res;
}
rational_t exp_rational(rational_t n, uint16_t p) {
   rational_t res = {0};
   res.numerator = pow(n.numerator, p);
   res.denominator = pow(n.denominator, p);
   return reduce(res);
}
float exp_real(uint16_t r, rational_t p) {
   return pow(pow(r, p.numerator), 1.0/p.denominator);
}
