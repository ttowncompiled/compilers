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
  c = (char)fgetc(file);
  line = head;
  while (c != EOF) {
    line->value[hare++] = c;
    if (c != '\n') {
      c = (char)fgetc(file);
    } else {
      line->value[hare] = '\0';
      hare = 0;
      if ((c = (char)fgetc(file)) != EOF) {
        Line* nextLine;
        nextLine = malloc(sizeof(Line));
        nextLine->value = malloc((BUFFER_SIZE+1) * sizeof(char));
        line->next = nextLine;
        line = nextLine;
      }
    }
  }

  loadReservedWords();
  return printListingFile(head);
}

Word* loadReservedWords() {
  FILE* file;
  Word* head;
  int hare;

  file = fopen("reserved_words.txt", "r");
  head = malloc(sizeof(Word));
  head->value = malloc(10 * sizeof(char));
  hare = 0;

  char c;
  Word* word;
  c = (char)fgetc(file);
  word = head;
  while (c != EOF) {
    // grab the reserved word
    if (c >= 'a' && c <= 'z') {
      word->value[hare++] = c;
      c = (char)fgetc(file);
    } else if (c >= '0' && c <= '9') {
      int type;
      int attr;
      type = c-48;
      // grab the token type
      while ((c = (char)fgetc(file)) >= '0' && c <= '9') {
        type = (type*10) + (c-48);
      }
      // skip to the attribute value
      while (c < '0' || c > '9') {
        c = (char)fgetc(file);
      }
      attr = c-48;
      // grab the attribute
      while ((c = (char)fgetc(file)) >= '0' && c <= '9') {
        attr = (attr*10) + (c-48);
      }
      word->type = type;
      word->attr = attr;
      // skip to the end of line
      while (c != '\n') {
        c = (char)fgetc(file);
      }
      word->value[hare] = '\0';
      hare = 0;
      // check if there is a next line
      if ((c = (char)fgetc(file)) != EOF && c != ' ') {
        Word* nextWord;
        nextWord = malloc(sizeof(Word));
        nextWord->value = malloc(10 * sizeof(char));
        word->next = nextWord;
        word = nextWord;
      }
    } else {
      c = (char)fgetc(file);
    }
  }

  return head;
}

int printListingFile(Line* head) {
  int lineNumber;
  lineNumber = 1;

  Line* line;
  line = head;
  while (line != NULL) {
    printf("%4d.    %s", lineNumber, line->value);
    printf("\n");
    line = line->next;
    lineNumber++;
  }
  
  return 0;
}

