// a lexical analyzer for a simplified version of Pascal.
#include "lexer.h"

int main(int argc, char* argv[]) {
  // check that we have been given the right amount of input
  if (argc > 2) {
    printf("This compiler only accepts one source file.\n");
    exit(1);
  }
  if (argc < 2) {
    printf("Please provide the name of a source file.\n");
    return 0;
  }

  FILE* file;
  Line* head;
  int hare;
  
  file = fopen(argv[1], "r");
  head = malloc(sizeof(Line));
  head->value = malloc((BUFFER_SIZE+1) * sizeof(char));
  hare = 0;

  char c;
  Line* line;
  line = head;
  while ((c = (char)fgetc(file)) != EOF) {
    line->value[hare++] = c;
    if (c == '\n') {
      Line* nextLine;
      nextLine = malloc(sizeof(Line));
      nextLine->value = malloc((BUFFER_SIZE+1) * sizeof(char));
      line->value[hare] = '\0';
      line->next = nextLine;
      line = nextLine;
      hare = 0;
    }
  }

  return printListingFile(head);
}

int printListingFile(Line* head) {
  int lineNumber;
  lineNumber = 1;

  Line* line;
  line = head;
  while (line != NULL && line->value[0] != '\0') {
    printf("%4d.    %s", lineNumber, line->value);
    printf("\n");
    line = line->next;
    lineNumber++;
  }
  
  return 0;
}

