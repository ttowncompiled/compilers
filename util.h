#ifndef UTIL_H
#define UTIL_H

int is_equal(char* left, char* right) {
  int idx = 0;
  while (left[idx] != '\0' && right[idx] != '\0' && left[idx] == right[idx]) {
    idx++;
  }
  return left[idx] == right[idx];
}

int is_digit(char c) {
  return c >= '0' && c <= '9';
}

int is_letter(char c) {
  return (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z');
}

int is_whitespace(char c) {
  return c == ' ' || c == '\t' || c == '\n';
}

char* substring(char* string, int first, int last) {
  if (last <= first) {
    return NULL;
  }
  char* sub = malloc((last-first+1) * sizeof(char));
  int idx = 0;
  while (idx < last-first) {
    sub[idx] = string[first+idx];
    idx++;
  }
  sub[idx] = '\0';
  return sub;
}

#endif
