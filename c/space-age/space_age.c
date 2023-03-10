#include "space_age.h"

float age(planet_t planet, int64_t seconds) {
   if (
         planet != EARTH
         && planet != MERCURY 
         && planet != VENUS 
         && planet != MARS 
         && planet != JUPITER 
         && planet != SATURN 
         && planet != URANUS 
         && planet != NEPTUNE 
      ) return -1.0F;
   return (float)seconds/planet;
}
