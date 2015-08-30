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
  return print_token_file(head) || print_listing_file(first);
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

int print_token_file(TokenNode* head) {
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
  TokenNode* curr = head;
  while (curr != NULL) {
    Token* token = curr->token;
    fprintf(file,
            "%-10d %-13s %-2d %-14s %-10d\n",
            token->line_number,
            token->lexeme,
            token->type,
            token->annotation,
            token->attr.value);
    curr = curr->next;
  }
  return fclose(file);
}

int print_listing_file(LineNode* head) {
  FILE* file;
  if ((file = fopen("build/listing_file.txt", "w")) == NULL) {
    printf("Cannot create file listing_file.txt\n");
    exit(1);
  }
  LineNode* curr = head;
  while (curr != NULL) {
    fprintf(file, "%4d.    %s", curr->line->number, curr->line->value);
    LineNode* error = curr->error;
    while (error != NULL) {
      fprintf(file, "%4d.    %s", error->line->number, error->line->value);
      error = error->error;
    }
    curr = curr->next;
  }
  return fclose(file);
}

TokenNode* analyze(LineNode* first, ReservedWordNode* reserved) {
  TokenNode* head = malloc(sizeof(TokenNode));
  TokenNode* curr = head;
  LineNode* node = first;
  int line_count = 0;
  int* trts = malloc(sizeof(int));
  while (node != NULL) {
    (*trts) = 0;
    while (node->line->value[(*trts)] != '\0') {
      Token* token;
      if (white_space_machine(node, trts)) {
        continue;
      }
      if ((token = id_machine(node, reserved, trts)) != NULL) {
      } else if ((token = long_real_machine(node, trts)) != NULL) {
      } else if ((token = real_machine(node, trts)) != NULL) {
      } else if ((token = int_machine(node, trts)) != NULL) {
      } else if ((token = relop_machine(node, trts)) != NULL) {
      } else if ((token = addop_machine(node, trts)) != NULL) {
      } else if ((token = mulop_machine(node, trts)) != NULL) {
      } else if ((token = assignop_machine(node, trts)) != NULL) {
      } else if ((token = catchall_machine(node, trts)) != NULL) {
      } else {
        // unrecognized symbol
        (*trts)++;
        continue;
      }
      curr = token_node_with(curr, token);
      curr->next = malloc(sizeof(TokenNode));
      curr = curr->next;
    }
    node = node->next;
    line_count++;
  }
  curr = token_node_with(curr,
                         token_of(++line_count,
                                  "",
                                  ENDFILE,
                                  annotation_of(ENDFILE),
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

Token* id_machine(LineNode* node, ReservedWordNode* reserved, int* trts) {
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
                        annotation_of(reserved->word->type),
                        reserved->word->attr);
      }
      reserved = reserved->next;
    }
    return token_of(node->line->number, lexeme, ID, annotation_of(ID), -2);
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
  return NULL;
}

Token* catchall_machine(LineNode* node, int* trts) {
  return NULL;
}
