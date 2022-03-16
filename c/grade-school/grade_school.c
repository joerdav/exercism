#include "grade_school.h"

void init_roster(roster_t *r) {
   *r = (roster_t){0};
}

static int cmp_students(const void *a, const void *b) {
   student_t *sa = (student_t*)a;
   student_t *sb = (student_t*)b;
   if (sa->grade == sb->grade) return strcmp(sa->name, sb->name);
   return sa->grade-sb->grade;
}


bool add_student(roster_t *r, char name[MAX_NAME_LENGTH], uint8_t grade) {
   for (size_t i = 0; i < r->count; ++i)
      if (strcmp(r->students[i].name, name) == 0)
         return false;
   student_t s;
   s.grade = grade;
   strcpy(s.name, name);
   r->students[r->count++] = s;
   qsort(r->students, r->count, sizeof(student_t), cmp_students);
   return true;
}

roster_t get_grade(roster_t *r, uint8_t grade) {
   roster_t res;
   init_roster(&res);
   for (size_t i = 0; i < r->count; i++)
      if (r->students[i].grade == grade) add_student(&res, r->students[i].name, grade);
   return res;
}
