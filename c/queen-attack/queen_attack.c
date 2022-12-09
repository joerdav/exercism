#include "queen_attack.h"
#include <stdbool.h>
#include <stdlib.h>

static bool valid_pos(position_t pos) {
   return pos.column < 8 && pos.row < 8;
}
attack_status_t can_attack(position_t queen_1, position_t queen_2) {
   if (queen_1.column == queen_2.column
         && queen_1.row == queen_2.row) return INVALID_POSITION;
   if (!valid_pos(queen_1) || !valid_pos(queen_2)) return INVALID_POSITION;
   if (queen_1.column == queen_2.column || queen_1.row == queen_2.row) return CAN_ATTACK;
   int row_diff =abs((int)queen_1.row - (int)queen_2.row);
   int col_diff =abs((int)queen_1.column - (int)queen_2.column);
   if (row_diff == col_diff) return CAN_ATTACK;
   return CAN_NOT_ATTACK;
}
