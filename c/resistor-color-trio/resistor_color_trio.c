#include "resistor_color_trio.h"

resistor_value_t color_code(const resistor_band_t* resistors) {
   int ohms = (resistors[0]*10 + resistors[1])*pow(10, resistors[2]);
   if (ohms % 1000 == 0)
      return (resistor_value_t) { 
         ohms/1000, 
         KILOOHMS,
      };
   return (resistor_value_t) { 
      ohms, 
      OHMS,
   };
}
