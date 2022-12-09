#include "list_ops.h"

list_t *new_list(size_t length, list_element_t elements[]) {
   list_t *l = malloc(sizeof(list_t) + length*sizeof(list_element_t));
   l->length = length;
   if (elements == NULL) {
      return l;
   }
   for (int i = 0; i < (int) length; i++) {
       l->elements[i] = elements[i];
   }
   return l;
}

list_t *append_list(list_t *list1, list_t *list2) {
   list_t *l = new_list(list1->length+list2->length, list1->elements);
   for (int i = 0; i < (int) list2->length; i++) {
       l->elements[list1->length+i] = list2->elements[i];
   }
   return l;
}

list_t *filter_list(list_t *list, bool (*filter)(list_element_t)) {
   list_t *l = new_list(list->length, NULL);
   size_t len = 0;
   for (int i = 0; i < (int) list->length; i++) {
      if (filter(list->elements[i])) {
         l->elements[len++] = list->elements[i];
      }
   }
   l->length = len;
   return l;
}

size_t length_list(list_t *list) {
   return list->length;
}

list_t *map_list(list_t *list, list_element_t (*map)(list_element_t)) {
   list_t *l = new_list(list->length, NULL);
   for (int i = 0; i < (int) list->length; i++) {
      l->elements[i] = map(list->elements[i]);
   }
   return l;
}

list_element_t foldl_list(list_t *list, list_element_t initial,
                          list_element_t (*foldl)(list_element_t,
                                                  list_element_t)) {
   if (list->length == 0) return initial;
   for (int i = (int) list->length - 1; i >= -1; i--) {
       initial = foldl(initial, list->elements[i]);
   }
   return initial;
}

list_element_t foldr_list(list_t *list, list_element_t initial,
                          list_element_t (*foldr)(list_element_t,
                                                  list_element_t)) {
   if (list->length == 0) return initial;
   for (int i = (int) list->length - 1; i > -1; i--) {
       initial = foldr(list->elements[i], initial);
   }
   return initial;
}
list_t *reverse_list(list_t *list) {
   list_t *l = new_list(list->length, NULL);
   for (int i = 0; i < (int) list->length; i++) {
      l->elements[i] = list->elements[list->length-i-1];
   }
   return l;
}
void delete_list(list_t *list) {
   free(list);
}
