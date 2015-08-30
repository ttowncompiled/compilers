#ifndef TYPES_H
#define TYPES_H

enum type {
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
  MOD,
  MULOP,
  NOT,          // 20
  NUM,
  OF,
  OPENBRACKET,
  OPENPAREN,
  OR,           // 25
  PERIOD,
  PROCEDURE,
  PROGRAM,
  RANGE,
  REAL,         // 30
  RELOP,
  SEMICOLON,
  THEN,
  VAR,
  WHILE,        // 35
  LEXERR
};

enum num {
  INT = 1,
  REAL_,
  LREAL
};

enum relop {
  EQ = 1,
  NEQ,
  LT,
  LTE,
  GTE,          // 5
  GT
};

enum mulop {
   MULT = 1,
   DIVISION,
   DIV_,
   MOD_,
   AND_         // 5
};

enum addop {
  PLUS = 1,
  MINUS,
  OR_
};

enum lexerr {
  UNRECOG = 1,
  ID_LENGTH,
  INT_LENGTH,
  XX_LENGTH,
  YY_LENGTH,    // 5
  ZZ_LENGTH
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
    case MOD: return "(MULOP)";
    case MULOP: return "(MULOP)";
    case NOT: return "(NOT)";
    case NUM: return "(NUM)";
    case OF: return "(OF)";
    case OPENBRACKET: return "(OPENBRACKET)";
    case OPENPAREN: return "(OPENPAREN)";
    case OR: return "(ADDOP)";
    case PERIOD: return "(PERIOD)";
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

char* attr_annotation_of(int type, int attr) {
  if (attr < -1) {
    return "(PTR)";
  }
  if (type == NUM) {
    switch (attr) {
      case NIL: return "(NULL)";
      case INT: return "(INT)";
      case REAL_: return "(REAL)";
      case LREAL: return "(LREAL)";
    }
  } else if (type == RELOP) {
    switch (attr) {
      case NIL: return "(NULL)";
      case EQ: return "(EQ)";
      case NEQ: return "(NEQ)";
      case LT: return "(LT)";
      case LTE: return "(LTE)";
      case GTE: return "(GTE)";
      case GT: return "(GT)";
    }
  } else if (type == MULOP) {
    switch (attr) {
      case NIL: return "(NULL)";
      case MULT: return "(MULT)";
      case DIVISION: return "(DVSN)";
      case DIV_: return "(DIV)";
      case MOD_: return "(MOD)";
      case AND_: return "(AND)";
    }
  } else if (type == ADDOP) {
    switch (attr) {
      case NIL: return "(NULL)";
      case PLUS: return "(PLUS)";
      case MINUS: return "(MINUS)";
      case OR_: return "(OR)";
    }
  } else if (type == LEXERR) {
    switch (attr) {
      case NIL: return "(NULL)";
      case UNRECOG: return "(?Symbol)";
      case ID_LENGTH: return "(#ID>10)";
      case INT_LENGTH: return "(#INT>10)";
      case XX_LENGTH: return "(#XX>5)";
      case YY_LENGTH: return "(YY>5)";
      case ZZ_LENGTH: return "(ZZ>2)";
    }
  }
  switch (attr) {
    case NIL: return "(NULL)";
  }
  return NULL;
}

#endif
