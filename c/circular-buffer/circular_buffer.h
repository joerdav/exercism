#ifndef CIRCULAR_BUFFER_H
#define CIRCULAR_BUFFER_H

#include <stdlib.h>
#include <stdbool.h>
#include <errno.h>

typedef int16_t buffer_value_t;

typedef struct {
	buffer_value_t *buf;
	size_t capacity;
	size_t length;
	size_t oldest;
} circular_buffer_t;

int16_t write(circular_buffer_t *buffer, buffer_value_t value);
int16_t read(circular_buffer_t *buffer, buffer_value_t *value);
int16_t overwrite(circular_buffer_t *buffer, buffer_value_t value);
circular_buffer_t *new_circular_buffer(size_t capacity);
int16_t delete_buffer(circular_buffer_t *buffer);
int16_t clear_buffer(circular_buffer_t *buffer);

#endif
