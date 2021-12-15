#include "armstrong_numbers.h"
#include <stdio.h>
#include <string.h>
#include <math.h>

int calc_armstrong(int number) {
	int result = 0;
	char s[11];
	sprintf(s, "%d", number);
	int length = strlen(s)/sizeof(char);
	for (int i = 0; i < length; i++) {
		int num = s[i] - '0';
		result += pow(num, length);
	}
	return result;
}

bool is_armstrong_number(int candidate) {
	return candidate == calc_armstrong(candidate);
}
