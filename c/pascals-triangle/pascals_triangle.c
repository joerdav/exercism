#include "pascals_triangle.h"
#include <stdlib.h>

#define MIN_SIZE(nrows) ((nrows) > 0 ? (nrows) : 1)

void free_triangle(uint8_t **triangle, size_t rows) {
   for (size_t i = 0; i < MIN_SIZE(rows); i++) {
      free(triangle[i]);
   }
   free(triangle);
}
uint8_t **create_triangle(size_t rows) {
   uint8_t **triangle = malloc(sizeof(uint8_t*) * MIN_SIZE(rows));
   if (rows == 0) {
      triangle[0] = calloc(1, sizeof(uint8_t));
      return triangle;
   }
   triangle = malloc(sizeof(uint8_t*)*rows);
   for (size_t i = 0; i < rows; i++) {
      triangle[i] = calloc(rows, sizeof(uint8_t));
      for (size_t j = 0; j < i+1; j++) {
         if (j == 0) {
            triangle[i][j] = 1;
            continue;
         }
         triangle[i][j] = triangle[i-1][j] + triangle[i-1][j-1];
      }
   }
   return triangle;
}
