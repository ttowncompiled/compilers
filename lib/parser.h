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

void program();
void program_body();
void program_subbody();
void id();
void identifier_list();
void identifier_list_prime();
void declarations();
void declarations_prime();
void type();
void standard_type();
void subprogram_declarations();
void subprogram_declarations_prime();
void subprogram_declaration();
void subprogram_body();
void subprogram_subbody();
void subprogram_head();
void subprogram_head_prime();
void arguments();
void parameter_list();
void parameter_list_prime();
void compound_statement();
void compound_statement_prime();
void optional_statements();
void statement_list();
void statement_list_prime();
void statement();
void statement_prime();
void variable();
void variable_prime();
void expression_list();
void expression_list_prime();
void expression();
void expression_prime();
void simple_expression();
void simple_expression_prime();
void term();
void term_prime();
void factor();
void factor_prime();
void sign();

void program() {
  Token* token;
  if ((token = match(PROGRAM)) != NULL) {
    // first production
    if ((token = match(ID)) == NULL) {
      
    }
    if ((token = match(OPENPAREN)) == NULL) {
      
    }
    identifier_list();
    if ((token = match(CLOSEPAREN)) == NULL) {
      
    }
    if ((token = match(SEMICOLON)) == NULL) {
      
    }
    program_body();
    return;
  }
}

void program_body() {
  Token* token;
  if ((token = match(VAR)) != NULL) {
    // first production
    declarations();
    program_subbody();
    return;
  }
  if ((token = match(FUNCTION)) != NULL) {
    // second production
    program_subbody();
    return;
  }
  if ((token = match(BEGIN)) != NULL) {
    // second production
    program_subbody();
    return;
  }
}

void program_subbody() {
  Token* token;
  if ((token = match(FUNCTION)) != NULL) {
    // first production
    subprogram_declarations();
    compound_statement();
    if ((token = match(PERIOD)) == NULL) {
      
    }
    return;
  }
  if ((token = match(BEGIN)) != NULL) {
    // second production
    compound_statement();
    if ((token = match(PERIOD)) == NULL) {
      
    }
    return;
  }
}

void id() {
  Token* token;
  if ((token = match(ID)) != NULL) {
    // first production
    return;
  }
}

void identifier_list() {
  Token* token;
  if ((token = match(ID)) != NULL) {
    // first production
    identifier_list_prime();
    return;
  }
}

void identifier_list_prime() {
  Token* token;
  if ((token = match(COMMA)) != NULL) {
    // first production
    if ((token = match(ID)) == NULL) {
      
    }
    identifier_list_prime();
    return;
  }
  if ((token = match(CLOSEPAREN)) != NULL) {
    // second production
    return;
  }
}

void declarations() {
  Token* token;
  if ((token = match(VAR)) != NULL) {
    // first production
    id();
    if ((token = match(COLON)) == NULL) {
      
    }
    type();
    if ((token = match(SEMICOLON)) == NULL) {
      
    }
    declarations_prime();
    return;
  }
}

void declarations_prime() {
  Token* token;
  if ((token = match(VAR)) != NULL) {
    // first production
    id();
    if ((token = match(COLON)) == NULL) {
      
    }
    type();
    if ((token = match(SEMICOLON)) == NULL) {
      
    }
    declarations_prime();
    return;
  }
  if ((token = match(FUNCTION)) != NULL) {
    // second production
    return;
  }
  if ((token = match(BEGIN)) != NULL) {
    // second production
    return;
  }
}

void type() {
  Token* token;
  if ((token = match(INTEGER)) != NULL) {
    // first production
    standard_type();
    return;
  }
  if ((token = match(REAL)) != NULL) {
    // first production
    standard_type();
    return;
  }
  if ((token = match(ARRAY)) != NULL) {
    // second production
    if ((token = match(OPENBRACKET)) == NULL) {
      
    }
    if ((token = match(NUM)) == NULL) {
      
    }
    if ((token = match(RANGE)) == NULL) {
      
    }
    if ((token = match(NUM)) == NULL) {
      
    }
    if ((token = match(CLOSEBRACKET)) == NULL) {
      
    }
    if ((token = match(OF)) == NULL) {
      
    }
    standard_type();
    return;
  }
}

void standard_type() {
  Token* token;
  if ((token = match(INTEGER)) != NULL) {
    // first production
    return;
  }
  if ((token = match(REAL)) != NULL) {
    // second production
    return;
  }
}

void subprogram_declarations() {
  Token* token;
  if ((token = match(FUNCTION)) != NULL) {
    // first production
    subprogram_declaration();
    if ((token = match(SEMICOLON)) == NULL) {
      
    }
    subprogram_declarations_prime();
    return;
  }
}

void subprogram_declarations_prime() {
  Token* token;
  if ((token = match(FUNCTION)) != NULL) {
    // first production
    subprogram_declaration();
    if ((token = match(SEMICOLON)) == NULL) {
      
    }
    subprogram_declarations_prime();
    return;
  }
  if ((token = match(BEGIN)) != NULL) {
    // second production
    return;
  }
}

void subprogram_declaration() {
  Token* token;
  if ((token = match(FUNCTION)) != NULL) {
    // first production
    subprogram_head();
    subprogram_body();
    return;
  }
}

void subprogram_body() {
  Token* token;
  if ((token = match(VAR)) != NULL) {
    // first production
    declarations();
    subprogram_subbody();
    return;
  }
  if ((token = match(FUNCTION)) != NULL) {
    // second production
    subprogram_subbody();
    return;
  }
  if ((token = match(BEGIN)) != NULL) {
    // second production
    subprogram_subbody();
    return;
  }
}

void subprogram_subbody() {
  Token* token;
  if ((token = match(FUNCTION)) != NULL) {
    // first production
    subprogram_declarations();
    compound_statement();
    return;
  }
  if ((token = match(BEGIN)) != NULL) {
    // second production
    compound_statement();
    return;
  }
}

void subprogram_head() {
  Token* token;
  if ((token = match(FUNCTION)) != NULL) {
    // first production
    if ((token = match(ID)) == NULL) {
      
    }
    subprogram_head_prime();
    return;
  }
}

void subprogram_head_prime() {
  Token* token;
  if ((token = match(OPENPAREN)) != NULL) {
    // first production
    arguments();
    if ((token = match(COLON)) == NULL) {
      
    }
    standard_type();
    if ((token = match(SEMICOLON)) == NULL) {
      
    }
    return;
  }
  if ((token = match(COLON)) != NULL) {
    // second production
    if ((token = match(COLON)) == NULL) {
      
    }
    standard_type();
    if ((token = match(SEMICOLON)) == NULL) {
      
    }
    return;
  }
}

void arguments() {
  Token* token;
  if ((token = match(OPENPAREN)) != NULL) {
    // first production
    parameter_list();
    if ((token = match(CLOSEPAREN)) == NULL) {
      
    }
    return;
  }
}

void parameter_list() {
  Token* token;
  if ((token = match(ID)) != NULL) {
    // first production
    id();
    if ((token = match(COLON)) == NULL) {
      
    }
    type();
    parameter_list_prime();
    return;
  }
}

void parameter_list_prime() {
  Token* token;
  if ((token = match(SEMICOLON)) != NULL) {
    // first production
    id();
    if ((token = match(COLON)) == NULL) {
      
    }
    type();
    parameter_list_prime();
    return;
  }
  if ((token = match(CLOSEPAREN)) != NULL) {
    // second production
    return;
  }
}

void compound_statement() {
  Token* token;
  if ((token = match(BEGIN)) != NULL) {
    // first production
    compound_statement_prime();
    return;
  }
}

void compound_statement_prime() {
  Token* token;
  if ((token = match(ID)) != NULL) {
    // first production
    optional_statements();
    if ((token = match(END)) == NULL) {
      
    }
    return;
  }
  if ((token = match(BEGIN)) != NULL) {
    // first production
    optional_statements();
    if ((token = match(END)) == NULL) {
      
    }
    return;
  }
  if ((token = match(IF)) != NULL) {
    // first production
    optional_statements();
    if ((token = match(END)) == NULL) {
      
    }
    return;
  }
  if ((token = match(WHILE)) != NULL) {
    // first production
    optional_statements();
    if ((token = match(END)) == NULL) {
      
    }
    return;
  }
  if ((token = match(END)) != NULL) {
    // second production
    return;
  }
}

void optional_statements() {
  Token* token;
  if ((token = match(ID)) != NULL) {
    // first production
    statement_list();
    return;
  }
  if ((token = match(BEGIN)) != NULL) {
    // first production
    statement_list();
    return;
  }
  if ((token = match(IF)) != NULL) {
    // first production
    statement_list();
    return;
  }
  if ((token = match(WHILE)) != NULL) {
    // first production
    statement_list();
    return;
  }
}

void statement_list() {
  Token* token;
  if ((token = match(ID)) != NULL) {
    // first production
    statement();
    statement_list_prime();
    return;
  }
  if ((token = match(BEGIN)) != NULL) {
    // first production
    statement();
    statement_list_prime();
    return;
  }
  if ((token = match(IF)) != NULL) {
    // first production
    statement();
    statement_list_prime();
    return;
  }
  if ((token = match(WHILE)) != NULL) {
    // first production
    statement();
    statement_list_prime();
    return;
  }
}

void statement_list_prime() {
  Token* token;
  if ((token = match(SEMICOLON)) != NULL) {
    // first production
    statement();
    statement_list_prime();
    return;
  }
  if ((token = match(END)) != NULL) {
    // second production
    return;
  }
}

void statement() {
  Token* token;
  if ((token = match(ID)) != NULL) {
    // first production
    variable();
    if ((token = match(ASSIGNOP)) == NULL) {
      
    }
    expression();
    return;
  }
  if ((token = match(BEGIN)) != NULL) {
    // second production
    compound_statement();
    return;
  }
  if ((token = match(IF)) != NULL) {
    // third production
    expression();
    if ((token = match(THEN)) == NULL) {
      
    }
    statement();
    statement_prime();
    return;
  }
  if ((token = match(WHILE)) != NULL) {
    // fourth production
    expression();
    if ((token = match(DO)) == NULL) {
      
    }
    statement();
    return;
  }
}

void statement_prime() {
  Token* token;
  if ((token = match(ELSE)) != NULL) {
    // first production
    statement();
    return;
  }
  if ((token = match(SEMICOLON)) != NULL) {
    // second production
    return;
  }
  if ((token = match(END)) != NULL) {
    // second production
    return;
  }
}

void variable() {
  Token* token;
  if ((token = match(ID)) != NULL) {
    // first production
    variable_prime();
    return;
  }
}

void variable_prime() {
  Token* token;
  if ((token = match(OPENBRACKET)) != NULL) {
    // first production
    expression();
    if ((token = match(CLOSEBRACKET)) == NULL) {
      
    }
    return;
  }
  if ((token = match(ASSIGNOP)) != NULL) {
    // second production
    return;
  }
}

void expression_list() {
  Token* token;
  if ((token = match(ID)) != NULL) {
    // first production
    expression();
    expression_list_prime();
    return;
  }
  if ((token = match(NUM)) != NULL) {
    // first production
    expression();
    expression_list_prime();
    return;
  }
  if ((token = match(OPENPAREN)) != NULL) {
    // first production
    expression();
    expression_list_prime();
    return;
  }
  if ((token = match(NOT)) != NULL) {
    // first production
    expression();
    expression_list_prime();
    return;
  }
  if ((token = match(ADDOP)) != NULL && token->attr.value == _PLUS_) {
    // first production
    expression();
    expression_list_prime();
    return;
  }
  if ((token = match(ADDOP)) != NULL && token->attr.value == _MINUS_) {
    // first production
    expression();
    expression_list_prime();
    return;
  }
}

void expression_list_prime() {
  Token* token;
  if ((token = match(COMMA)) != NULL) {
    // first production
    expression();
    expression_list_prime();
    return;
  }
  if ((token = match(CLOSEPAREN)) != NULL) {
    // second production
    return;
  }
}

void expression() {
  Token* token;
  if ((token = match(ID)) != NULL) {
    // first production
    simple_expression();
    expression_prime();
    return;
  }
  if ((token = match(NUM)) != NULL) {
    // first production
    simple_expression();
    expression_prime();
    return;
  }
  if ((token = match(OPENPAREN)) != NULL) {
    // first production
    simple_expression();
    expression_prime();
    return;
  }
  if ((token = match(NOT)) != NULL) {
    // first production
    simple_expression();
    expression_prime();
    return;
  }
  if ((token = match(ADDOP)) != NULL && token->attr.value == _PLUS_) {
    // first production
    simple_expression();
    expression_prime();
    return;
  }
  if ((token = match(ADDOP)) != NULL && token->attr.value == _MINUS_) {
    // first production
    simple_expression();
    expression_prime();
    return;
  }
}

void expression_prime() {
  Token* token;
  if ((token = match(RELOP)) != NULL) {
    // first production
    simple_expression();
    return;
  }
  if ((token = match(ELSE)) != NULL) {
    // second production
    return;
  }
  if ((token = match(SEMICOLON)) != NULL) {
    // second production
    return;
  }
  if ((token = match(END)) != NULL) {
    // second production
    return;
  }
  if ((token = match(THEN)) != NULL) {
    // second production
    return;
  }
  if ((token = match(DO)) != NULL) {
    // second production
    return;
  }
  if ((token = match(CLOSEBRACKET)) != NULL) {
    // second production
    return;
  }
  if ((token = match(CLOSEPAREN)) != NULL) {
    // second production
    return;
  }
}

void simple_expression() {
  Token* token;
  if ((token = match(ID)) != NULL) {
    // first production
    term();
    simple_expression_prime();
    return;
  }
  if ((token = match(NUM)) != NULL) {
    // first production
    term();
    simple_expression_prime();
    return;
  }
  if ((token = match(OPENPAREN)) != NULL) {
    // first production
    term();
    simple_expression_prime();
    return;
  }
  if ((token = match(NOT)) != NULL) {
    // first production
    term();
    simple_expression_prime();
    return;
  }
  if ((token = match(ADDOP)) != NULL && token->attr.value == _PLUS_) {
    // second production
    sign();
    term();
    simple_expression_prime();
    return;
  }
  if ((token = match(ADDOP)) != NULL && token->attr.value == _MINUS_) {
    // second production
    sign();
    term();
    simple_expression_prime();
    return;
  }
}

void simple_expression_prime() {
  Token* token;
  if ((token = match(ADDOP)) != NULL) {
    // first production
    term();
    simple_expression_prime();
    return;
  }
  if ((token = match(RELOP)) != NULL) {
    // second production
    return;
  }
  if ((token = match(ELSE)) != NULL) {
    // second production
    return;
  }
  if ((token = match(SEMICOLON)) != NULL) {
    // second production
    return;
  }
  if ((token = match(END)) != NULL) {
    // second production
    return;
  }
  if ((token = match(THEN)) != NULL) {
    // second production
    return;
  }
  if ((token = match(DO)) != NULL) {
    // second production
    return;
  }
  if ((token = match(CLOSEBRACKET)) != NULL) {
    // second production
    return;
  }
  if ((token = match(CLOSEPAREN)) != NULL) {
    // second production
    return;
  }
}

void term() {
  Token* token;
  if ((token = match(ID)) != NULL) {
    // first production
    factor();
    term_prime();
    return;
  }
  if ((token = match(NUM)) != NULL) {
    // first production
    factor();
    term_prime();
    return;
  }
  if ((token = match(OPENPAREN)) != NULL) {
    // first production
    factor();
    term_prime();
    return;
  }
  if ((token = match(NOT)) != NULL) {
    // first production
    factor();
    term_prime();
    return;
  }
}

void term_prime() {
  Token* token;
  if ((token = match(MULOP)) != NULL) {
    // first production
    factor();
    term_prime();
    return;
  }
  if ((token = match(ADDOP)) != NULL) {
    // second production
    return;
  }
  if ((token = match(RELOP)) != NULL) {
    // second production
    return;
  }
  if ((token = match(ELSE)) != NULL) {
    // second production
    return;
  }
  if ((token = match(SEMICOLON)) != NULL) {
    // second production
    return;
  }
  if ((token = match(END)) != NULL) {
    // second production
    return;
  }
  if ((token = match(THEN)) != NULL) {
    // second production
    return;
  }
  if ((token = match(DO)) != NULL) {
    // second production
    return;
  }
  if ((token = match(CLOSEBRACKET)) != NULL) {
    // second production
    return;
  }
  if ((token = match(CLOSEPAREN)) != NULL) {
    // second production
    return;
  }
}

void factor() {
  Token* token;
  if ((token = match(ID)) != NULL) {
    // first production
    factor_prime();
    return;
  }
  if ((token = match(NUM)) != NULL) {
    // second production
    return;
  }
  if ((token = match(OPENPAREN)) != NULL) {
    // third production
    expression();
    if ((token = match(CLOSEPAREN)) == NULL) {
      
    }
    return;
  }
  if ((token = match(NOT)) != NULL) {
    // fourth production
    factor();
    return;
  }
}

void factor_prime() {
  Token* token;
  if ((token = match(OPENPAREN)) != NULL) {
    // first production
    expression_list();
    if ((token = match(CLOSEPAREN)) == NULL) {
      
    }
    return;
  }
  if ((token = match(OPENBRACKET)) != NULL) {
    // second production
    expression();
    if ((token = match(CLOSEBRACKET)) == NULL) {
      
    }
    return;
  }
  if ((token = match(MULOP)) != NULL) {
    // third production
    return;
  }
  if ((token = match(ADDOP)) != NULL) {
    // third production
    return;
  }
  if ((token = match(RELOP)) != NULL) {
    // third production
    return;
  }
  if ((token = match(ELSE)) != NULL) {
    // third production
    return;
  }
  if ((token = match(SEMICOLON)) != NULL) {
    // third production
    return;
  }
  if ((token = match(END)) != NULL) {
    // third production
    return;
  }
  if ((token = match(THEN)) != NULL) {
    // third production
    return;
  }
  if ((token = match(DO)) != NULL) {
    // third production
    return;
  }
  if ((token = match(CLOSEBRACKET)) != NULL) {
    // third production
    return;
  }
  if ((token = match(CLOSEPAREN)) != NULL) {
    // third production
    return;
  }
}

void sign() {
  Token* token;
  if ((token = match(ADDOP)) != NULL && token->attr.value == _PLUS_) {
    // first production
    return;
  }
  if ((token = match(ADDOP)) != NULL && token->attr.value == _MINUS_) {
    // second production
    return;
  }
}

#endif
