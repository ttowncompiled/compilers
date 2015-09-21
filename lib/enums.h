#ifndef ENUMS_H
#define ENUMS_H

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
  _INT_ = 1,
  _REAL_,
  _LREAL_
};

enum relop {
  _EQ_ = 1,
  _NEQ_,
  _LT_,
  _LTE_,
  _GTE_,          // 5
  _GT_
};

enum mulop {
   _MULT_ = 1,
   _DIVISION_,
   _DIV_,
   _MOD_,
   _AND_         // 5
};

enum addop {
  _PLUS_ = 1,
  _MINUS_,
  _OR_
};

enum lexerr {
  UNREC = 1,
  ID_LENGTH,
  INT_LENGTH,
  XX_LENGTH,
  YY_LENGTH,    // 5
  ZZ_LENGTH
};

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
      case _INT_: return "(INT)";
      case _REAL_: return "(REAL)";
      case _LREAL_: return "(LREAL)";
    }
  } else if (type == RELOP) {
    switch (attr) {
      case NIL: return "(NULL)";
      case _EQ_: return "(EQ)";
      case _NEQ_: return "(NEQ)";
      case _LT_: return "(LT)";
      case _LTE_: return "(LTE)";
      case _GTE_: return "(GTE)";
      case _GT_: return "(GT)";
    }
  } else if (type == MULOP) {
    switch (attr) {
      case NIL: return "(NULL)";
      case _MULT_: return "(MULT)";
      case _DIVISION_: return "(DVSN)";
      case _DIV_: return "(DIV)";
      case _MOD_: return "(MOD)";
      case _AND_: return "(AND)";
    }
  } else if (type == ADDOP) {
    switch (attr) {
      case NIL: return "(NULL)";
      case _PLUS_: return "(PLUS)";
      case _MINUS_: return "(MINUS)";
      case _OR_: return "(OR)";
    }
  } else if (type == LEXERR) {
    switch (attr) {
      case NIL: return "(NULL)";
      case UNREC: return "(?Symbol)";
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
