#ifndef LEXER_H 
#define LEXER_H
#include "types.h"
#include "util.h"

const int MAX_BUFFER_LENGTH = 72;
const int MAX_ID_LENGTH = 10;
const int MAX_INTEGER_LENGTH = 10;
const int MAX_XX_LENGTH = 5;
const int MAX_YY_LENGTH = 5;
const int MAX_ZZ_LENGTH = 5;

size_t const MAX_BUFFER_SIZE = ((MAX_BUFFER_LENGTH+1) * sizeof(char));
size_t const MAX_ID_SIZE = ((MAX_ID_LENGTH+1) * sizeof(char));

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

void check_buffer_size(size_t buffer_size, LineNode* node) {
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

LineNode* organize(char* filename);

int print_token_file(TokenNode* head);

int print_listing_file(LineNode* head);

TokenNode* analyze(LineNode* first, ReservedWordNode* reserved);

int white_space_machine(LineNode* node, int* trts);
Token* id_machine(LineNode* node, ReservedWordNode* reserved,
    SymbolNode* symbols, int* trts);
Token* long_real_machine(LineNode* node, int* trts);
Token* real_machine(LineNode* node, int* trts);
Token* int_machine(LineNode* node, int* trts);
Token* relop_machine(LineNode* node, int* trts);
Token* addop_machine(LineNode* node, int* trts);
Token* mulop_machine(LineNode* node, int* trts);
Token* assignop_machine(LineNode* node, int* trts);
Token* catchall_machine(LineNode* node, int* trts);

#endif
