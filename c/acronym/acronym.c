#include "acronym.h"
#include <stdlib.h>
#include <string.h>
#include <stdbool.h>
#include <ctype.h>

static bool is_letter(const char ch) {
   return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') || ch == '\'';
}

char *abbreviate(const char *phrase) {
   if (!phrase || strlen(phrase) == 0) return NULL;
   char *abbr = malloc(sizeof(char)*strlen(phrase));
   char *curr = abbr;
   for (const char *letter = phrase; *letter; letter++) {
      if (is_letter(*letter)) {
         *curr++ = toupper(*letter);
         while (letter[1] != '\0' && is_letter(letter[1])) {
            letter++;
         }
      }
   }
   *curr = '\0';
   return abbr;
}
