#ifndef LEXER_H 
#define LEXER_H

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

typedef struct LineList {
  Line* line;
  struct LineList* next;
} LineList;

typedef struct ReservedWord {
  char* value;
  int type;
  int attr;
} ReservedWord;

typedef struct ReservedWordList {
  ReservedWord* word;
  struct ReservedWordList* next;
} ReservedWordList;

ReservedWordList* load_reserved_words();
LineList* analyze(char* filename);
int print_listing_file(LineList* head);
#endif

