#ifndef LEXERR_H
#define LEXERR_H
#include <string.h>

Token* check_id_length(LineNode* node, char* lexeme, int trts, int hare) {
  if (hare-trts <= MAX_ID_LENGTH) {
    return NULL;
  }
  throw_error(node->line->number, strcat(lexeme, ": Id cannot be longer than 10 characters.\n"));
  return new_token(node->line->number, lexeme, LEXERR, ID_LENGTH);
}

Token* check_int_length(LineNode* node, char* lexeme, int trts, int hare) {
  if (hare-trts <= MAX_INTEGER_LENGTH) {
    return NULL;
  }
  throw_error(node->line->number, strcat(lexeme, ": Integer cannot be longer than 10 digits.\n"));
  return new_token(node->line->number, lexeme, LEXERR, INT_LENGTH);
}

Token* check_xx_length(LineNode* node, char* lexeme, int trts, int hare) {
  if (hare-trts <= MAX_XX_LENGTH) {
    return NULL;
  }
  throw_error(node->line->number, strcat(lexeme, ": (xx.yyEzz) xx cannot be longer than 5 digits.\n"));
  return new_token(node->line->number, lexeme, LEXERR, XX_LENGTH);
}

Token* check_yy_length(LineNode* node, char* lexeme, int trts, int hare) {
  if (hare-trts <= MAX_YY_LENGTH) {
    return NULL;
  }
  throw_error(node->line->number, strcat(lexeme, ": (xx.yyEzz) yy cannot be longer than 5 digits.\n"));
  return new_token(node->line->number, lexeme, LEXERR, YY_LENGTH);
}

Token* check_zz_length(LineNode* node, char* lexeme, int trts, int hare) {
  if (hare-trts <= MAX_ZZ_LENGTH) {
    return NULL;
  }
  throw_error(node->line->number, strcat(lexeme, ": (xx.yyEzz) zz cannot be longer than 2 digits.\n"));
  return new_token(node->line->number, lexeme, LEXERR, ZZ_LENGTH);
}

#endif
