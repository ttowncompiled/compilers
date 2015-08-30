#include <stdio.h>
#include <stdlib.h>
#include "lexer.h"

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

ReservedWordNode* load_reserved_words() {
  FILE* file;
  if((file = fopen("build/reserved_words.txt", "r")) == NULL) {
    printf("Cannot open file reserved_words.txt\n");
    exit(1);
  }
  ReservedWordNode* head = malloc(sizeof(ReservedWordNode));
  size_t buffer_size = MAX_BUFFER_SIZE;
  char* buffer = malloc(buffer_size);
  if (-1 == getline(&buffer, &buffer_size, file) || buffer[0] != '"') {
    return head;
  }
  head = word_node_from(buffer);
  ReservedWordNode* prev = head;
  while (-1 != getline(&buffer, &buffer_size, file) && buffer[0] == '"') {
    ReservedWordNode* curr = word_node_from(buffer);
    prev->next = curr;
    prev = curr;
  }
  fclose(file);
  return head;
}

LineNode* organize(char* filename) {
  FILE* file;
  if ((file = fopen(filename, "r")) == NULL) {
    printf("Cannot open file %s\n", filename);
    exit(1);
  }
  LineNode* head = malloc(sizeof(LineNode));
  size_t buffer_size = MAX_BUFFER_SIZE;
  char* buffer = malloc(buffer_size);
  if (-1 == getline(&buffer, &buffer_size, file)) {
    return head;
  }
  int line_number = 1;
  head = line_node_of(buffer, line_number);
  check_buffer_size(buffer_size, head);
  buffer = malloc(MAX_BUFFER_SIZE);
  LineNode* prev = head;
  while (-1 != getline(&buffer, &buffer_size, file)) {
    LineNode* curr = line_node_of(buffer, ++line_number);
    check_buffer_size(buffer_size, curr);
    prev->next = curr; 
    prev = curr;
    buffer = malloc(MAX_BUFFER_SIZE);
  }
  return head;
}

int print_token_file(TokenNode* tokens) {
  FILE* file;
  if ((file = fopen("build/token_file.txt", "w")) == NULL) {
    printf("Cannot create file token_file.txt\n");
    exit(1);
  }
  fprintf(file,
          "%-10s %-13s %-17s %-10s\n",
          "Line No.",
          "Lexeme",
          "TOKEN-TYPE",
          "ATTRIBUTE");
  while (tokens != NULL) {
    Token* token = tokens->token;
    if (token->attr.value < -1) {
      SymbolNode* address = token->attr.address;
      fprintf(file,
              "%-10d %-13s %-2d %-14s %p %-10s\n",
              token->line_number,
              token->lexeme,
              token->type,
              type_annotation_of(token->type),
              address,
              attr_annotation_of(token->attr.value));
    } else {
      fprintf(file,
              "%-10d %-13s %-2d %-14s %-14d %-10s\n",
              token->line_number,
              token->lexeme,
              token->type,
              type_annotation_of(token->type),
              token->attr.value,
              attr_annotation_of(token->attr.value));
    }
    tokens = tokens->next;
  }
  return fclose(file);
}

int print_listing_file(LineNode* lines) {
  FILE* file;
  if ((file = fopen("build/listing_file.txt", "w")) == NULL) {
    printf("Cannot create file listing_file.txt\n");
    exit(1);
  }
  while (lines != NULL) {
    fprintf(file, "%4d.    %s", lines->line->number, lines->line->value);
    LineNode* error = lines->error;
    while (error != NULL) {
      fprintf(file, "%4d.    %s", error->line->number, error->line->value);
      error = error->error;
    }
    lines = lines->next;
  }
  return fclose(file);
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
      } else if ((token = addop_machine(lines, trts)) != NULL) {
      } else if ((token = mulop_machine(lines, trts)) != NULL) {
      } else if ((token = assignop_machine(lines, trts)) != NULL) {
      } else if ((token = catchall_machine(lines, trts)) != NULL) {
      } else {
        // unrecognized symbol
        (*trts)++;
        continue;
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
    hare++;
    while (is_letter(buffer[hare]) || is_digit(buffer[hare])) {
      hare++;
    }
    char* lexeme = substring(buffer, (*trts), hare);
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
    Token* token = token_of(node->line->number, lexeme, ID, NIL);
    Attribute attribute;
    attribute.address = save_symbol(symbols, lexeme);
    token->attr = attribute;
    return token;
  }
  return NULL;
}

Token* long_real_machine(LineNode* node, int* trts) {
  return NULL;
}

Token* real_machine(LineNode* node, int* trts) {
  return NULL;
}

Token* int_machine(LineNode* node, int* trts) {
  return NULL;
}

Token* relop_machine(LineNode* node, int* trts) {
  return NULL;
}

Token* addop_machine(LineNode* node, int* trts) {
  return NULL;
}

Token* mulop_machine(LineNode* node, int* trts) {
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
    (*trts) = (*trts) + 1;
    return token_of(node->line->number, lexeme, RANGE, NIL);
  }
  char* lexeme;
  switch (buffer[(*trts)]) {
    case '(':
      lexeme = substring(buffer, (*trts), (*trts)+1);
      (*trts) = (*trts) + 1;
      return token_of(node->line->number, lexeme, OPENPAREN, NIL);
    case ')':
      lexeme = substring(buffer, (*trts), (*trts)+1);
      (*trts) = (*trts) + 1;
      return token_of(node->line->number, lexeme, CLOSEPAREN, NIL);
    case '[':
      lexeme = substring(buffer, (*trts), (*trts)+1);
      (*trts) = (*trts) + 1;
      return token_of(node->line->number, lexeme, OPENBRACKET, NIL);
    case ']':
      lexeme = substring(buffer, (*trts), (*trts)+1);
      (*trts) = (*trts) + 1;
      return token_of(node->line->number, lexeme, CLOSEBRACKET, NIL);
    case ':':
      lexeme = substring(buffer, (*trts), (*trts)+1);
      (*trts) = (*trts) + 1;
      return token_of(node->line->number, lexeme, COLON, NIL);
    case ';':
      lexeme = substring(buffer, (*trts), (*trts)+1);
      (*trts) = (*trts) + 1;
      return token_of(node->line->number, lexeme, SEMICOLON, NIL);
    case ',':
      lexeme = substring(buffer, (*trts), (*trts)+1);
      (*trts) = (*trts) + 1;
      return token_of(node->line->number, lexeme, COMMA, NIL);
    case '.':
      lexeme = substring(buffer, (*trts), (*trts)+1);
      (*trts) = (*trts) + 1;
      return token_of(node->line->number, lexeme, PERIOD, NIL);
  }
  return NULL;
}
