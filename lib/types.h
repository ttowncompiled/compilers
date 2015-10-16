#ifndef TYPES_H
#define TYPES_H

const int MAX_ID_LENGTH = 10;
size_t const MAX_ID_SIZE = ((MAX_ID_LENGTH+1) * sizeof(char));

typedef struct Error {
  int line_number;
  char* reason;
} Error;

typedef struct ErrorNode {
  Error* error;
  struct ErrorNode* next;
} ErrorNode;

typedef struct Line {
  char* value;
  int number;
} Line;

typedef struct LineNode {
  Line* line;
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

#endif
