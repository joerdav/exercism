#include "resistor_color_duo.h"

uint16_t color_code(resistor_band_t *resistors) {
   return resistors[0]*10 + resistors[1];
}
