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
    case ENDFILE: return "(EOF)";
    case NIL: return "(NULL)";
    case ADDOP: return "(ADDOP)";
    case AND: return "(MULOP)";
    case ARRAY: return "(ARRAY)";
    case ASSIGNOP: return "(ASSIGNOP)";
    case BEGIN: return "(BEGIN)";
    case CLOSEBRACKET: return "(CLOSEBRACKET)";
    case CLOSEPAREN: return "(CLOSEPAREN)";
    case COLON: return "(COLON)";
    case COMMA: return "(COMMA)";
    case DIV: return "(MULOP)";
    case DO: return "(DO)";
    case ELSE: return "(ELSE)";
    case END: return "(END)";
    case FUNCTION: return "(FUNCTION)";
    case ID: return "(ID)";
    case IF: return "(IF)";
    case INTEGER: return "(INTEGER)";
    case MINUS: return "(MINUS)";
    case MOD: return "(MULOP)";
    case MULOP: return "(MULOP)";
    case NOT: return "(NOT)";
    case NUM: return "(NUM)";
    case OF: return "(OF)";
    case OPENBRACKET: return "(OPENBRACKET)";
    case OPENPAREN: return "(OPENPAREN)";
    case OR: return "(ADDOP)";
    case PERIOD: return "(PERIOD)";
    case PLUS: return "(PLUS)";
    case PROCEDURE: return "(PROCEDURE)";
    case PROGRAM: return "(PROGRAM)";
    case RANGE: return "(RANGE)";
    case REAL: return "(REAL)";
    case RELOP: return "(RELOP)";
    case SEMICOLON: return "(SEMICOLON)";
    case THEN: return "(THEN)";
    case VAR: return "(VAR)";
    case WHILE: return "(WHILE)";
    case LEXERR: return "(LEXERR)";
  }
  return NULL;
}

char* attr_annotation_of(int attr) {
  if (attr < -1) {
    return "(PTR)";
  }
  switch (attr) {
    case NIL: return "(NULL)";
  }
  return NULL;
}

#endif
