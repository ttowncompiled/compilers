#ifndef PARSER_H
#define PARSER_H
#include "global.h"

TokenNode* lead_token = NULL;

Token* get_token() {
  if (lead_token == NULL) {
    lead_token = tokens;
    return lead_token->token;
  }
  if (lead_token->next == NULL) {
    return NULL;
  }
  lead_token = lead_token->next;
  return lead_token->token;
}

Token* match(int type) {
  Token* token = get_token();
  if (token->type == type) {
    return token;
  }
  return NULL;
}

void program() {
  Token* token;
  if ((token = match(PROGRAM)) != NULL) {
    // first production
  }
}

void program_body() {
  Token* token;
  if ((token = match(VAR)) != NULL) {
    // first production
  }
  if ((token = match(FUNCTION)) != NULL) {
    // second production
  }
  if ((token = match(BEGIN)) != NULL) {
    // second production
  }
}

void program_subbody() {
  Token* token;
  if ((token = match(FUNCTION)) != NULL) {
    // first production
  }
  if ((token = match(BEGIN)) != NULL) {
    // second production
  }
}

void id() {
  Token* token;
  if ((token = match(ID)) != NULL) {
    // first production
  }
}

void identifier_list() {
  Token* token;
  if ((token = match(ID)) != NULL) {
    // first production
  }
}

void identifier_list_prime() {
  Token* token;
  if ((token = match(COMMA)) != NULL) {
    // first production
  }
  if ((token = match(CLOSEPAREN)) != NULL) {
    // second production
  }
}

void declarations() {
  Token* token;
  if ((token = match(VAR)) != NULL) {
    // first production
  }
}

void declarations_prime() {
  Token* token;
  if ((token = match(VAR)) != NULL) {
    // first production
  }
  if ((token = match(FUNCTION)) != NULL) {
    // second production
  }
  if ((token = match(BEGIN)) != NULL) {
    // second production
  }
}

void type() {
  Token* token;
  if ((token = match(INTEGER)) != NULL) {
    // first production
  }
  if ((token = match(REAL)) != NULL) {
    // first production
  }
  if ((token = match(ARRAY)) != NULL) {
    // second production
  }
}

void standard_type() {
  Token* token;
  if ((token = match(INTEGER)) != NULL) {
    // first production
  }
  if ((token = match(REAL)) != NULL) {
    // second production
  }
}

void subprogram_declarations() {
  Token* token;
  if ((token = match(FUNCTION)) != NULL) {
    // first production
  }
}

void subprogram_declarations_prime() {
  Token* token;
  if ((token = match(FUNCTION)) != NULL) {
    // first production
  }
  if ((token = match(BEGIN)) != NULL) {
    // second production
  }
}

void subprogram_declaration() {
  Token* token;
  if ((token = match(FUNCTION)) != NULL) {
    // first production
  }
}

void subprogram_body() {
  Token* token;
  if ((token = match(VAR)) != NULL) {
    // first production
  }
  if ((token = match(FUNCTION)) != NULL) {
    // second production
  }
  if ((token = match(BEGIN)) != NULL) {
    // second production
  }
}

void subprogram_subbody() {
  Token* token;
  if ((token = match(FUNCTION)) != NULL) {
    // first production
  }
  if ((token = match(BEGIN)) != NULL) {
    // second production
  }
}

void subprogram_head() {
  Token* token;
  if ((token = match(FUNCTION)) != NULL) {
    // first production
  }
}

void subprogram_head_prime() {
  Token* token;
  if ((token = match(OPENPAREN)) != NULL) {
    // first production
  }
  if ((token = match(COLON)) != NULL) {
    // second production
  }
}

void arguments() {
  Token* token;
  if ((token = match(OPENPAREN)) != NULL) {
    // first production
  }
}

void parameter_list() {
  Token* token;
  if ((token = match(ID)) != NULL) {
    // first production
  }
}

void parameter_list_prime() {
  Token* token;
  if ((token = match(SEMICOLON)) != NULL) {
    // first production
  }
  if ((token = match(CLOSEPAREN)) != NULL) {
    // second production
  }
}

void compound_statement() {
  Token* token;
  if ((token = match(BEGIN)) != NULL) {
    // first production
  }
}

void compound_statement_prime() {
  Token* token;
  if ((token = match(ID)) != NULL) {
    // first production
  }
  if ((token = match(BEGIN)) != NULL) {
    // first production
  }
  if ((token = match(IF)) != NULL) {
    // first production
  }
  if ((token = match(WHILE)) != NULL) {
    // first production
  }
  if ((token = match(END)) != NULL) {
    // second production
  }
}

void optional_statements() {
  Token* token;
  if ((token = match(ID)) != NULL) {
    // first production
  }
  if ((token = match(BEGIN)) != NULL) {
    // first production
  }
  if ((token = match(IF)) != NULL) {
    // first production
  }
  if ((token = match(WHILE)) != NULL) {
    // first production
  }
}

void statement_list() {
  Token* token;
  if ((token = match(ID)) != NULL) {
    // first production
  }
  if ((token = match(BEGIN)) != NULL) {
    // first production
  }
  if ((token = match(IF)) != NULL) {
    // first production
  }
  if ((token = match(WHILE)) != NULL) {
    // first production
  }
}

void statement_list_prime() {
  Token* token;
  if ((token = match(SEMICOLON)) != NULL) {
    // first production
  }
  if ((token = match(END)) != NULL) {
    // second production
  }
}

void statement() {
  Token* token;
  if ((token = match(ID)) != NULL) {
    // first production
  }
  if ((token = match(BEGIN)) != NULL) {
    // second production
  }
  if ((token = match(IF)) != NULL) {
    // third production
  }
  if ((token = match(WHILE)) != NULL) {
    // fourth production
  }
}

void statement_prime() {
  Token* token;
  if ((token = match(ELSE)) != NULL) {
    // first production
  }
  if ((token = match(SEMICOLON)) != NULL) {
    // second production
  }
  if ((token = match(END)) != NULL) {
    // second production
  }
}

void variable() {
  Token* token;
  if ((token = match(ID)) != NULL) {
    // first production
  }
}

void variable_prime() {
  Token* token;
  if ((token = match(OPENBRACKET)) != NULL) {
    // first production
  }
  if ((token = match(ASSIGNOP)) != NULL) {
    // second production
  }
}

void expression_list() {
  Token* token;
  if ((token = match(ID)) != NULL) {
    // first production
  }
  if ((token = match(NUM)) != NULL) {
    // first production
  }
  if ((token = match(OPENPAREN)) != NULL) {
    // first production
  }
  if ((token = match(NOT)) != NULL) {
    // first production
  }
  if ((token = match(ADDOP)) != NULL && token->attr.value == _PLUS_) {
    // first production
  }
  if ((token = match(ADDOP)) != NULL && token->attr.value == _MINUS_) {
    // first production
  }
}

void expression_list_prime() {
  Token* token;
  if ((token = match(COMMA)) != NULL) {
    // first production
  }
  if ((token = match(CLOSEPAREN)) != NULL) {
    // second production
  }
}

void expression() {
  Token* token;
  if ((token = match(ID)) != NULL) {
    // first production
  }
  if ((token = match(NUM)) != NULL) {
    // first production
  }
  if ((token = match(OPENPAREN)) != NULL) {
    // first production
  }
  if ((token = match(NOT)) != NULL) {
    // first production
  }
  if ((token = match(ADDOP)) != NULL && token->attr.value == _PLUS_) {
    // first production
  }
  if ((token = match(ADDOP)) != NULL && token->attr.value == _MINUS_) {
    // first production
  }
}

void expression_prime() {
  Token* token;
  if ((token = match(RELOP)) != NULL) {
    // first production
  }
  if ((token = match(ELSE)) != NULL) {
    // second production
  }
  if ((token = match(SEMICOLON)) != NULL) {
    // second production
  }
  if ((token = match(END)) != NULL) {
    // second production
  }
  if ((token = match(THEN)) != NULL) {
    // second production
  }
  if ((token = match(DO)) != NULL) {
    // second production
  }
  if ((token = match(CLOSEBRACKET)) != NULL) {
    // second production
  }
  if ((token = match(CLOSEPAREN)) != NULL) {
    // second production
  }
}

void simple_expression() {
  Token* token;
  if ((token = match(ID)) != NULL) {
    // first production
  }
  if ((token = match(NUM)) != NULL) {
    // first production
  }
  if ((token = match(OPENPAREN)) != NULL) {
    // first production
  }
  if ((token = match(NOT)) != NULL) {
    // first production
  }
  if ((token = match(ADDOP)) != NULL && token->attr.value == _PLUS_) {
    // second production
  }
  if ((token = match(ADDOP)) != NULL && token->attr.value == _MINUS_) {
    // second production
  }
}

void simple_expression_prime() {
  Token* token;
  if ((token = match(ADDOP)) != NULL) {
    // first production
  }
  if ((token = match(RELOP)) != NULL) {
    // second production
  }
  if ((token = match(ELSE)) != NULL) {
    // second production
  }
  if ((token = match(SEMICOLON)) != NULL) {
    // second production
  }
  if ((token = match(END)) != NULL) {
    // second production
  }
  if ((token = match(THEN)) != NULL) {
    // second production
  }
  if ((token = match(DO)) != NULL) {
    // second production
  }
  if ((token = match(CLOSEBRACKET)) != NULL) {
    // second production
  }
  if ((token = match(CLOSEPAREN)) != NULL) {
    // second production
  }
}

void term() {
  Token* token;
  if ((token = match(ID)) != NULL) {
    // first production
  }
  if ((token = match(NUM)) != NULL) {
    // first production
  }
  if ((token = match(OPENPAREN)) != NULL) {
    // first production
  }
  if ((token = match(NOT)) != NULL) {
    // first production
  }
}

void term_prime() {
  Token* token;
  if ((token = match(MULOP)) != NULL) {
    // first production
  }
  if ((token = match(ADDOP)) != NULL) {
    // second production
  }
  if ((token = match(RELOP)) != NULL) {
    // second production
  }
  if ((token = match(ELSE)) != NULL) {
    // second production
  }
  if ((token = match(SEMICOLON)) != NULL) {
    // second production
  }
  if ((token = match(END)) != NULL) {
    // second production
  }
  if ((token = match(THEN)) != NULL) {
    // second production
  }
  if ((token = match(DO)) != NULL) {
    // second production
  }
  if ((token = match(CLOSEBRACKET)) != NULL) {
    // second production
  }
  if ((token = match(CLOSEPAREN)) != NULL) {
    // second production
  }
}

void factor() {
  Token* token;
  if ((token = match(ID)) != NULL) {
    // first production
  }
  if ((token = match(NUM)) != NULL) {
    // second production
  }
  if ((token = match(OPENPAREN)) != NULL) {
    // third production
  }
  if ((token = match(NOT)) != NULL) {
    // fourth production
  }
}

void factor_prime() {
  Token* token;
  if ((token = match(OPENPAREN)) != NULL) {
    // first production
  }
  if ((token = match(OPENBRACKET)) != NULL) {
    // second production
  }
  if ((token = match(MULOP)) != NULL) {
    // third production
  }
  if ((token = match(ADDOP)) != NULL) {
    // third production
  }
  if ((token = match(RELOP)) != NULL) {
    // third production
  }
  if ((token = match(ELSE)) != NULL) {
    // third production
  }
  if ((token = match(SEMICOLON)) != NULL) {
    // third production
  }
  if ((token = match(END)) != NULL) {
    // third production
  }
  if ((token = match(THEN)) != NULL) {
    // third production
  }
  if ((token = match(DO)) != NULL) {
    // third production
  }
  if ((token = match(CLOSEBRACKET)) != NULL) {
    // third production
  }
  if ((token = match(CLOSEPAREN)) != NULL) {
    // third production
  }
}

void sign() {
  Token* token;
  if ((token = match(ADDOP)) != NULL && token->attr.value == _PLUS_) {
    // first production
  }
  if ((token = match(ADDOP)) != NULL && token->attr.value == _MINUS_) {
    // second production
  }
}

#endif
