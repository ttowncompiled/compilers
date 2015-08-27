#include <stdio.h>
#include <stdlib.h>
#include "lexer.h"

int main(int argc, char* argv[]) {
  if (argc > 2) {
    printf("This compiler only accepts one source file.\n");
    exit(1);
  }
  if (argc < 2) {
    printf("Please provide the name of a source file.\n");
    return 0;
  }
  load_reserved_words();
  char* filename = argv[1];
  LineNode* head = organize(filename);
  return print_listing_file(head);
}

ReservedWordNode* load_reserved_words() {
  FILE* file;
  if((file = fopen("reserved_words.txt", "r")) == NULL) {
    printf("Cannot open file reserved_words.txt");
    exit(1);
  }
  ReservedWordNode* head = malloc(sizeof(ReservedWordNode));
  size_t buffer_size = MAX_BUFFER_SIZE;
  char* buffer = malloc(buffer_size);
  if (-1 == getline(&buffer, &buffer_size, file) || buffer[0] != '"') {
    return head;
  }
  head = parse_reserved_word_entry(buffer);
  ReservedWordNode* prev = head;
  while (-1 != getline(&buffer, &buffer_size, file) && buffer[0] == '"') {
    ReservedWordNode* curr = parse_reserved_word_entry(buffer);
    prev->next = curr;
    prev = curr;
  }
  return head;
}

ReservedWordNode* parse_reserved_word_entry(char* entry) {
  char c;
  char* value = malloc(MAX_ID_SIZE);
  int type = 0;
  int attr = 0;
  int hare = 0;
  int trts = 0;
  hare++;
  while ((c = entry[hare++]) != '"') {
    value[trts++] = c;
  }
  value[trts] = '\0';
  hare++;
  while ((c = entry[hare++]) != ' ') {
    type = (type*10) + (c-48);
  }
  while ((c = entry[hare++]) != '\n') {
    attr = (attr*10) + (c-48);
  }
  return word_node_of(value, type, attr);
}

LineNode* organize(char* filename) {
  FILE* file;
  if ((file = fopen(filename, "r")) == NULL) {
    printf("Cannot open file %s\n", filename);
    exit(1);
  }
  LineNode* head = malloc(sizeof(LineNode));
  size_t buffer_size = MAX_BUFFER_SIZE;
  char* buffer = malloc(buffer_size);
  if (-1 == getline(&buffer, &buffer_size, file)) {
    return head;
  }
  int line_number = 1;
  head = line_node_of(buffer, line_number);
  assert_buffer_size(buffer_size, head);
  buffer = malloc(MAX_BUFFER_SIZE);
  LineNode* prev = head;
  while (-1 != getline(&buffer, &buffer_size, file)) {
    LineNode* curr = line_node_of(buffer, ++line_number);
    assert_buffer_size(buffer_size, curr);
    prev->next = curr; 
    prev = curr;
    buffer = malloc(MAX_BUFFER_SIZE);
  }
  return head;
}

int print_listing_file(LineNode* head) {
  LineNode* curr = head;
  printf("\n");
  while (curr != NULL) {
    printf("%4d.    %s", curr->line->number, curr->line->value);
    LineNode* error = curr->error;
    while (error != NULL) {
      printf("%4d.    %s", error->line->number, error->line->value);
      error = error->error;
    }
    curr = curr->next;
    printf("\n");
  }
  return 0;
}
