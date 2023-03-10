#ifndef ROMAN_NUMERALS_H
#define ROMAN_NUMERALS_H

#define NUMERAL_LEN 10
#define MAX_BUFFER_LEN 10

char *to_roman_numeral(unsigned int number);

typedef struct {
	unsigned int number;
	char numeral[NUMERAL_LEN];
} numeral_t;

#endif
