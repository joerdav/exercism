#include "word_count.h"
#include <ctype.h>
#include <stdbool.h>
#include <stdint.h>
#include <string.h>

static bool is_letter(const char ch, bool in_word) {
   return isalnum(ch) || (in_word && ch == '\'');
}

static size_t mark_word(char word[MAX_WORD_LENGTH], word_count_word_t *words, size_t count) {
   for (size_t i = 0; i < count; i++) {
      if (strcmp(word, words[i].text) == 0) {
         words[i].count++;
         return count;
      }
   }
   words[count] = (word_count_word_t) {0};
   strcpy(words[count].text, word);
   words[count].count = 1;
   return count+1;
}

int count_words(const char *sentence, word_count_word_t *words) {
   if (!sentence || !words) return -1;
   size_t uniq = 0;
   size_t count = 0;
   char word[MAX_WORD_LENGTH+1] = {0};
   int word_idx = 0;
   bool in_word = false;
   for (; *sentence && *sentence != '\0'; sentence++) {
      if (!is_letter(*sentence, in_word)) continue;
      in_word = true;
      if (word_idx+1 > MAX_WORD_LENGTH) return EXCESSIVE_LENGTH_WORD;
      word[word_idx++] = tolower(*sentence);
      if (!sentence[1] || !is_letter(sentence[1], in_word)) {
         if (++count > MAX_WORDS) return EXCESSIVE_NUMBER_OF_WORDS;
         if (word[word_idx-1] == '\'') word[word_idx-1] = '\0';
         uniq = mark_word(word, words, uniq);
         memset(word, 0, sizeof(word));
         word_idx = 0;
         in_word = false;
      }
   }
   return uniq;
}
