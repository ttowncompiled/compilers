#ifndef GLOBAL_H
#define GLOBAL_H
#include <string.h>
#include <ctype.h>
#include "enums.h"
#include "types.h"

const int MAX_BUFFER_LENGTH = 72;
size_t const MAX_BUFFER_SIZE = ((MAX_BUFFER_LENGTH+1) * sizeof(char));
char* const LINE_TOO_LONG = "ERROR: Lines can be only 72 characters long.\n";

const int MAX_INTEGER_LENGTH = 10;
const int MAX_XX_LENGTH = 5;
const int MAX_YY_LENGTH = 5;
const int MAX_ZZ_LENGTH = 2;

ReservedWordNode* reserved_word_table = NULL;
ErrorNode* error_list = NULL;
SymbolNode* symbol_table = NULL;

char* substring(char* string, int first, int last) {
  if (last <= first) {
    return NULL;
  }
  char* sub = malloc((last-first+1) * sizeof(char));
  int idx = 0;
  while (idx < last-first) {
    sub[idx] = string[first+idx];
    idx++;
  }
  sub[idx] = '\0';
  return sub;
}

SymbolNode* save_symbol(char* symbol) {
  if (symbol_table == NULL) {
    symbol_table = new_symbol_node(symbol);
    return symbol_table;
  }
  SymbolNode* node = symbol_table;
  while (node->next != NULL) {
    if (strcmp(node->symbol, symbol) == 0) {
      return node;
    }
    node = node->next;
  }
  if (strcmp(node->symbol, symbol) == 0) {
    return node;
  }
  node->next = new_symbol_node(symbol);
  return node->next;
}

void throw_error(int line_number, char* reason) {
  ErrorNode* node = malloc(sizeof(ErrorNode));
  node->error = malloc(sizeof(Error));
  node->error->line_number = line_number;
  node->error->reason = reason;
  node->next = NULL;
  if (error_list == NULL) {
    error_list = node;
    return;
  }
  ErrorNode* curr = error_list;
  while (curr->next != NULL) {
    curr = curr->next;
  }
  curr->next = node;
}

#endif
