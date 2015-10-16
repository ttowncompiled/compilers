#include <stdio.h>
#include <stdlib.h>
#include "../lib/io.h"
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
  load_reserved_words();
  char* filename = argv[1];
  LineNode* lines = read_file(filename);
  TokenNode* tokens = tokenize(lines);
  return print_token_file(tokens) || print_listing_file(lines);
}

