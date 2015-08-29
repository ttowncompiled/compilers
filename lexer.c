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
  LineNode* first = organize(filename);
  TokenNode* head = analyze(first, reserved);
  print_token_file(head);
  return print_listing_file(first);
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

ReservedWordNode* word_node_from(char* entry) {
  char c;
  char* value = malloc(MAX_ID_SIZE);
  int type = 0;
  int attr = 0;
  int hare = 0;
  int trts = 0;
  hare++;
  while ((c = entry[hare++]) != '"') {
    value[trts++] = c;
  }
  value[trts] = '\0';
  hare++;
  while ((c = entry[hare++]) != ' ') {
    type = (type*10) + (c-48);
  }
  while ((c = entry[hare++]) != '\n') {
    attr = (attr*10) + (c-48);
  }
  return word_node_of(value, type, attr);
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
  assert_buffer_size(buffer_size, head);
  buffer = malloc(MAX_BUFFER_SIZE);
  LineNode* prev = head;
  while (-1 != getline(&buffer, &buffer_size, file)) {
    LineNode* curr = line_node_of(buffer, ++line_number);
    assert_buffer_size(buffer_size, curr);
    prev->next = curr; 
    prev = curr;
    buffer = malloc(MAX_BUFFER_SIZE);
  }
  return head;
}

TokenNode* analyze(LineNode* first, ReservedWordNode* reserved) {
  TokenNode* head = malloc(sizeof(TokenNode));
  TokenNode* curr = head;
  LineNode* node = first;
  int line_count = 0;
  while (node != NULL) {
    char* buffer = node->line->value;
    int hare = 0;
    while (buffer[hare] != '\0') {
      hare++;
    }
    // white space machine
    // id machine - check reserved words
    // long real machine
    // real machine
    // int machine
    // relop machine
    // addop machine
    // mulop machine
    // assignop machine
    // unrecognized symbol
    node = node->next;
    line_count++;
  }
  // EOF token
  curr = token_node_with(curr, ++line_count, "EOF", 0, "(EOF)", 0);
  return head;
}

int print_token_file(TokenNode* head) {
  FILE* file;
  if ((file = fopen("build/token_file.txt", "w")) == NULL) {
    printf("Cannot create file token_file.txt\n");
    exit(1);
  }
  fprintf(file, "%-10s %-13s %-17s %-10s\n", "Line No.", "Lexeme", "TOKEN-TYPE", "ATTRIBUTE");
  TokenNode* curr = head;
  while (curr != NULL) {
    Token* token = curr->token;
    fprintf(file, "%-10d %-13s %-2d %-14s %-10d\n", token->line_number, token->lexeme, token->type, token->annotation, token->attr);
    curr = curr->next;
  }
  return fclose(file);
}

int print_listing_file(LineNode* head) {
  LineNode* curr = head;
  printf("\n");
  while (curr != NULL) {
    printf("%4d.    %s", curr->line->number, curr->line->value);
    LineNode* error = curr->error;
    while (error != NULL) {
      printf("%4d.    %s", error->line->number, error->line->value);
      error = error->error;
    }
    curr = curr->next;
    printf("\n");
  }
  return 0;
}
