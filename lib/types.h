#ifndef TYPES_H
#define TYPES_H
#include "util.h"

const int MAX_ID_LENGTH = 10;
size_t const MAX_ID_SIZE = ((MAX_ID_LENGTH+1) * sizeof(char));

typedef struct Line {
  char* value;
  int number;
} Line;

typedef struct LineNode {
  Line* line;
  struct LineNode* error;
  struct LineNode* next;
} LineNode;

typedef struct ReservedWord {
  char* value;
  int type;
  int attr;
} ReservedWord;

typedef struct ReservedWordNode {
  ReservedWord* word;
  struct ReservedWordNode* next;
} ReservedWordNode;

typedef struct SymbolNode {
  char* symbol;
  struct SymbolNode* next;
} SymbolNode;

typedef union Attribute {
  int value;
  SymbolNode* address;
} Attribute;

typedef struct Token {
  int line_number;
  char* lexeme;
  int type;
  Attribute attr;
} Token;

typedef struct TokenNode {
  Token* token;
  struct TokenNode* next;
} TokenNode;

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

ReservedWordNode* word_node_from(char* entry) {
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

Token* token_of(int line_number, char* lexeme, int type, int attr) {
  Token* token = malloc(sizeof(Token));
  token->line_number = line_number;
  token->lexeme = lexeme;
  token->type = type;
  Attribute attribute;
  attribute.value = attr;
  token->attr = attribute;
  return token;
}

TokenNode* token_node_with(TokenNode* node, Token* token) {
  node->token = token;
  node->next = NULL;
  return node;
}

SymbolNode* new_symbol_table() {
  SymbolNode* node = malloc(sizeof(SymbolNode));
  node->symbol = NULL;
  node->next = NULL;
  return node;
}

SymbolNode* symbol_node_with(SymbolNode* node, char* symbol) {
  node->symbol = symbol;
  node->next = NULL;
  return node;
}

SymbolNode* save_symbol(SymbolNode* symbols, char* symbol) {
  if (symbols->symbol == NULL) {
    return symbol_node_with(symbols, symbol);
  }
  while (symbols->next != NULL) {
    if (is_equal(symbol, symbols->symbol)) {
      return symbols;
    }
    symbols = symbols->next;
  }
  if (is_equal(symbol, symbols->symbol)) {
    return symbols;
  }
  symbols->next = malloc(sizeof(SymbolNode));
  return symbol_node_with(symbols->next, symbol);
}

#endif
