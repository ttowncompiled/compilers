#ifndef IO_H 
#define IO_H
#include "global.h"
#include "types.h"
#include "types_util.h"

void check_buffer_size(size_t buffer_size, LineNode* node) {
  if (buffer_size <= MAX_BUFFER_SIZE) {
    return;
  }
  throw_error(node->line->number, LINE_TOO_LONG);
}

void load_reserved_words() {
  FILE* file;
  if((file = fopen("build/reserved_words.txt", "r")) == NULL) {
    printf("Cannot open file reserved_words.txt\n");
    exit(1);
  }
  ReservedWordNode* head = malloc(sizeof(ReservedWordNode));
  size_t buffer_size = MAX_BUFFER_SIZE;
  char* buffer = malloc(buffer_size);
  if (-1 == getline(&buffer, &buffer_size, file) || buffer[0] != '"') {
    reserved_word_table = head;
    return;
  }
  head = reserved_word_node_from(buffer);
  ReservedWordNode* prev = head;
  while (-1 != getline(&buffer, &buffer_size, file) && buffer[0] == '"') {
    ReservedWordNode* curr = reserved_word_node_from(buffer);
    prev->next = curr;
    prev = curr;
  }
  fclose(file);
  reserved_word_table = head;
}

LineNode* read_file(char* filename) {
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
  head = new_line_node(line_number, buffer);
  check_buffer_size(buffer_size, head);
  buffer = malloc(MAX_BUFFER_SIZE);
  LineNode* prev = head;
  while (-1 != getline(&buffer, &buffer_size, file)) {
    LineNode* curr = new_line_node(++line_number, buffer);
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
              attr_annotation_of(token->type, token->attr.value));
    } else {
      fprintf(file,
              "%-10d %-13s %-2d %-14s %-14d %-10s\n",
              token->line_number,
              token->lexeme,
              token->type,
              type_annotation_of(token->type),
              token->attr.value,
              attr_annotation_of(token->type, token->attr.value));
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
  ErrorNode* errors = error_list;
  while (lines != NULL) {
    fprintf(file, "%4d.    %s", lines->line->number, lines->line->value);
    while (errors != NULL && errors->error->line_number == lines->line->number) {
      fprintf(file, "%4d.    %s", errors->error->line_number, errors->error->reason);
      errors = errors->next;
    }
    lines = lines->next;
  }
  return fclose(file);
}

#endif
