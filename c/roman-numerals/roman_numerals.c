#include "roman_numerals.h"
#include <stdlib.h>
#include <stdint.h>
#include <string.h>

static const numeral_t numerals[] = {
   {1, "I"},
   {4, "IV"},
   {5, "V"},
   {9, "IX"},
   {10, "X"},
   {40, "XL"},
   {50, "L"},
   {90, "XC"},
   {100, "C"},
   {400, "CD"},
   {500, "D"},
   {900, "CM"},
   {1000, "M"},
};
const size_t NUMERALS_COUNT = sizeof(numerals)/sizeof(numeral_t);

static unsigned int write_numeral(char *buf, unsigned int number) {
   for (int i = NUMERALS_COUNT-1; i >= 0; i--) {
      if (number >= numerals[i].number) {
         strcat(buf, numerals[i].numeral);
         return number - numerals[i].number;
      }
   }
   return number;
}
char *to_roman_numeral(unsigned int number) {
   char* buf = calloc(MAX_BUFFER_LEN, sizeof(char));
   if (!buf) return NULL;
   while (number != 0) {
      number = write_numeral(buf, number);
   }
   return buf;
}
