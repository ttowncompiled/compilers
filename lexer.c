// a lexical analyzer for a simplified version of Pascal.
#include <stdio.h>
#include <stdlib.h>
#include "lexer.h"

int const MAX_BUFFER_SIZE = 72;
int const MAX_ID_SIZE = 10;

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

  ReservedWordList* reserved = load_reserved_words();
  printf("%s\n", reserved->word->value);
  char* filename = argv[1];
  LineList* head = analyze(filename);
  return print_listing_file(head);
}

ReservedWordList* load_reserved_words() {
  FILE* file;
  if((file = fopen("reserved_words.txt", "r")) == NULL) {
    printf("Cannot open file reserved_words.txt");
    exit(1);
  }
  ReservedWordList* head = malloc(sizeof(ReservedWordList));
  ReservedWordList* node = head;
  size_t buffer_size = (MAX_BUFFER_SIZE+1) * sizeof(char);
  char* buffer = malloc(buffer_size);

  while (-1 != getline(&buffer, &buffer_size, file) && buffer[0] == '"') {
    char c;
    char* value = malloc((MAX_ID_SIZE+1) * sizeof(char));
    int type = 0;
    int attr = 0;
    int hare = 0;
    int trts = 0;
    
    // move past the first "
    hare++;
    // grab the value of the reserved word
    while ((c = buffer[hare++]) != '"') {
      value[trts++] = c;
    }
    value[trts] = '\0';
    // move the past the space
    hare++;
    // grab the type of the reserved word
    while ((c = buffer[hare++]) != ' ') {
      type = (type*10) + (c-48);
    }
    // grab the attr of the reserved word
    while ((c = buffer[hare++]) != '\n') {
      attr = (attr*10) + (c-48);
    }

    ReservedWord* word = malloc(sizeof(ReservedWord));
    ReservedWordList* next = malloc(sizeof(ReservedWordList));

    word->value = value;
    word->type = type;
    word->attr = attr;
    node->word = word;
    node->next = next;

    node = node->next;
  }

  return head;
}

LineList* analyze(char* filename) {
  FILE* file;
  if ((file = fopen(filename, "r")) == NULL) {
    printf("Cannot open file %s\n", filename);
    exit(1);
  }
  LineList* head = malloc(sizeof(LineList));
  LineList* node = head;
  size_t buffer_size = ((MAX_BUFFER_SIZE+1) * sizeof(char));
  char* buffer = malloc(buffer_size);
  int line_number = 0;

  while (-1 != getline(&buffer, &buffer_size, file)) {
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

