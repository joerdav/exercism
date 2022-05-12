#include "circular_buffer.h"

circular_buffer_t *new_circular_buffer(size_t capacity) {
   circular_buffer_t *b = malloc(sizeof(circular_buffer_t));
   *b = (circular_buffer_t) {
      .oldest = 0, .length = 0,
      .capacity = capacity,
      .buf = malloc(sizeof(int16_t)*capacity)
   };
   return b;
}

int16_t read(circular_buffer_t *buffer, buffer_value_t *value) {
   if (buffer->length == 0) {
      errno = ENODATA;
      return EXIT_FAILURE;
   }
   *value = buffer->buf[buffer->oldest];
   buffer->oldest = (buffer->oldest+1)%buffer->capacity;
   buffer->length--;
   return EXIT_SUCCESS;
}

int16_t write(circular_buffer_t *buffer, buffer_value_t value) {
   if (buffer->length==buffer->capacity) {
      errno = ENOBUFS;
      return EXIT_FAILURE;
   }
   size_t idx = buffer->oldest+buffer->length;
   buffer->length++;
   buffer->buf[idx] = value;
   return EXIT_SUCCESS;
}
int16_t overwrite(circular_buffer_t *buffer, buffer_value_t value) {
   if (buffer->length==buffer->capacity) {
      buffer->buf[buffer->oldest] = value;
      buffer->oldest = (buffer->oldest+1)%buffer->capacity;
      return EXIT_SUCCESS;
   }
   return write(buffer, value);
}

int16_t delete_buffer(circular_buffer_t *buffer) {
   free(buffer->buf);
   free(buffer);
   return EXIT_SUCCESS;
}
int16_t clear_buffer(circular_buffer_t *buffer) {
   buffer->length=0;
   buffer->oldest=0;
   return EXIT_SUCCESS;
}
