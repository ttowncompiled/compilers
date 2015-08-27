#ifndef TYPES_H
#define TYPES_H

enum types {
  ADDOP = 1,
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

#endif
