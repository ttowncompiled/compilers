#ifndef TYPES_UTIL_H
#define TYPES_UTIL_H
#include "global.h"
#include "types.h"
#include "util.h"

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

ReservedWordNode* new_reserved_word_node(char* value, int type, int attr) {
  ReservedWordNode* node = malloc(sizeof(ReservedWordNode));
  node->word = malloc(sizeof(ReservedWord));
  node->word->value = value;
  node->word->type = type;
  node->word->attr = attr;
  node->next = NULL;
  return node;
}

ReservedWordNode* reserved_word_node_from(char* entry) {
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
  return new_reserved_word_node(value, type, attr);
}

LineNode* new_line_node(int line_number, char* value) {
  LineNode* node = malloc(sizeof(LineNode));
  node->line = malloc(sizeof(Line));
  node->line->value = value;
  node->line->number = line_number;
  node->next = NULL;
  return node;
}

Token* new_token(int line_number, char* lexeme, int type, int attr) {
  Token* token = malloc(sizeof(Token));
  token->line_number = line_number;
  token->lexeme = lexeme;
  token->type = type;
  Attribute attribute;
  attribute.value = attr;
  token->attr = attribute;
  return token;
}

TokenNode* new_token_node(Token* token) {
  TokenNode* node = malloc(sizeof(TokenNode));
  node->token = token;
  node->next = NULL;
  return node;
}

SymbolNode* new_symbol_node(char* symbol) {
  SymbolNode* node = malloc(sizeof(SymbolNode));
  node->symbol = symbol;
  node->next = NULL;
  return node;
}

SymbolNode* save_symbol(char* symbol) {
  if (symbol_table == NULL) {
    symbol_table = new_symbol_node(symbol);
    return symbol_table;
  }
  SymbolNode* node = symbol_table;
  while (node->next != NULL) {
    if (is_equal(node->symbol, symbol)) {
      return node;
    }
    node = node->next;
  }
  if (is_equal(node->symbol, symbol)) {
    return node;
  }
  node->next = new_symbol_node(symbol);
  return node->next;
}

#endif
