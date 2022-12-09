#include "clock.h"
#include <stdio.h>
#include <string.h>
#include <stdlib.h>

static int clock_minutes(clock_t clock) {
   char *hour_str = strtok(clock.text, ":");
   char *minute_str = strtok(NULL, ":");
   int hour = strtol(hour_str, NULL, 10);
   int minute = strtol(minute_str, NULL, 10);
   return hour*60+minute;
}

clock_t clock_create(int hour, int minute) {
   int minutes = (hour*60+minute);
   while (minutes<0) minutes += 1440;
   minutes %= 1440;
   clock_t c;
   sprintf(c.text, "%02d:%02d", minutes/60, minutes%60);
   return c;
}
clock_t clock_add(clock_t clock, int minute_add) {
   return clock_create(0, clock_minutes(clock)+minute_add);
}
clock_t clock_subtract(clock_t clock, int minute_subtract) {
   return clock_add(clock, -minute_subtract);
}
bool clock_is_equal(clock_t a, clock_t b) {
   return strcmp(a.text, b.text) == 0;
}
