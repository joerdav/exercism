#ifndef ROBOT_SIMULATOR_H
#define ROBOT_SIMULATOR_H

typedef enum {
   DIRECTION_NORTH = 0,
   DIRECTION_EAST,
   DIRECTION_SOUTH,
   DIRECTION_WEST,
   DIRECTION_DEFAULT = DIRECTION_NORTH,
   DIRECTION_MAX
} robot_direction_t;

typedef struct {
   int x;
   int y;
} robot_position_t;

typedef struct {
   robot_direction_t direction;
   robot_position_t position;
} robot_status_t;

robot_status_t robot_create(robot_direction_t direction, int x, int y);
void robot_move(robot_status_t *robot, const char *commands);


#define NORTH_TRANSFORM (robot_position_t) {0,1}
#define EAST_TRANSFORM (robot_position_t) {1,0}
#define SOUTH_TRANSFORM (robot_position_t) {0,-1}
#define WEST_TRANSFORM (robot_position_t) {-1,0}

#endif
