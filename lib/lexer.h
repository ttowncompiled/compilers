#ifndef LEXER_H 
#define LEXER_H
#include "global.h"
#include "lexerr.h"

TokenNode* tokens = NULL;

int white_space_machine(LineNode* node, int* trts) {
  char* buffer = node->line->value;
  int hare = (*trts);
  while (isspace(buffer[hare])) {
    hare++;
  }
  (*trts) = hare;
  return buffer[hare] == '\0';
}

Token* id_machine(LineNode* node, int* trts) {
  char* buffer = node->line->value;
  int hare = (*trts);
  if (isalpha(buffer[hare])) {
    Token* token;
    hare++;
    while (isalpha(buffer[hare]) || isdigit(buffer[hare])) {
      hare++;
    }
    char* lexeme = substring(buffer, (*trts), hare);
    if ((token = check_id_length(node, lexeme, (*trts), hare)) != NULL) {
      (*trts) = hare;
      return token;
    }
    (*trts) = hare;
    ReservedWordNode* reserved = reserved_word_table;
    while (reserved != NULL) {
      if (strcmp(lexeme, reserved->word->value) == 0) {
        return new_token(node->line->number,
                        lexeme,
                        reserved->word->type,
                        reserved->word->attr);
      }
      reserved = reserved->next;
    }
    token = new_token(node->line->number, lexeme, ID, NIL);
    Attribute attribute;
    attribute.address = save_symbol(lexeme);
    token->attr = attribute;
    return token;
  }
  return NULL;
}

Token* long_real_machine(LineNode* node, int* trts) {
  char* buffer = node->line->value;
  int hare = (*trts);
  if (isdigit(buffer[hare])) {
    Token* token;
    while(isdigit(buffer[hare])) {
      hare++;
    }
    token = check_xx_length(node,
                            substring(buffer, (*trts), hare),
                            (*trts),
                            hare);
    if (buffer[hare] != '.' || !isdigit(buffer[hare+1])) {
      return NULL;
    }
    hare++;
    while(isdigit(buffer[hare])) {
      hare++;
    }
    token = check_yy_length(node,
                            substring(buffer, (*trts), hare),
                            (*trts),
                            hare);
    if (buffer[hare] != 'E' || !isdigit(buffer[hare+1])) {
      return NULL;
    }
    hare++;
    while(isdigit(buffer[hare])) {
      hare++;
    }
    token = check_zz_length(node,
                            substring(buffer, (*trts), hare),
                            (*trts),
                            hare);
    char* lexeme = substring(buffer, (*trts), hare);
    (*trts) = hare;
    if (token != NULL) {
      token->lexeme = lexeme;
      return token;
    }
    return new_token(node->line->number, lexeme, NUM, _LREAL_);
  }
  return NULL;
}

Token* real_machine(LineNode* node, int* trts) {
  char* buffer = node->line->value;
  int hare = (*trts);
  if (isdigit(buffer[hare])) {
    Token* token;
    while(isdigit(buffer[hare])) {
      hare++;
    }
    token = check_xx_length(node,
                            substring(buffer, (*trts), hare),
                            (*trts),
                            hare);
    if (buffer[hare] != '.' || !isdigit(buffer[hare+1])) {
      return NULL;
    }
    hare++;
    while(isdigit(buffer[hare])) {
      hare++;
    }
    token = check_yy_length(node,
                            substring(buffer, (*trts), hare),
                            (*trts),
                            hare);
    char* lexeme = substring(buffer, (*trts), hare);
    (*trts) = hare;
    if (token != NULL) {
      token->lexeme = lexeme;
      return token;
    }
    return new_token(node->line->number, lexeme, NUM, _REAL_);
  }
  return NULL;
}

Token* int_machine(LineNode* node, int* trts) {
  char* buffer = node->line->value;
  int hare = (*trts);
  if (isdigit(buffer[hare])) {
    Token* token;
    while(isdigit(buffer[hare])) {
      hare++;
    }
    char* lexeme = substring(buffer, (*trts), hare);
    if ((token = check_int_length(node, lexeme, (*trts), hare)) != NULL) {
      (*trts) = hare;
      return token;
    }
    (*trts) = hare;
    return new_token(node->line->number, lexeme, NUM, _INT_);
  }
  return NULL;
}

Token* relop_machine(LineNode* node, int* trts) {
  char* buffer = node->line->value;
  char* lexeme;
  switch(buffer[(*trts)]) {
    case '=':
      lexeme = substring(buffer, (*trts), (*trts)+1);
      (*trts)++;
      return new_token(node->line->number, lexeme, RELOP, _EQ_);
    case '<':
      if (buffer[(*trts)+1] == '>') {
        lexeme = substring(buffer, (*trts), (*trts)+2);
        (*trts) = (*trts) + 2;
        return new_token(node->line->number, lexeme, RELOP, _NEQ_);
      }
      if (buffer[(*trts)+1] == '=') {
        lexeme = substring(buffer, (*trts), (*trts)+2);
        (*trts) = (*trts) + 2;
        return new_token(node->line->number, lexeme, RELOP, _LTE_);
      }
      lexeme = substring(buffer, (*trts), (*trts)+1);
      (*trts)++;
      return new_token(node->line->number, lexeme, RELOP, _LT_);
    case '>':
      if (buffer[(*trts)+1] == '=') {
        lexeme = substring(buffer, (*trts), (*trts)+2);
        (*trts) = (*trts) + 2;
        return new_token(node->line->number, lexeme, RELOP, _GTE_);
      }
      lexeme = substring(buffer, (*trts), (*trts)+1);
      (*trts)++;
      return new_token(node->line->number, lexeme, RELOP, _GT_);
  }
  return NULL;
}

Token* mulop_machine(LineNode* node, int* trts) {
  char* buffer = node->line->value;
  char* lexeme;
  switch (buffer[(*trts)]) {
    case '*':
      lexeme = substring(buffer, (*trts), (*trts)+1);
      (*trts)++;
      return new_token(node->line->number, lexeme, MULOP, NIL);
    case '/':
      lexeme = substring(buffer, (*trts), (*trts)+1);
      (*trts)++;
      return new_token(node->line->number, lexeme, MULOP, NIL);
  }
  return NULL;
}

Token* addop_machine(LineNode* node, int* trts) {
  char* buffer = node->line->value;
  char* lexeme;
  switch (buffer[(*trts)]) {
    case '+':
      lexeme = substring(buffer, (*trts), (*trts)+1);
      (*trts)++;
      return new_token(node->line->number, lexeme, ASSIGNOP, NIL);
    case '-':
      lexeme = substring(buffer, (*trts), (*trts)+1);
      (*trts)++;
      return new_token(node->line->number, lexeme, ASSIGNOP, NIL);
  }
  return NULL;
}

Token* assignop_machine(LineNode* node, int* trts) {
  char* buffer = node->line->value;
  if (buffer[(*trts)] == ':' && buffer[(*trts)+1] == '=') {
    char* lexeme = substring(buffer, (*trts), (*trts)+2);
    (*trts) = (*trts) + 2;
    return new_token(node->line->number, lexeme, ASSIGNOP, NIL);
  }
  return NULL;
}

Token* catchall_machine(LineNode* node, int* trts) {
  char* buffer = node->line->value;
  if (buffer[(*trts)] == '.' && buffer[(*trts)+1] == '.') {
    char* lexeme = substring(buffer, (*trts), (*trts)+2);
    (*trts)++;
    return new_token(node->line->number, lexeme, RANGE, NIL);
  }
  char* lexeme;
  switch (buffer[(*trts)]) {
    case '(':
      lexeme = substring(buffer, (*trts), (*trts)+1);
      (*trts)++;
      return new_token(node->line->number, lexeme, OPENPAREN, NIL);
    case ')':
      lexeme = substring(buffer, (*trts), (*trts)+1);
      (*trts)++;
      return new_token(node->line->number, lexeme, CLOSEPAREN, NIL);
    case '[':
      lexeme = substring(buffer, (*trts), (*trts)+1);
      (*trts)++;
      return new_token(node->line->number, lexeme, OPENBRACKET, NIL);
    case ']':
      lexeme = substring(buffer, (*trts), (*trts)+1);
      (*trts)++;
      return new_token(node->line->number, lexeme, CLOSEBRACKET, NIL);
    case ':':
      lexeme = substring(buffer, (*trts), (*trts)+1);
      (*trts)++;
      return new_token(node->line->number, lexeme, COLON, NIL);
    case ';':
      lexeme = substring(buffer, (*trts), (*trts)+1);
      (*trts)++;
      return new_token(node->line->number, lexeme, SEMICOLON, NIL);
    case ',':
      lexeme = substring(buffer, (*trts), (*trts)+1);
      (*trts)++;
      return new_token(node->line->number, lexeme, COMMA, NIL);
    case '.':
      lexeme = substring(buffer, (*trts), (*trts)+1);
      (*trts)++;
      return new_token(node->line->number, lexeme, PERIOD, NIL);
  }
  return NULL;
}

void tokenize(LineNode* lines) {
  TokenNode* head = NULL;
  TokenNode* prev;
  int line_count = 0;
  int* trts = malloc(sizeof(int));
  while (lines != NULL) {
    (*trts) = 0;
    while (lines->line->value[(*trts)] != '\0') {
      Token* token;
      if (white_space_machine(lines, trts)) {
        continue;
      }
      if ((token = id_machine(lines, trts)) != NULL) {
      } else if ((token = long_real_machine(lines, trts)) != NULL) {
      } else if ((token = real_machine(lines, trts)) != NULL) {
      } else if ((token = int_machine(lines, trts)) != NULL) {
      } else if ((token = relop_machine(lines, trts)) != NULL) {
      } else if ((token = mulop_machine(lines, trts)) != NULL) {
      } else if ((token = addop_machine(lines, trts)) != NULL) {
      } else if ((token = assignop_machine(lines, trts)) != NULL) {
      } else if ((token = catchall_machine(lines, trts)) != NULL) {
      } else {
        throw_error(lines->line->number, substring(lines->line->value, (*trts), (*trts)+1));
        token = new_token(lines->line->number,
                         substring(lines->line->value, (*trts), (*trts)+1),
                         LEXERR,
                         UNREC);
        (*trts)++;
      }
      TokenNode* curr = new_token_node(token);
      if (head == NULL) {
        head = curr;
        prev = head;
      } else {
        prev->next = curr;
        prev = curr;
      }
    }
    lines = lines->next;
    line_count++;
  }
  prev->next = new_token_node(
                         new_token(++line_count,
                                  "",
                                  ENDFILE,
                                  NIL)
                         );
  tokens = head;
}

#endif
