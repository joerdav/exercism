#include "robot_simulator.h"
#include <stdbool.h>

static robot_direction_t rotate(robot_direction_t direction, bool left) {
   if (left) return (direction + 3)%4;
   return (direction + 1)%4;
}
robot_status_t robot_create(robot_direction_t direction, int x, int y) {
   return (robot_status_t) {direction, {x, y}};
}
void robot_move(robot_status_t *robot, const char *commands) {
   for (const char *command = commands; *command; ++command) {
      switch (*command) {
         case 'R': 
            robot->direction = rotate(robot->direction, false);
            break;
         case 'L': 
            robot->direction = rotate(robot->direction, true);
            break;
         case 'A': {
               robot_position_t transform;
               if (robot->direction == DIRECTION_NORTH) transform = NORTH_TRANSFORM;
               if (robot->direction == DIRECTION_EAST) transform = EAST_TRANSFORM;
               if (robot->direction == DIRECTION_SOUTH) transform = SOUTH_TRANSFORM;
               if (robot->direction == DIRECTION_WEST) transform = WEST_TRANSFORM;
               robot->position.x += transform.x;
               robot->position.y += transform.y;
               break;
         }
      }
   }
}
