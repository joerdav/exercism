#include "darts.h"

uint8_t score(coordinate_t coordinate) {
   float distance = sqrtf(fabs(pow(coordinate.x,2)+pow(coordinate.y,2)));
   if (distance > 10.0F) return 0;
   if (distance > 5.0F) return 1;
   if (distance > 1.0F) return 5;
   return 10;
}
