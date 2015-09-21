#ifndef GLOBAL_H
#define GLOBAL_H
#include "types.h"

const int MAX_BUFFER_LENGTH = 72;
size_t const MAX_BUFFER_SIZE = ((MAX_BUFFER_LENGTH+1) * sizeof(char));
char* const LINE_TOO_LONG = "ERROR: Lines can be only 72 characters long.\n";

const int MAX_ID_LENGTH = 10;
size_t const MAX_ID_SIZE = ((MAX_ID_LENGTH+1) * sizeof(char));

const int MAX_INTEGER_LENGTH = 10;
const int MAX_XX_LENGTH = 5;
const int MAX_YY_LENGTH = 5;
const int MAX_ZZ_LENGTH = 2;

ReservedWordNode* reserved_word_table = NULL;
ErrorNode* error_list = NULL;
SymbolNode* symbol_table = NULL;

#endif
