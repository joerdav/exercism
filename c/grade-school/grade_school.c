#include "grade_school.h"

static int compare_students(const void *a, const void *b) {
   const student_t *sa = a;
   const student_t *sb = b;
   int g = sa->grade - sb->grade;
   if (g != 0) return g;
   return strcmp(sa->name, sb->name);
}
static void sort(roster_t *roster) {
   qsort(roster->students, roster->count, sizeof(student_t), compare_students);
}
static bool in_roster(roster_t *roster, student_t s) {
   for (size_t i = 0; i < roster->count; i++)
      if (strcmp(roster->students[i].name, s.name) == 0) return true;
   return false;
}
void init_roster(roster_t *roster) {
   *roster = (roster_t){ 0 };
}
bool add_student(roster_t *roster, const char *name, uint8_t grade) {
   if (roster->count == MAX_STUDENTS || strlen(name) > MAX_NAME_LENGTH) return false;
   student_t new_student = { .grade = grade };
   strncpy(new_student.name, name, MAX_NAME_LENGTH);
   if (in_roster(roster, new_student)) return false;
   roster->students[roster->count++] = new_student;
   sort(roster);
   return true;
}
roster_t get_grade(const roster_t *roster, uint8_t grade) {
   roster_t g = {0};
   for (size_t i = 0; i < roster->count; i++) {
      if (roster->students[i].grade == grade) {
         g.students[g.count++] = roster->students[i];
      }
   }
   return g;
}

