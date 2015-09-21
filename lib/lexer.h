#ifndef LEXER_H 
#define LEXER_H
#include "global.h"
#include "enums.h"
#include "types.h"
#include "io.h"
#include "util.h"

Token* check_id_length(LineNode* node, char* lexeme, int trts, int hare) {
  if (hare-trts <= MAX_ID_LENGTH) {
    return NULL;
  }
  throw_error(node->line->number, concat(lexeme, ": Id cannot be longer than 10 characters.\n"));
  return new_token(node->line->number, lexeme, LEXERR, ID_LENGTH);
}

Token* check_int_length(LineNode* node, char* lexeme, int trts, int hare) {
  if (hare-trts <= MAX_INTEGER_LENGTH) {
    return NULL;
  }
  throw_error(node->line->number, concat(lexeme, ": Integer cannot be longer than 10 digits.\n"));
  return new_token(node->line->number, lexeme, LEXERR, INT_LENGTH);
}

Token* check_xx_length(LineNode* node, char* lexeme, int trts, int hare) {
  if (hare-trts <= MAX_XX_LENGTH) {
    return NULL;
  }
  throw_error(node->line->number, concat(lexeme, ": (xx.yyEzz) xx cannot be longer than 5 digits.\n"));
  return new_token(node->line->number, lexeme, LEXERR, XX_LENGTH);
}

Token* check_yy_length(LineNode* node, char* lexeme, int trts, int hare) {
  if (hare-trts <= MAX_YY_LENGTH) {
    return NULL;
  }
  throw_error(node->line->number, concat(lexeme, ": (xx.yyEzz) yy cannot be longer than 5 digits.\n"));
  return new_token(node->line->number, lexeme, LEXERR, YY_LENGTH);
}

Token* check_zz_length(LineNode* node, char* lexeme, int trts, int hare) {
  if (hare-trts <= MAX_ZZ_LENGTH) {
    return NULL;
  }
  throw_error(node->line->number, concat(lexeme, ": (xx.yyEzz) zz cannot be longer than 2 digits.\n"));
  return new_token(node->line->number, lexeme, LEXERR, ZZ_LENGTH);
}

TokenNode* tokenize(LineNode* lines);

int white_space_machine(LineNode* node, int* trts);

Token* id_machine(LineNode* node, int* trts);
    
Token* long_real_machine(LineNode* node, int* trts);

Token* real_machine(LineNode* node, int* trts);

Token* int_machine(LineNode* node, int* trts);

Token* relop_machine(LineNode* node, int* trts);

Token* addop_machine(LineNode* node, int* trts);

Token* mulop_machine(LineNode* node, int* trts);

Token* assignop_machine(LineNode* node, int* trts);

Token* catchall_machine(LineNode* node, int* trts);

#endif
