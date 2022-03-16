#include "binary_search.h"
#include "stdio.h"
#include <math.h>

const int *binary_search(int value, const int *arr, size_t length) {
   if (!length) return NULL;
   size_t mid = floor(length/2);
   if (arr[mid] == value) {
      return (int*) arr+mid;
   }
   if (arr[mid] > value) {
      return binary_search(value, arr, mid);
   }
   return binary_search(value, arr+mid+1, length-(mid+1));
}
