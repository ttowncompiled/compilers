// a lexical analyzer for a simplified version of Pascal.
#include <stdio.h>
#include <stdlib.h>
#include "lexer.h"

int main(int argc, char* argv[]) {
  // check that we have been given the right amount of input
  if (argc > 2) {
    printf("This compiler only accepts one source file.\n");
    exit(1);
  }
  if (argc < 2) {
    printf("Please provide the name of a source file.\n");
    return 0;
  }

  char* filename = argv[1];
  LineList* head = analyze(filename);
  return print_listing_file(head);
}

Word* loadReservedWords() {
  FILE* file;
  Word* head;
  int hare;

  file = fopen("../reserved_words.txt", "r");
  head = malloc(sizeof(Word));
  head->value = malloc(10 * sizeof(char));
  hare = 0;

  char c;
  Word* word;
  c = (char)fgetc(file);
  word = head;
  while (c != EOF) {
    // grab the reserved word
    if (c >= 'a' && c <= 'z') {
      word->value[hare++] = c;
      c = (char)fgetc(file);
    } else if (c >= '0' && c <= '9') {
      int type;
      int attr;
      type = c-48;
      // grab the token type
      while ((c = (char)fgetc(file)) >= '0' && c <= '9') {
        type = (type*10) + (c-48);
      }
      // skip to the attribute value
      while (c < '0' || c > '9') {
        c = (char)fgetc(file);
      }
      attr = c-48;
      // grab the attribute
      while ((c = (char)fgetc(file)) >= '0' && c <= '9') {
        attr = (attr*10) + (c-48);
      }
      word->type = type;
      word->attr = attr;
      // skip to the end of line
      while (c != '\n') {
        c = (char)fgetc(file);
      }
      word->value[hare] = '\0';
      hare = 0;
      // check if there is a next line
      if ((c = (char)fgetc(file)) != EOF && c != ' ') {
        Word* nextWord;
        nextWord = malloc(sizeof(Word));
        nextWord->value = malloc(10 * sizeof(char));
        word->next = nextWord;
        word = nextWord;
      }
    } else {
      c = (char)fgetc(file);
    }
  }

  return head;
}

LineList* analyze(char* filename) {
  size_t const MAX_BUFFER_SIZE = 72;

  FILE* file;
  if ((file = fopen(filename, "r")) == NULL) {
    printf("Cannot open file %s\n", filename);
  }
  LineList* head = malloc(sizeof(LineList));
  LineList* node = head;
  char* buffer = malloc(MAX_BUFFER_SIZE * sizeof(char));
  int line_number = 0;
  size_t actual_size = MAX_BUFFER_SIZE;

  while (-1 != getline(&buffer, &actual_size, file)) {
    line_number++;

    Line* line = malloc(sizeof(Line));
    LineList* next = malloc(sizeof(LineList));

    line->value = buffer;
    line->number = line_number;

    node->line = line;
    node->next = next;

    node = node->next;
    buffer = malloc(MAX_BUFFER_SIZE * sizeof(char));
  }

  return head;
}

int print_listing_file(LineList* head) {
  LineList* node = head;

  while (node != NULL && node->line != NULL) {
    printf("%4d.    %s\n", node->line->number, node->line->value);
    node = node->next;
  }
  
  return 0;
}

