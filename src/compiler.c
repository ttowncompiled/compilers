#include <stdio.h>
#include <stdlib.h>
#include "../lib/io.h"
#include "../lib/lexer.h"
#include "../lib/parser.h"

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
  print_listing_file(lines);
  tokenize(lines);
  print_token_file(tokens);
  program();
  return 0;
}

