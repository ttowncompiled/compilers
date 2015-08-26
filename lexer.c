#include "lexer.h"

int main(int argc, char *argv[]) {
  if (argc > 2) {
    printf("This compiler only accepts one source file.");
    exit(1);
  }
  if (argc < 2) {
    printf("Please provide the name of a source file.");
    return 0;
  }
  printListingFile(argv[1]);
  return 0;
}

void printListingFile(char *filename) {
  FILE *file;
  char *line;
  int hare;
  int lineNumber;

  file = fopen(filename, "r");
  line = malloc(72 * sizeof(char));
  hare = 0;
  lineNumber = 1;

  char c;
  while ((c = (char)fgetc(file)) != EOF) {
    line[hare++] = c;

    if (c == '\n') {
      line[hare] = '\0';
      printf("%4d.    %s", lineNumber, line);
      printf("%c", '\n');

      line = malloc(72 * sizeof(char));
      hare = 0;
      lineNumber++;
    }
  }
}

