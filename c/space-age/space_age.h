#ifndef SPACE_AGE_H
#define SPACE_AGE_H

#include <stdint.h>

typedef float planet_t;

#define EARTH (planet_t)31557600.0F
#define MERCURY (planet_t)0.2408467F*EARTH
#define VENUS (planet_t)0.61519726F*EARTH
#define MARS (planet_t)1.8808158F*EARTH
#define JUPITER (planet_t)11.862615F*EARTH
#define SATURN (planet_t)29.447498F*EARTH
#define URANUS (planet_t)84.016846F*EARTH
#define NEPTUNE (planet_t)164.79132*EARTH

float age(planet_t planet, int64_t seconds);

#endif
