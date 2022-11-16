#include "armstrong_numbers.h"
#include <math.h>
#include <stdio.h>

static int count_digits(int n)
{
    if (n < 10) return 1;
    return 1+count_digits(n/10);
}

static int sum_of_digits_powered(int candidate)
{
    int remainder = candidate;
    int sum = 0;
    int count = count_digits(candidate);
    printf("Count: %d\n", sum);
    while (remainder != 0)
    {
        sum += pow(remainder%10, count);
        remainder /= 10;
    }
    printf("Result: %d\n", sum);
    return sum;
}

bool is_armstrong_number(int candidate)
{
    return candidate == sum_of_digits_powered(candidate);
}

