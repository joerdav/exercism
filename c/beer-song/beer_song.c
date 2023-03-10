#include "beer_song.h"
#include <stdlib.h>
#include <stdio.h>
#include <string.h>

void recite(uint8_t start_bottles, uint8_t take_down, char **song) {
   size_t line = 0;
   for (uint8_t i = start_bottles; i > start_bottles - take_down; i--) {
      if (i == 0) {
         strcpy(song[line++], "No more bottles of beer on the wall, no more bottles of beer.");
         strcpy(song[line++], "Go to the store and buy some more, 99 bottles of beer on the wall.");
         return;
      } else {
      sprintf(song[line++], "%d bottle%s of beer on the wall, %d bottle%s of beer.",
            i, 
            i==1?"":"s",
            i,
            i==1?"":"s");
      if (i == 1)
         strcpy(song[line++], "Take it down and pass it around, no more bottles of beer on the wall.");
      else
         sprintf(song[line++], "Take one down and pass it around, %d bottle%s of beer on the wall.", 
               i-1,
               i==2?"":"s");
      }
      strcpy(song[line++], "");
   }
}
