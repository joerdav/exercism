#include "circular_buffer.h"

static size_t normalize_idx(size_t idx, size_t cap) {
   return (idx+cap)%cap;
}
circular_buffer_t* new_circular_buffer(size_t cap) {
   circular_buffer_t *b = calloc(1, sizeof(circular_buffer_t));
   if (!b) {
      return NULL;
   }
   b->cap = cap;
   b->buf = malloc(cap * sizeof(buffer_value_t));
   if (!b->buf) {
      free(b);
      return NULL;
   }
   return b;
}
void delete_buffer(circular_buffer_t *buffer) {
   free(buffer->buf);
   free(buffer);
}
void clear_buffer(circular_buffer_t *buffer) {
   buffer->len = 0;
}
int16_t read(circular_buffer_t *buffer, buffer_value_t *read_value) {
   if (buffer == NULL) {
      errno = EINVAL;
      return EXIT_FAILURE;
   }
   if (buffer->len == 0) {
      errno = ENODATA;
      return EXIT_FAILURE;
   }
   size_t read_idx = normalize_idx(buffer->idx-buffer->len, buffer->cap);
   *read_value = buffer->buf[read_idx];
   buffer->len--;
   return 0;
}
int16_t write(circular_buffer_t *buffer, buffer_value_t value) {
   if (buffer == NULL) {
      errno = EINVAL;
      return EXIT_FAILURE;
   }
   if (buffer->len == buffer->cap) {
      errno = ENOBUFS;
      return EXIT_FAILURE;
   }
   return overwrite(buffer, value);
}
int16_t overwrite(circular_buffer_t *buffer, buffer_value_t value) {
   if (buffer == NULL) {
      errno = ENODATA;
      return EXIT_FAILURE;
   }
   buffer->buf[buffer->idx] = value;
   buffer->idx = normalize_idx(buffer->idx+1, buffer->cap);
   if (buffer->len != buffer->cap)
      buffer->len++;
   return 0;
}
