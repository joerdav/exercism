#include "palindrome_products.h"
#include <stdlib.h>

product_t *get_palindrome_product(int from, int to);
void free_product(product_t *p) {
  free(p->factors_lg);
  free(p->factors_sm);
  free(p);
}
