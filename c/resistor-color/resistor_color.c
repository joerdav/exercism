#include "resistor_color.h"
#include <stdlib.h>

int color_code(resistor_band_t b) {
   return (int)b;
}

resistor_band_t colorsArray[] = { COLORS };

resistor_band_t* colors() {
   return colorsArray;
}
