#include "resistor_color.h"

int color_code(resistor_band_t r) {
   return r;

}

int * colors() {
   static int r[] = {
      BLACK, BROWN, RED, ORANGE, YELLOW,
      GREEN, BLUE, VIOLET, GREY, WHITE
   };
   return r;
}
