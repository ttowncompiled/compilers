#ifndef LEXER_H 
#define LEXER_H
#include "types.h"

size_t const MAX_BUFFER_SIZE = (73 * sizeof(char));
size_t const MAX_ID_SIZE = (11 * sizeof(char));

char* const LINE_TOO_LONG = "ERROR: Lines can be only 72 characters long.\n";

LineNode* line_node_of(char* value, int line_number) {
  LineNode* node = malloc(sizeof(LineNode));
  node->line = malloc(sizeof(Line));
  node->line->value = value;
  node->line->number = line_number;
  node->error = NULL;
  node->next = NULL;
  return node;
}

ReservedWordNode* word_node_of(char* value, int type, int attr) {
  ReservedWordNode* node = malloc(sizeof(ReservedWordNode));
  node->word = malloc(sizeof(ReservedWord));
  node->word->value = value;
  node->word->type = type;
  node->word->attr = attr;
  node->next = NULL;
  return node;
}

void assert_buffer_size(size_t buffer_size, LineNode* node) {
  if (buffer_size <= MAX_BUFFER_SIZE) {
    return;
  }
  LineNode* error = line_node_of(LINE_TOO_LONG, node->line->number);
  LineNode* curr = node;
  while (curr->error != NULL) {
    curr = curr->error;
  }
  curr->error = error;
}

ReservedWordNode* load_reserved_words();
ReservedWordNode* word_node_from(char* entry);
LineNode* organize(char* filename);
int print_listing_file(LineNode* head);

#endif
