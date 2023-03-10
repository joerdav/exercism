#include "minesweeper.h"
#include <stdlib.h>
#include <string.h>

static int surround_count(const char **minefield, size_t width, size_t rows,
                          size_t x, size_t y) {
  return
      // TL
      (x > 0 && y > 0 && minefield[y - 1][x - 1] == '*') +
      // TM
      (y > 0 && minefield[y - 1][x] == '*') +
      // TR
      (x < width - 1 && y > 0 && minefield[y - 1][x + 1] == '*') +
      // ML
      (x > 0 && minefield[y][x - 1] == '*') +
      // MR
      (x < width - 1 && minefield[y][x + 1] == '*') +
      // BL
      (x > 0 && y < rows - 1 && minefield[y + 1][x - 1] == '*') +
      // BM
      (y < rows - 1 && minefield[y + 1][x] == '*') +
      // BR
      (x < width - 1 && y < rows - 1 && minefield[y + 1][x + 1] == '*');
}

char **annotate(const char **minefield, const size_t rows) {
  if (minefield == NULL || rows == 0) {
    return NULL;
  }
  char **result = malloc(sizeof(*result) * rows);
  size_t width = strlen(minefield[0]);
  for (size_t i = 0; i < rows; i++) {
    result[i] = calloc(width, sizeof(char));
    for (size_t j = 0; j < strlen(minefield[0]); j++) {
      if (minefield[i][j] == '*') {
        result[i][j] = '*';
        continue;
      }
      if (minefield[i][j] == ' ') {
        int amount = surround_count(minefield, width, rows, j, i);
        result[i][j] = amount ? '0' + amount : ' ';
      }
    }
  }
  return result;
}
void free_annotation(char **annotation) { free(annotation); }
