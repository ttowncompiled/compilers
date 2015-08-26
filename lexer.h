// the header file for the lexical analyzer
#include <stdio.h>
#include <stdlib.h>

enum types {
  ADDOP = 1,
  ARRAY,
  ASSIGNOP,
  BEGIN,
  CLOSEBRACKET,
  CLOSEPAREN,
  COLON,
  COMMA,
  DO,
  ELSE,
  END,
  FUNCTION,
  ID,
  IF,
  INTEGER,
  MINUS,
  MULOP,
  NOT,
  NUM,
  OF,
  OPENBRACKET,
  OPENPAREN,
  PERIOD,
  PLUS,
  PROCEDURE,
  PROGRAM,
  RANGE,
  REAL,
  RELOP,
  SEMICOLON,
  THEN,
  VAR,
  WHILE
};

void printListingFile(char* filename);

