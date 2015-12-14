package lib

import "container/list"

func ProgramFollows() *list.List {
  return list.New()
}

func ProgramBodyFollows() *list.List {
  return list.New()
}

func ProgramSubbodyFollows() *list.List {
  return list.New()
}

func IdFollows() *list.List {
  follows := list.New()
  follows.PushBack(COLON)
  return follows
}

func IdentifierListFollows() *list.List {
  follows := list.New()
  follows.PushBack(CLOSE_PAREN)
  return follows
}

func IdentifierListPrimeFollows() *list.List {
  follows := list.New()
  follows.PushBack(CLOSE_PAREN)
  return follows
}

func DeclarationsFollows() *list.List {
  follows := list.New()
  follows.PushBack(FUNCTION)
  follows.PushBack(BEGIN)
  return follows
}

func DeclarationsPrimeFollows() *list.List {
  follows := list.New()
  follows.PushBack(FUNCTION)
  follows.PushBack(BEGIN)
  return follows
}

func TypeFollows() *list.List {
  follows := list.New()
  follows.PushBack(SEMICOLON)
  follows.PushBack(CLOSE_PAREN)
  return follows
}

func StandardTypeFollows() *list.List {
  follows := list.New()
  follows.PushBack(SEMICOLON)
  follows.PushBack(CLOSE_PAREN)
  return follows
}

func SubprogramDeclarationsFollows() *list.List {
  follows := list.New()
  follows.PushBack(BEGIN)
  return follows
}

func SubprogramDeclarationsPrimeFollows() *list.List {
  follows := list.New()
  follows.PushBack(BEGIN)
  return follows
}

func SubprogramDeclarationFollows() *list.List {
  follows := list.New()
  follows.PushBack(SEMICOLON)
  return follows
}

func SubprogramBodyFollows() *list.List {
  follows := list.New()
  follows.PushBack(SEMICOLON)
  return follows
}

func SubprogramSubbodyFollows() *list.List {
  follows := list.New()
  follows.PushBack(SEMICOLON)
  return follows
}

func SubprogramHeadFollows() *list.List {
  follows := list.New()
  follows.PushBack(VAR)
  follows.PushBack(FUNCTION)
  follows.PushBack(BEGIN)
  return follows
}

func SubprogramHeadPrimeFollows() *list.List {
  follows := list.New()
  follows.PushBack(VAR)
  follows.PushBack(FUNCTION)
  follows.PushBack(BEGIN)
  return follows
}

func ArgumentsFollows() *list.List {
  follows := list.New()
  follows.PushBack(COLON)
  return follows
}

func ParameterListFollows() *list.List {
  follows := list.New()
  follows.PushBack(CLOSE_PAREN)
  return follows
}

func ParameterListPrimeFollows() *list.List {
  follows := list.New()
  follows.PushBack(CLOSE_PAREN)
  return follows
}

func CompoundStatementFollows() *list.List {
  follows := list.New()
  follows.PushBack(PERIOD)
  follows.PushBack(SEMICOLON)
  follows.PushBack(ELSE)
  follows.PushBack(END)
  return follows
}

func CompoundStatementPrimeFollows() *list.List {
  follows := list.New()
  follows.PushBack(PERIOD)
  follows.PushBack(SEMICOLON)
  follows.PushBack(ELSE)
  follows.PushBack(END)
  return follows
}

func OptionalStatementsFollows() *list.List {
  follows := list.New()
  follows.PushBack(END)
  return follows
}

func StatementListFollows() *list.List {
  follows := list.New()
  follows.PushBack(END)
  return follows
}

func StatementListPrimeFollows() *list.List {
  follows := list.New()
  follows.PushBack(END)
  return follows
}

func StatementFollows() *list.List {
  follows := list.New()
  follows.PushBack(ELSE)
  follows.PushBack(SEMICOLON)
  follows.PushBack(END)
  return follows
}

func StatementPrimeFollows() *list.List {
  follows := list.New()
  follows.PushBack(ELSE)
  follows.PushBack(SEMICOLON)
  follows.PushBack(END)
  return follows
}

func VariableFollows() *list.List {
  follows := list.New()
  follows.PushBack(ASSIGNOP)
  return follows
}

func VariablePrimeFollows() *list.List {
  follows := list.New()
  follows.PushBack(ASSIGNOP)
  return follows
}

func ExpressionListFollows() *list.List {
  follows := list.New()
  follows.PushBack(CLOSE_PAREN)
  return follows
}

func ExpressionListPrimeFollows() *list.List {
  follows := list.New()
  follows.PushBack(CLOSE_PAREN)
  return follows
}

func ExpressionFollows() *list.List {
  follows := list.New()
  follows.PushBack(ELSE)
  follows.PushBack(SEMICOLON)
  follows.PushBack(END)
  follows.PushBack(THEN)
  follows.PushBack(DO)
  follows.PushBack(CLOSE_BRACKET)
  follows.PushBack(CLOSE_PAREN)
  return follows
}

func ExpressionPrimeFollows() *list.List {
  follows := list.New()
  follows.PushBack(ELSE)
  follows.PushBack(SEMICOLON)
  follows.PushBack(END)
  follows.PushBack(THEN)
  follows.PushBack(DO)
  follows.PushBack(CLOSE_BRACKET)
  follows.PushBack(CLOSE_PAREN)
  return follows
}

func SimpleExpressionFollows() *list.List {
  follows := list.New()
  follows.PushBack(RELOP)
  follows.PushBack(ELSE)
  follows.PushBack(SEMICOLON)
  follows.PushBack(END)
  follows.PushBack(THEN)
  follows.PushBack(DO)
  follows.PushBack(CLOSE_BRACKET)
  follows.PushBack(CLOSE_PAREN)
  return follows
}

func SimpleExpressionPrimeFollows() *list.List {
  follows := list.New()
  follows.PushBack(RELOP)
  follows.PushBack(ELSE)
  follows.PushBack(SEMICOLON)
  follows.PushBack(END)
  follows.PushBack(THEN)
  follows.PushBack(DO)
  follows.PushBack(CLOSE_BRACKET)
  follows.PushBack(CLOSE_PAREN)
  return follows
}

func TermFollows() *list.List {
  follows := list.New()
  follows.PushBack(ADDOP)
  follows.PushBack(RELOP)
  follows.PushBack(ELSE)
  follows.PushBack(SEMICOLON)
  follows.PushBack(END)
  follows.PushBack(THEN)
  follows.PushBack(DO)
  follows.PushBack(CLOSE_BRACKET)
  follows.PushBack(CLOSE_PAREN)
  return follows
}

func TermPrimeFollows() *list.List {
  follows := list.New()
  follows.PushBack(ADDOP)
  follows.PushBack(RELOP)
  follows.PushBack(ELSE)
  follows.PushBack(SEMICOLON)
  follows.PushBack(END)
  follows.PushBack(THEN)
  follows.PushBack(DO)
  follows.PushBack(CLOSE_BRACKET)
  follows.PushBack(CLOSE_PAREN)
  return follows
}

func FactorFollows() *list.List {
  follows := list.New()
  follows.PushBack(MULOP)
  follows.PushBack(ADDOP)
  follows.PushBack(RELOP)
  follows.PushBack(ELSE)
  follows.PushBack(SEMICOLON)
  follows.PushBack(END)
  follows.PushBack(THEN)
  follows.PushBack(DO)
  follows.PushBack(CLOSE_BRACKET)
  follows.PushBack(CLOSE_PAREN)
  return follows
}

func FactorPrimeFollows() *list.List {
  follows := list.New()
  follows.PushBack(MULOP)
  follows.PushBack(ADDOP)
  follows.PushBack(RELOP)
  follows.PushBack(ELSE)
  follows.PushBack(SEMICOLON)
  follows.PushBack(END)
  follows.PushBack(THEN)
  follows.PushBack(DO)
  follows.PushBack(CLOSE_BRACKET)
  follows.PushBack(CLOSE_PAREN)
  return follows
}

func SignFollows() *list.List {
  follows := list.New()
  follows.PushBack(ID)
  follows.PushBack(NUM)
  follows.PushBack(CLOSE_PAREN)
  follows.PushBack(NOT)
  return follows
}