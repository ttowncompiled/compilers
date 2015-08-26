// the header file for the lexical analyzer
#include <stdio.h>
#include <stdlib.h>

#define BUFFER_SIZE 72

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
  WHILE
};

typedef struct Line {
  char* value;
  struct Line* next;
} Line;

typedef struct Word {
  char* value;
  int type;
  int attr;
  struct Word* next;
} Word;

Word* loadReservedWords();
int printListingFile(Line* head);

