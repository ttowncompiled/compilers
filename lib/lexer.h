#ifndef LEXER_H 
#define LEXER_H
#include "enums.h"
#include "types.h"
#include "io.h"
#include "util.h"

const int MAX_INTEGER_LENGTH = 10;
const int MAX_XX_LENGTH = 5;
const int MAX_YY_LENGTH = 5;
const int MAX_ZZ_LENGTH = 2;

Token* check_id_length(LineNode* node, char* lexeme, int trts, int hare) {
  if (hare-trts <= MAX_ID_LENGTH) {
    return NULL;
  }
  LineNode* error = line_node_of(
      concat(lexeme, ": Id cannot be longer than 10 characters.\n"),
      node->line->number);
  LineNode* curr = node;
  while (curr->error != NULL) {
    curr = curr->error;
  }
  curr->error = error;
  return token_of(node->line->number, lexeme, LEXERR, ID_LENGTH);
}

Token* check_int_length(LineNode* node, char* lexeme, int trts, int hare) {
  if (hare-trts <= MAX_INTEGER_LENGTH) {
    return NULL;
  }
  LineNode* error = line_node_of(
      concat(lexeme, ": Integer cannot be longer than 10 digits.\n"),
      node->line->number);
  LineNode* curr = node;
  while (curr->error != NULL) {
    curr = curr->error;
  }
  curr->error = error;
  return token_of(node->line->number, lexeme, LEXERR, INT_LENGTH);
}

Token* check_xx_length(LineNode* node, char* lexeme, int trts, int hare) {
  if (hare-trts <= MAX_XX_LENGTH) {
    return NULL;
  }
  LineNode* error = line_node_of(
      concat(lexeme, ": (xx.yyEzz) xx cannot be longer than 5 digits.\n"),
      node->line->number);
  LineNode* curr = node;
  while (curr->error != NULL) {
    curr = curr->error;
  }
  curr->error = error;
  return token_of(node->line->number, lexeme, LEXERR, XX_LENGTH);
}

Token* check_yy_length(LineNode* node, char* lexeme, int trts, int hare) {
  if (hare-trts <= MAX_YY_LENGTH) {
    return NULL;
  }
  LineNode* error = line_node_of(
      concat(lexeme, ": (xx.yyEzz) yy cannot be longer than 5 digits.\n"),
      node->line->number);
  LineNode* curr = node;
  while (curr->error != NULL) {
    curr = curr->error;
  }
  curr->error = error;
  return token_of(node->line->number, lexeme, LEXERR, YY_LENGTH);
}

Token* check_zz_length(LineNode* node, char* lexeme, int trts, int hare) {
  if (hare-trts <= MAX_ZZ_LENGTH) {
    return NULL;
  }
  LineNode* error = line_node_of(
      concat(lexeme, ": (xx.yyEzz) zz cannot be longer than 2 digits.\n"),
      node->line->number);
  LineNode* curr = node;
  while (curr->error != NULL) {
    curr = curr->error;
  }
  curr->error = error;
  return token_of(node->line->number, lexeme, LEXERR, ZZ_LENGTH);
}

ReservedWordNode* load_reserved_words();

LineNode* organize(char* filename);

int print_token_file(TokenNode* head);

int print_listing_file(LineNode* head);

TokenNode* analyze(LineNode* first, ReservedWordNode* reserved);

int white_space_machine(LineNode* node, int* trts);

Token* id_machine(LineNode* node, ReservedWordNode* reserved,
    SymbolNode* symbols, int* trts);
    
Token* long_real_machine(LineNode* node, int* trts);

Token* real_machine(LineNode* node, int* trts);

Token* int_machine(LineNode* node, int* trts);

Token* relop_machine(LineNode* node, int* trts);

Token* addop_machine(LineNode* node, int* trts);

Token* mulop_machine(LineNode* node, int* trts);

Token* assignop_machine(LineNode* node, int* trts);

Token* catchall_machine(LineNode* node, int* trts);

#endif
