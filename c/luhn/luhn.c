#include "luhn.h"

bool luhn(const char *num) {
  int *nums = calloc(strlen(num), sizeof(*nums));
  int len = 0;
  for (; *num; num++) {
    if (!isdigit(*num) && *num != ' ')
      return false;
    if (isdigit(*num))
      nums[len++] = *num - '0';
  }
  if (len < 2)
    return false;
  int back_idx = 1;
  int sum = 0;
  for (int i = len - 1; i >= 0; i--) {
    if (!(back_idx & 1)) {
      nums[i] = nums[i] * 2;
      if (nums[i] > 9)
        nums[i] = nums[i] - 9;
    }
    sum += nums[i];
    back_idx++;
  }
  free(nums);
  return sum % 10 == 0;
}
