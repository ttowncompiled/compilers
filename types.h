#ifndef TYPES_H
#define TYPES_H

enum types {
  ENDFILE = -1,
  NIL,
  ADDOP,
  AND,
  ARRAY,
  ASSIGNOP,
  BEGIN,        // 5
  CLOSEBRACKET,
  CLOSEPAREN,
  COLON,
  COMMA,
  DIV,          // 10
  DO,
  ELSE,
  END,
  FUNCTION,
  ID,           // 15
  IF,
  INTEGER,
  MINUS,
  MOD,
  MULOP,        // 20
  NOT,
  NUM,
  OF,
  OPENBRACKET,
  OPENPAREN,    // 25
  OR,
  PERIOD,
  PLUS,
  PROCEDURE,
  PROGRAM,      // 30
  RANGE,
  REAL,
  RELOP,
  SEMICOLON,
  THEN,         // 35
  VAR,
  WHILE,
  LEXERR
};

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

char* type_annotation_of(int type) {
  switch (type) {
    case -1: return "(EOF)";
    case 0: return "(NULL)";
    case 1: return "(ADDOP)";
    case 2: return "(MULOP)";
    case 3: return "(ARRAY)";
    case 4: return "(ASSIGNOP)";
    case 5: return "(BEGIN)";
    case 6: return "(CLOSEBRACKET)";
    case 7: return "(CLOSEPAREN)";
    case 8: return "(COLON)";
    case 9: return "(COMMA)";
    case 10: return "(MULOP)";
    case 11: return "(DO)";
    case 12: return "(ELSE)";
    case 13: return "(END)";
    case 14: return "(FUNCTION)";
    case 15: return "(ID)";
    case 16: return "(IF)";
    case 17: return "(INTEGER)";
    case 18: return "(MINUS)";
    case 19: return "(MULOP)";
    case 20: return "(MULOP)";
    case 21: return "(NOT)";
    case 22: return "(NUM)";
    case 23: return "(OF)";
    case 24: return "(OPENBRACKET)";
    case 25: return "(OPENPAREN)";
    case 26: return "(ADDOP)";
    case 27: return "(PERIOD)";
    case 28: return "(PLUS)";
    case 29: return "(PROCEDURE)";
    case 30: return "(PROGRAM)";
    case 31: return "(RANGE)";
    case 32: return "(REAL)";
    case 33: return "(RELOP)";
    case 34: return "(SEMICOLON)";
    case 35: return "(THEN)";
    case 36: return "(VAR)";
    case 37: return "(WHILE)";
    case 38: return "(LEXERR)";
  }
  return NULL;
}

char* attr_annotation_of(int attr) {
  if (attr < -1) {
    return "(PTR)";
  }
  switch(attr) {
    case 0: return "(NULL)";
  }
  return NULL;
}

#endif
