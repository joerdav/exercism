#include "binary_search.h"
#include <stdio.h>

const int *binary_search(int value, const int *arr, size_t length) {
   if (length == 0) return NULL;
   size_t mid = length / 2;
   if (arr[mid] == value) return &arr[mid];
   if (length == 1) return NULL;
   if (arr[mid] > value) return binary_search(value, arr, mid);
   return binary_search(value, &arr[mid], length-mid);
}
