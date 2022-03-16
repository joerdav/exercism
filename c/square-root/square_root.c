#include "square_root.h"

unsigned int square_root(unsigned int radicand) {
   if (radicand == 0 || radicand == 1) return radicand;
   unsigned int start = 1;
   unsigned int end = radicand;
   unsigned int ans;
   while (start <= end) {
      unsigned int mid = (start + end) / 2;
      unsigned int sqr = mid*mid;
      if (sqr == radicand) return mid;
      if (sqr < radicand) {
         start = mid+1;
         ans = mid;
         continue;
      }
      end = mid-1;
   }
   return ans;
}
