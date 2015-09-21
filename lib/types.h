#ifndef TYPES_H
#define TYPES_H

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

#endif
