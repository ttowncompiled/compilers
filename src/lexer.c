#include <stdio.h>
#include <stdlib.h>
#include "../lib/lexer.h"

int main(int argc, char* argv[]) {
  if (argc > 2) {
    printf("This compiler only accepts one source file.\n");
    exit(1);
  }
  if (argc < 2) {
    printf("Please provide the name of a source file.\n");
    return 0;
  }
  ReservedWordNode* reserved = load_reserved_words();
  char* filename = argv[1];
  LineNode* lines = organize(filename);
  TokenNode* tokens = analyze(lines, reserved);
  return print_token_file(tokens) || print_listing_file(lines);
}

TokenNode* analyze(LineNode* lines, ReservedWordNode* reserved) {
  SymbolNode* symbols = new_symbol_table();
  TokenNode* head = malloc(sizeof(TokenNode));
  TokenNode* curr = head;
  int line_count = 0;
  int* trts = malloc(sizeof(int));
  while (lines != NULL) {
    (*trts) = 0;
    while (lines->line->value[(*trts)] != '\0') {
      Token* token;
      if (white_space_machine(lines, trts)) {
        continue;
      }
      if ((token = id_machine(lines, reserved, symbols, trts)) != NULL) {
      } else if ((token = long_real_machine(lines, trts)) != NULL) {
      } else if ((token = real_machine(lines, trts)) != NULL) {
      } else if ((token = int_machine(lines, trts)) != NULL) {
      } else if ((token = relop_machine(lines, trts)) != NULL) {
      } else if ((token = mulop_machine(lines, trts)) != NULL) {
      } else if ((token = addop_machine(lines, trts)) != NULL) {
      } else if ((token = assignop_machine(lines, trts)) != NULL) {
      } else if ((token = catchall_machine(lines, trts)) != NULL) {
      } else {
        token = token_of(lines->line->number,
                         substring(lines->line->value, (*trts), (*trts)+1),
                         LEXERR,
                         UNRECOG);
        (*trts)++;
      }
      curr = token_node_with(curr, token);
      curr->next = malloc(sizeof(TokenNode));
      curr = curr->next;
    }
    lines = lines->next;
    line_count++;
  }
  curr = token_node_with(curr,
                         token_of(++line_count,
                                  "",
                                  ENDFILE,
                                  NIL)
                         );
  return head;
}

int white_space_machine(LineNode* node, int* trts) {
  char* buffer = node->line->value;
  int hare = (*trts);
  while (is_whitespace(buffer[hare])) {
    hare++;
  }
  (*trts) = hare;
  return buffer[hare] == '\0';
}

Token* id_machine(LineNode* node, ReservedWordNode* reserved,
    SymbolNode* symbols, int* trts) {
  char* buffer = node->line->value;
  int hare = (*trts);
  if (is_letter(buffer[hare])) {
    Token* token;
    hare++;
    while (is_letter(buffer[hare]) || is_digit(buffer[hare])) {
      hare++;
    }
    char* lexeme = substring(buffer, (*trts), hare);
    if ((token = check_id_length(node, lexeme, (*trts), hare)) != NULL) {
      (*trts) = hare;
      return token;
    }
    (*trts) = hare;
    while (reserved != NULL) {
      if (is_equal(lexeme, reserved->word->value)) {
        return token_of(node->line->number,
                        lexeme,
                        reserved->word->type,
                        reserved->word->attr);
      }
      reserved = reserved->next;
    }
    token = token_of(node->line->number, lexeme, ID, NIL);
    Attribute attribute;
    attribute.address = save_symbol(symbols, lexeme);
    token->attr = attribute;
    return token;
  }
  return NULL;
}

Token* long_real_machine(LineNode* node, int* trts) {
  char* buffer = node->line->value;
  int hare = (*trts);
  if (is_digit(buffer[hare])) {
    Token* token;
    while(is_digit(buffer[hare])) {
      hare++;
    }
    token = check_xx_length(node,
                            substring(buffer, (*trts), hare),
                            (*trts),
                            hare);
    if (buffer[hare] != '.' || !is_digit(buffer[hare+1])) {
      return NULL;
    }
    hare++;
    while(is_digit(buffer[hare])) {
      hare++;
    }
    token = check_yy_length(node,
                            substring(buffer, (*trts), hare),
                            (*trts),
                            hare);
    if (buffer[hare] != 'E' || !is_digit(buffer[hare+1])) {
      return NULL;
    }
    hare++;
    while(is_digit(buffer[hare])) {
      hare++;
    }
    token = check_xx_length(node,
                            substring(buffer, (*trts), hare),
                            (*trts),
                            hare);
    char* lexeme = substring(buffer, (*trts), hare);
    (*trts) = hare;
    return token_of(node->line->number, lexeme, NUM, LREAL);
  }
  return NULL;
}

Token* real_machine(LineNode* node, int* trts) {
  char* buffer = node->line->value;
  int hare = (*trts);
  if (is_digit(buffer[hare])) {
    while(is_digit(buffer[hare])) {
      hare++;
    }
    if (buffer[hare] != '.' || !is_digit(buffer[hare+1])) {
      return NULL;
    }
    hare++;
    while(is_digit(buffer[hare])) {
      hare++;
    }
    char* lexeme = substring(buffer, (*trts), hare);
    (*trts) = hare;
    return token_of(node->line->number, lexeme, NUM, REAL_);
  }
  return NULL;
}

Token* int_machine(LineNode* node, int* trts) {
  char* buffer = node->line->value;
  int hare = (*trts);
  if (is_digit(buffer[hare])) {
    while(is_digit(buffer[hare])) {
      hare++;
    }
    char* lexeme = substring(buffer, (*trts), hare);
    (*trts) = hare;
    return token_of(node->line->number, lexeme, NUM, INT);
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
      return token_of(node->line->number, lexeme, RELOP, EQ);
    case '<':
      if (buffer[(*trts)+1] == '>') {
        lexeme = substring(buffer, (*trts), (*trts)+2);
        (*trts) = (*trts) + 2;
        return token_of(node->line->number, lexeme, RELOP, NEQ);
      }
      if (buffer[(*trts)+1] == '=') {
        lexeme = substring(buffer, (*trts), (*trts)+2);
        (*trts) = (*trts) + 2;
        return token_of(node->line->number, lexeme, RELOP, LTE);
      }
      lexeme = substring(buffer, (*trts), (*trts)+1);
      (*trts)++;
      return token_of(node->line->number, lexeme, RELOP, LT);
    case '>':
      if (buffer[(*trts)+1] == '=') {
        lexeme = substring(buffer, (*trts), (*trts)+2);
        (*trts) = (*trts) + 2;
        return token_of(node->line->number, lexeme, RELOP, GTE);
      }
      lexeme = substring(buffer, (*trts), (*trts)+1);
      (*trts)++;
      return token_of(node->line->number, lexeme, RELOP, GT);
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
      return token_of(node->line->number, lexeme, MULOP, NIL);
    case '/':
      lexeme = substring(buffer, (*trts), (*trts)+1);
      (*trts)++;
      return token_of(node->line->number, lexeme, MULOP, NIL);
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
      return token_of(node->line->number, lexeme, ASSIGNOP, NIL);
    case '-':
      lexeme = substring(buffer, (*trts), (*trts)+1);
      (*trts)++;
      return token_of(node->line->number, lexeme, ASSIGNOP, NIL);
  }
  return NULL;
}

Token* assignop_machine(LineNode* node, int* trts) {
  char* buffer = node->line->value;
  if (buffer[(*trts)] == ':' && buffer[(*trts)+1] == '=') {
    char* lexeme = substring(buffer, (*trts), (*trts)+2);
    (*trts) = (*trts) + 2;
    return token_of(node->line->number, lexeme, ASSIGNOP, NIL);
  }
  return NULL;
}

Token* catchall_machine(LineNode* node, int* trts) {
  char* buffer = node->line->value;
  if (buffer[(*trts)] == '.' && buffer[(*trts)+1] == '.') {
    char* lexeme = substring(buffer, (*trts), (*trts)+2);
    (*trts)++;
    return token_of(node->line->number, lexeme, RANGE, NIL);
  }
  char* lexeme;
  switch (buffer[(*trts)]) {
    case '(':
      lexeme = substring(buffer, (*trts), (*trts)+1);
      (*trts)++;
      return token_of(node->line->number, lexeme, OPENPAREN, NIL);
    case ')':
      lexeme = substring(buffer, (*trts), (*trts)+1);
      (*trts)++;
      return token_of(node->line->number, lexeme, CLOSEPAREN, NIL);
    case '[':
      lexeme = substring(buffer, (*trts), (*trts)+1);
      (*trts)++;
      return token_of(node->line->number, lexeme, OPENBRACKET, NIL);
    case ']':
      lexeme = substring(buffer, (*trts), (*trts)+1);
      (*trts)++;
      return token_of(node->line->number, lexeme, CLOSEBRACKET, NIL);
    case ':':
      lexeme = substring(buffer, (*trts), (*trts)+1);
      (*trts)++;
      return token_of(node->line->number, lexeme, COLON, NIL);
    case ';':
      lexeme = substring(buffer, (*trts), (*trts)+1);
      (*trts)++;
      return token_of(node->line->number, lexeme, SEMICOLON, NIL);
    case ',':
      lexeme = substring(buffer, (*trts), (*trts)+1);
      (*trts)++;
      return token_of(node->line->number, lexeme, COMMA, NIL);
    case '.':
      lexeme = substring(buffer, (*trts), (*trts)+1);
      (*trts)++;
      return token_of(node->line->number, lexeme, PERIOD, NIL);
  }
  return NULL;
}
