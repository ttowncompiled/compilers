package compiler

import (
  "container/list"
  "lib"
)

func match(tokens *list.List, expectedType int) (lib.Token, bool) {
  t := tokens.Front().Value.(lib.Token)
  if (t.Type == expectedType) {
    return t, true
  }
  return t, false
}

func matchYank(tokens *list.List, expectedType int) (lib.Token, bool) {
  t := tokens.Front().Value.(lib.Token)
  if (t.Type == expectedType) {
    tokens.Remove(tokens.Front())
    return t, true
  }
  return t, false
}

func report(listing *list.List, expected string, t lib.Token) {
  if t.Type == lib.EOF {
    return
  }
  e := listing.Front()
  for i := 1; i < t.LineNumber; i++ {
    e = e.Next()
  }
  line := e.Value.(lib.Line)
  line.Errors.PushBack(lib.Error{"SYNERR: EXPECTED: " + expected + " RECEIVED: " + t.Lexeme, &t})
}

func sync(tokens *list.List, follows *list.List) bool {
  t := tokens.Front()
  for t.Value.(lib.Token).Type != lib.EOF {
    for f := follows.Front(); f != nil; f = f.Next() {
      if (f.Value.(int) == t.Value.(lib.Token).Type) {
        return true
      }
    }
    t = t.Next()
    tokens.Remove(tokens.Front())
  }
  return false
}

func identifierListPrime(listing *list.List, tokens *list.List) {
  t, ok := matchYank(tokens, lib.COMMA)
  if ok {
    if t, ok = matchYank(tokens, lib.ID); !ok {
      report(listing, "ID", t)
      sync(tokens, lib.IdentifierListPrimeFollows())
      return
    }
    identifierListPrime(listing, tokens)
    return
  }
  t, ok = match(tokens, lib.CLOSE_PAREN)
  if !ok {
    report(listing, ", OR )", t)
    sync(tokens, lib.IdentifierListPrimeFollows())
    return
  }
  // epsilon production
}

func identifierList(listing *list.List, tokens *list.List) {
  t, ok := matchYank(tokens, lib.ID)
  if !ok {
    report(listing, "ID", t)
    sync(tokens, lib.IdentifierListFollows())
    return
  }
  identifierListPrime(listing, tokens)
}

func standardType(listing *list.List, tokens *list.List) {
  t, ok := matchYank(tokens, lib.INTEGER)
  if ok {
    return
  }
  t, ok = matchYank(tokens, lib.REAL)
  if !ok {
    report(listing, "integer OR real", t)
    sync(tokens, lib.StandardTypeFollows())
    return
  }
}

func type_(listing *list.List, tokens *list.List) {
  t, ok := match(tokens, lib.INTEGER)
  if !ok {
    t, ok = match(tokens, lib.REAL)
  }
  if ok {
    standardType(listing, tokens)
    return
  }
  t, ok = matchYank(tokens, lib.ARRAY)
  if !ok {
    report(listing, "integer OR real OR array", t)
    sync(tokens, lib.TypeFollows())
    return
  }
  if t, ok = matchYank(tokens, lib.OPEN_BRACKET); !ok {
    report(listing, "[", t)
    sync(tokens, lib.TypeFollows())
    return
  }
  if t, ok = matchYank(tokens, lib.NUM); !ok {
    report(listing, "NUM", t)
    sync(tokens, lib.TypeFollows())
    return
  }
  if t, ok = matchYank(tokens, lib.RANGE); !ok {
    report(listing, "..", t)
    sync(tokens, lib.TypeFollows())
    return
  }
  if t, ok = matchYank(tokens, lib.NUM); !ok {
    report(listing, "NUM", t)
    sync(tokens, lib.TypeFollows())
    return
  }
  if t, ok = matchYank(tokens, lib.CLOSE_BRACKET); !ok {
    report(listing, "]", t)
    sync(tokens, lib.TypeFollows())
    return
  }
  if t, ok = matchYank(tokens, lib.OF); !ok {
    report(listing, "of", t)
    sync(tokens, lib.TypeFollows())
    return
  }
  standardType(listing, tokens)
}

func declarationsPrime(listing *list.List, tokens *list.List) {
  t, ok := matchYank(tokens, lib.VAR)
  if ok {
    if t, ok = matchYank(tokens, lib.ID); !ok {
      report(listing, "ID", t)
      sync(tokens, lib.DeclarationsPrimeFollows())
      return
    }
    if t, ok = matchYank(tokens, lib.COLON); !ok {
      report(listing, ":", t)
      sync(tokens, lib.DeclarationsPrimeFollows())
      return
    }
    type_(listing, tokens)
    if t, ok = matchYank(tokens, lib.SEMICOLON); !ok {
      report(listing, ";", t)
      sync(tokens, lib.DeclarationsPrimeFollows())
      return
    }
    declarationsPrime(listing, tokens)
    return
  }
  t, ok = match(tokens, lib.FUNCTION)
  if !ok {
    t, ok = match(tokens, lib.BEGIN)
  }
  if !ok {
    report(listing, "var OR function OR begin", t)
    sync(tokens, lib.DeclarationsPrimeFollows())
    return
  }
  // epsilon production
}

func declarations(listing *list.List, tokens *list.List) {
  t, ok := matchYank(tokens, lib.VAR)
  if !ok {
    report(listing, "var", t)
    sync(tokens, lib.DeclarationsFollows())
    return
  }
  if t, ok = matchYank(tokens, lib.ID); !ok {
    report(listing, "ID", t)
    sync(tokens, lib.DeclarationsFollows())
    return
  }
  if t, ok = matchYank(tokens, lib.COLON); !ok {
    report(listing, ":", t)
    sync(tokens, lib.DeclarationsFollows())
    return
  }
  type_(listing, tokens)
  if t, ok = matchYank(tokens, lib.SEMICOLON); !ok {
    report(listing, ";", t)
    sync(tokens, lib.DeclarationsFollows())
    return
  }
  declarationsPrime(listing, tokens)
}

func parameterListPrime(listing *list.List, tokens *list.List) {
  t, ok := matchYank(tokens, lib.SEMICOLON)
  if ok {
    if t, ok = matchYank(tokens, lib.ID); !ok {
      report(listing, "ID", t)
      sync(tokens, lib.ParameterListPrimeFollows())
      return
    }
    if t, ok = matchYank(tokens, lib.COLON); !ok {
      report(listing, ":", t)
      sync(tokens, lib.ParameterListPrimeFollows())
      return
    }
    type_(listing, tokens)
    parameterListPrime(listing, tokens)
    return
  }
  if t, ok = match(tokens, lib.CLOSE_PAREN); !ok {
    report(listing, "; OR )", t)
    sync(tokens, lib.ParameterListPrimeFollows())
    return
  }
  // epsilon production
}

func parameterList(listing *list.List, tokens *list.List) {
  t, ok := matchYank(tokens, lib.ID)
  if !ok {
    report(listing, "ID", t)
    sync(tokens, lib.ParameterListFollows())
    return
  }
  if t, ok = matchYank(tokens, lib.COLON); !ok {
    report(listing, ":", t)
    sync(tokens, lib.ParameterListFollows())
    return
  }
  type_(listing, tokens)
  parameterListPrime(listing, tokens)
}

func arguments(listing *list.List, tokens *list.List) {
  t, ok := matchYank(tokens, lib.OPEN_PAREN)
  if !ok {
    report(listing, "(", t)
    sync(tokens, lib.ArgumentsFollows())
    return
  }
  parameterList(listing, tokens)
  if t, ok = matchYank(tokens, lib.CLOSE_PAREN); !ok {
    report(listing, ")", t)
    sync(tokens, lib.ArgumentsFollows())
    return
  }
}

func subprogramHeadPrime(listing *list.List, tokens *list.List) {
  t, ok := match(tokens, lib.OPEN_PAREN)
  if ok {
    arguments(listing, tokens)
    if t, ok = matchYank(tokens, lib.COLON); !ok {
      report(listing, ":", t)
      sync(tokens, lib.SubprogramHeadPrimeFollows())
      return
    }
    standardType(listing, tokens)
    if t, ok = matchYank(tokens, lib.SEMICOLON); !ok {
      report(listing, ";", t)
      sync(tokens, lib.SubprogramHeadPrimeFollows())
      return
    }
    return
  }
  t, ok = matchYank(tokens, lib.COLON)
  if !ok {
    report(listing, "( OR :", t)
    sync(tokens, lib.SubprogramHeadPrimeFollows())
    return
  }
  standardType(listing, tokens)
  if t, ok = matchYank(tokens, lib.SEMICOLON); !ok {
    report(listing, ";", t)
    sync(tokens, lib.SubprogramHeadPrimeFollows())
    return
  }
}

func subprogramHead(listing *list.List, tokens *list.List) {
  t, ok := matchYank(tokens, lib.FUNCTION)
  if !ok {
    report(listing, "function", t)
    sync(tokens, lib.SubprogramHeadFollows())
    return
  }
  if t, ok = matchYank(tokens, lib.ID); !ok {
    report(listing, "ID", t)
    sync(tokens, lib.SubprogramHeadFollows())
    return
  }
  subprogramHeadPrime(listing, tokens)
}

func statementPrime(listing *list.List, tokens *list.List) {
  t, ok := matchYank(tokens, lib.ELSE)
  if ok {
    statement(listing, tokens)
    return
  }
  t, ok = match(tokens, lib.SEMICOLON)
  if !ok {
    t, ok = match(tokens, lib.END)
    if !ok {
      report(listing, "else OR ; OR end", t)
      sync(tokens, lib.StatementPrimeFollows())
      return
    }
  }
  // epsilon production
}

func expressionListPrime(listing *list.List, tokens *list.List) {
  t, ok := matchYank(tokens, lib.COMMA)
  if ok {
    expression(listing, tokens)
    expressionListPrime(listing, tokens)
    return
  }
  t, ok = match(tokens, lib.CLOSE_PAREN)
  if !ok {
    report(listing, ", OR )", t)
    sync(tokens, lib.ExpressionListPrimeFollows())
    return
  }
  // epsilon production
}

func expressionList(listing *list.List, tokens *list.List) {
  t, ok := match(tokens, lib.ID)
  if !ok {
    t, ok = match(tokens, lib.NUM)
  }
  if !ok {
    t, ok = match(tokens, lib.OPEN_PAREN)
  }
  if !ok {
    t, ok = match(tokens, lib.NOT)
  }
  if !ok {
    t, ok = match(tokens, lib.ADDOP)
  }
  if !ok || (t.Type == lib.ADDOP && t.Attr != lib.PLUS && t.Attr != lib.MINUS) {
    report(listing, "ID OR NUM OR ( OR not OR + OR -", t)
    sync(tokens, lib.ExpressionListFollows())
    return
  }
  expression(listing, tokens)
  expressionListPrime(listing, tokens)
}

func factorPrime(listing *list.List, tokens *list.List) {
  t, ok := matchYank(tokens, lib.OPEN_PAREN)
  if ok {
    expressionList(listing, tokens)
    if t, ok = matchYank(tokens, lib.CLOSE_PAREN); !ok {
      report(listing, ")", t)
      sync(tokens, lib.FactorPrimeFollows())
      return
    }
    return
  }
  t, ok = matchYank(tokens, lib.OPEN_BRACKET)
  if ok {
    expression(listing, tokens)
    if t, ok = matchYank(tokens, lib.CLOSE_BRACKET); !ok {
      report(listing, "]", t)
      sync(tokens, lib.FactorPrimeFollows())
      return
    }
    return
  }
  t, ok = match(tokens, lib.MULOP)
  if !ok {
    t, ok = match(tokens, lib.ADDOP)
  }
  if !ok {
    t, ok = match(tokens, lib.RELOP)
  }
  if !ok {
    t, ok = match(tokens, lib.ELSE)
  }
  if !ok {
    t, ok = match(tokens, lib.SEMICOLON)
  }
  if !ok {
    t, ok = match(tokens, lib.END)
  }
  if !ok {
    t, ok = match(tokens, lib.THEN)
  }
  if !ok {
    t, ok = match(tokens, lib.DO)
  }
  if !ok {
    t, ok = match(tokens, lib.CLOSE_BRACKET)
  }
  if !ok {
    t, ok = match(tokens, lib.COMMA)
  }
  if !ok {
    t, ok = match(tokens, lib.CLOSE_PAREN)
  }
  if !ok {
    report(listing, "( OR [ OR MULOP OR ADDOP OR RELOP OR else OR ; OR end OR then OR do OR ] OR )", t)
    sync(tokens, lib.FactorPrimeFollows())
    return
  }
  // epsilon production
}

func factor(listing *list.List, tokens *list.List) {
  t, ok := matchYank(tokens, lib.ID)
  if ok {
    factorPrime(listing, tokens)
    return
  }
  t, ok = matchYank(tokens, lib.NUM)
  if ok {
    return
  }
  t, ok = matchYank(tokens, lib.OPEN_PAREN)
  if ok {
    expression(listing, tokens)
    if t, ok = matchYank(tokens, lib.CLOSE_PAREN); !ok {
      report(listing, ")", t)
      sync(tokens, lib.FactorFollows())
      return
    }
    return
  }
  t, ok = matchYank(tokens, lib.NOT)
  if !ok {
    report(listing, "ID OR NUM OR ( OR not", t)
    sync(tokens, lib.FactorFollows())
    return
  }
  factor(listing, tokens)
}

func termPrime(listing *list.List, tokens *list.List) {
  t, ok := matchYank(tokens, lib.MULOP)
  if ok {
    factor(listing, tokens)
    termPrime(listing, tokens)
    return
  }
  t, ok = match(tokens, lib.ADDOP)
  if !ok {
    t, ok = match(tokens, lib.RELOP)
  }
  if !ok {
    t, ok = match(tokens, lib.ELSE)
  }
  if !ok {
    t, ok = match(tokens, lib.SEMICOLON)
  }
  if !ok {
    t, ok = match(tokens, lib.END)
  }
  if !ok {
    t, ok = match(tokens, lib.THEN)
  }
  if !ok {
    t, ok = match(tokens, lib.DO)
  }
  if !ok {
    t, ok = match(tokens, lib.CLOSE_BRACKET)
  }
  if !ok {
    t, ok = match(tokens, lib.COMMA)
  }
  if !ok {
    t, ok = match(tokens, lib.CLOSE_PAREN)
  }
  if !ok {
    report(listing, "MULOP OR ADDOP OR RELOP OR else OR ; OR end OR then OR do OR ] OR )", t)
    sync(tokens, lib.TermPrimeFollows())
    return
  }
  // epsilon production
}

func term(listing *list.List, tokens *list.List) {
  t, ok := match(tokens, lib.ID)
  if !ok {
    t, ok = match(tokens, lib.NUM)
  }
  if !ok {
    t, ok = match(tokens, lib.OPEN_PAREN)
  }
  if !ok {
    t, ok = match(tokens, lib.NOT)
  }
  if !ok {
    report(listing, "ID OR NUM OR ( OR not", t)
    sync(tokens, lib.TermFollows())
    return
  }
  factor(listing, tokens)
  termPrime(listing, tokens)
}

func simpleExpressionPrime(listing *list.List, tokens *list.List) {
  t, ok := matchYank(tokens, lib.ADDOP)
  if ok {
    term(listing, tokens)
    simpleExpressionPrime(listing, tokens)
    return
  }
  t, ok = match(tokens, lib.RELOP)
  if !ok {
    t, ok = match(tokens, lib.ELSE)
  }
  if !ok {
    t, ok = match(tokens, lib.SEMICOLON)
  }
  if !ok {
    t, ok = match(tokens, lib.END)
  }
  if !ok {
    t, ok = match(tokens, lib.THEN)
  }
  if !ok {
    t, ok = match(tokens, lib.DO)
  }
  if !ok {
    t, ok = match(tokens, lib.CLOSE_BRACKET)
  }
  if !ok {
    t, ok = match(tokens, lib.COMMA)
  }
  if !ok {
    t, ok = match(tokens, lib.CLOSE_PAREN)
  }
  if !ok {
    report(listing, "ADDOP OR RELOP OR else OR ; OR end OR then OR do OR ] OR )", t)
    sync(tokens, lib.SimpleExpressionPrimeFollows())
    return
  }
  // epsilon production
}

func sign(listing *list.List, tokens *list.List) {
  t, ok := matchYank(tokens, lib.ADDOP)
  if !ok || (t.Attr != lib.PLUS && t.Attr != lib.MINUS) {
    report(listing, "+ OR -", t)
    sync(tokens, lib.SignFollows())
    return
  }
}

func simpleExpression(listing *list.List, tokens *list.List) {
  t, ok := match(tokens, lib.ID)
  if !ok {
    t, ok = match(tokens, lib.NUM)
  }
  if !ok {
    t, ok = match(tokens, lib.OPEN_PAREN)
  }
  if !ok {
    t, ok = match(tokens, lib.NOT)
  }
  if ok {
    term(listing, tokens)
    simpleExpressionPrime(listing, tokens)
    return
  }
  t, ok = match(tokens, lib.ADDOP)
  if !ok || (t.Type == lib.ADDOP && t.Attr != lib.PLUS && t.Attr != lib.MINUS) {
    report(listing, "ID OR NUM OR ( OR not OR + OR -", t)
    sync(tokens, lib.SimpleExpressionFollows())
    return
  }
  sign(listing, tokens)
  term(listing, tokens)
  simpleExpressionPrime(listing, tokens)
}

func expressionPrime(listing *list.List, tokens *list.List) {
  t, ok := matchYank(tokens, lib.RELOP)
  if ok {
    simpleExpression(listing, tokens)
    return
  }
  t, ok = match(tokens, lib.ELSE)
  if !ok {
    t, ok = match(tokens, lib.SEMICOLON)
  }
  if !ok {
    t, ok = match(tokens, lib.END)
  }
  if !ok {
    t, ok = match(tokens, lib.THEN)
  }
  if !ok {
    t, ok = match(tokens, lib.DO)
  }
  if !ok {
    t, ok = match(tokens, lib.CLOSE_BRACKET)
  }
  if !ok {
    t, ok = match(tokens, lib.COMMA)
  }
  if !ok {
    t, ok = match(tokens, lib.CLOSE_PAREN)
  }
  if !ok {
    report(listing, "RELOP OR else OR ; OR end OR then OR do OR ] OR )", t)
    sync(tokens, lib.ExpressionPrimeFollows())
    return
  }
  // epsilon production
} 

func expression(listing *list.List, tokens *list.List) {
  t, ok := match(tokens, lib.ID)
  if !ok {
    t, ok = match(tokens, lib.NUM)
  }
  if !ok {
    t, ok = match(tokens, lib.OPEN_PAREN)
  }
  if !ok {
    t, ok = match(tokens, lib.NOT)
  }
  if !ok {
    t, ok = match(tokens, lib.ADDOP)
    if !ok || (t.Type == lib.ADDOP && t.Attr != lib.PLUS && t.Attr != lib.MINUS) {
      report(listing, "ID OR NUM OR ( OR not OR + OR -", t)
      sync(tokens, lib.ExpressionFollows())
      return
    }
  }
  simpleExpression(listing, tokens)
  expressionPrime(listing, tokens)
}

func variablePrime(listing *list.List, tokens *list.List) {
  t, ok := matchYank(tokens, lib.OPEN_BRACKET)
  if ok {
    expression(listing, tokens)
    if t, ok = matchYank(tokens, lib.CLOSE_BRACKET); !ok {
      report(listing, "]", t)
      sync(tokens, lib.VariablePrimeFollows())
      return
    }
    return
  }
  t, ok = match(tokens, lib.ASSIGNOP)
  if !ok {
    report(listing, "[ OR :=", t)
    sync(tokens, lib.VariablePrimeFollows())
    return
  }
  // epsilon production
}

func variable(listing *list.List, tokens *list.List) {
  t, ok := matchYank(tokens, lib.ID)
  if !ok {
    report(listing, "ID", t)
    sync(tokens, lib.VariableFollows())
    return
  }
  variablePrime(listing, tokens)
}

func statement(listing *list.List, tokens *list.List) {
  t, ok := match(tokens, lib.ID)
  if ok {
    variable(listing, tokens)
    if t, ok = matchYank(tokens, lib.ASSIGNOP); !ok {
      report(listing, ":=", t)
      sync(tokens, lib.StatementFollows())
      return
    }
    expression(listing, tokens)
    return
  }
  t, ok = match(tokens, lib.BEGIN)
  if ok {
    compoundStatement(listing, tokens)
    return
  }
  t, ok = matchYank(tokens, lib.IF)
  if ok {
    expression(listing, tokens)
    if t, ok = matchYank(tokens, lib.THEN); !ok {
      report(listing, "then", t)
      sync(tokens, lib.StatementFollows())
      return
    }
    statement(listing, tokens)
    statementPrime(listing, tokens)
    return
  }
  t, ok = matchYank(tokens, lib.WHILE)
  if !ok {
    report(listing, "ID OR begin OR if OR while", t)
    sync(tokens, lib.StatementFollows())
    return
  }
  expression(listing, tokens)
  if t, ok = matchYank(tokens, lib.DO); !ok {
    report(listing, "do", t)
    sync(tokens, lib.StatementFollows())
    return
  }
  statement(listing, tokens)
}

func statementListPrime(listing *list.List, tokens *list.List) {
  t, ok := matchYank(tokens, lib.SEMICOLON)
  if ok {
    statement(listing, tokens)
    statementListPrime(listing, tokens)
    return
  }
  t, ok = match(tokens, lib.END)
  if !ok {
    report(listing, "; OR end", t)
    sync(tokens, lib.StatementListPrimeFollows())
    return
  }
  // epsilon production
}

func statementList(listing *list.List, tokens *list.List) {
  t, ok := match(tokens, lib.ID)
  if !ok {
    t, ok = match(tokens, lib.BEGIN)
  }
  if !ok {
    t, ok = match(tokens, lib.IF)
  }
  if !ok {
    t, ok = match(tokens, lib.WHILE)
  }
  if !ok {
    report(listing, "ID OR begin OR if OR while", t)
    sync(tokens, lib.StatementListFollows())
    return
  }
  statement(listing, tokens)
  statementListPrime(listing, tokens)
}

func optionalStatements(listing *list.List, tokens *list.List) {
  t, ok := match(tokens, lib.ID)
  if !ok {
    t, ok = match(tokens, lib.BEGIN)
  }
  if !ok {
    t, ok = match(tokens, lib.IF)
  }
  if !ok {
    t, ok = match(tokens, lib.WHILE)
  }
  if !ok {
    report(listing, "ID OR begin OR if OR while", t)
    sync(tokens, lib.OptionalStatementsFollows())
    return
  }
  statementList(listing, tokens)
}

func compoundStatementPrime(listing *list.List, tokens *list.List) {
  t, ok := match(tokens, lib.ID)
  if !ok {
    t, ok = match(tokens, lib.BEGIN)
  }
  if !ok {
    t, ok = match(tokens, lib.IF)
  }
  if !ok {
    t, ok = match(tokens, lib.WHILE)
  }
  if ok {
    optionalStatements(listing, tokens)
    if t, ok = matchYank(tokens, lib.END); !ok {
      report(listing, "end", t)
      sync(tokens, lib.CompoundStatementPrimeFollows())
      return
    }
    return
  }
  t, ok = matchYank(tokens, lib.END)
  if !ok {
    report(listing, "ID OR begin OR if OR while OR end", t)
    sync(tokens, lib.CompoundStatementPrimeFollows())
    return
  }
}

func compoundStatement(listing *list.List, tokens *list.List) {
  t, ok := matchYank(tokens, lib.BEGIN)
  if !ok {
    report(listing, "begin", t)
    sync(tokens, lib.CompoundStatementFollows())
    return
  }
  compoundStatementPrime(listing, tokens)
}

func subprogramSubbody(listing *list.List, tokens *list.List) {
  t, ok := match(tokens, lib.FUNCTION)
  if ok {
    subprogramDeclarations(listing, tokens)
    compoundStatement(listing, tokens)
    return
  }
  t, ok = match(tokens, lib.BEGIN)
  if !ok {
    report(listing, "function OR begin", t)
    sync(tokens, lib.SubprogramSubbodyFollows())
    return
  }
  compoundStatement(listing, tokens)
}

func subprogramBody(listing *list.List, tokens *list.List) {
  t, ok := match(tokens, lib.VAR)
  if ok {
    declarations(listing, tokens)
    subprogramSubbody(listing, tokens)
    return
  }
  t, ok = match(tokens, lib.FUNCTION)
  if !ok {
    t, ok = match(tokens, lib.BEGIN)
    if !ok {
      report(listing, "var OR function OR begin", t)
      sync(tokens, lib.SubprogramBodyFollows())
      return
    }
  }
  subprogramSubbody(listing, tokens)
}

func subprogramDeclaration(listing *list.List, tokens *list.List) {
  t, ok := match(tokens, lib.FUNCTION)
  if !ok {
    report(listing, "function", t)
    sync(tokens, lib.SubprogramDeclarationFollows())
    return
  }
  subprogramHead(listing, tokens)
  subprogramBody(listing, tokens)
}

func subprogramDeclarationsPrime(listing *list.List, tokens *list.List) {
  t, ok := match(tokens, lib.FUNCTION)
  if ok {
    subprogramDeclaration(listing, tokens)
    if t, ok = matchYank(tokens, lib.SEMICOLON); !ok {
      report(listing, ";", t)
      sync(tokens, lib.SubprogramDeclarationsPrimeFollows())
      return
    }
    subprogramDeclarationsPrime(listing, tokens)
    return
  }
  if t, ok = match(tokens, lib.BEGIN); !ok {
    report(listing, "function OR begin", t)
    sync(tokens, lib.SubprogramDeclarationsPrimeFollows())
    return
  }
  // epsilon production
}

func subprogramDeclarations(listing *list.List, tokens *list.List) {
  t, ok := match(tokens, lib.FUNCTION)
  if !ok {
    report(listing, "function", t)
    sync(tokens, lib.SubprogramDeclarationsFollows())
    return
  }
  subprogramDeclaration(listing, tokens)
  if t, ok = matchYank(tokens, lib.SEMICOLON); !ok {
    report(listing, ";", t)
    sync(tokens, lib.SubprogramDeclarationsFollows())
    return
  }
  subprogramDeclarationsPrime(listing, tokens)
}

func programSubbody(listing *list.List, tokens *list.List) {
  t, ok := match(tokens, lib.FUNCTION)
  if ok {
    subprogramDeclarations(listing, tokens)
    compoundStatement(listing, tokens)
    if t, ok = matchYank(tokens, lib.PERIOD); !ok {
      report(listing, ".", t)
      sync(tokens, lib.ProgramSubbodyFollows())
      return
    }
    return
  }
  t, ok = match(tokens, lib.BEGIN)
  if !ok {
    report(listing, "function OR begin", t)
    sync(tokens, lib.ProgramSubbodyFollows())
    return
  }
  compoundStatement(listing, tokens)
  if t, ok = matchYank(tokens, lib.PERIOD); !ok {
    report(listing, ".", t)
    sync(tokens, lib.ProgramSubbodyFollows())
    return
  }
}

func programBody(listing *list.List, tokens *list.List) {
  t, ok := match(tokens, lib.VAR)
  if ok {
    declarations(listing, tokens)
    programSubbody(listing, tokens)
    return
  }
  t, ok = match(tokens, lib.FUNCTION)
  if !ok {
    t, ok = match(tokens, lib.BEGIN)
    if !ok {
      report(listing, "var OR function OR begin", t)
      sync(tokens, lib.ProgramBodyFollows())
    }
  }
  programSubbody(listing, tokens)
}

func program(listing *list.List, tokens *list.List) {
  t, ok := matchYank(tokens, lib.PROGRAM)
  if !ok {
    report(listing, "program", t)
    sync(tokens, lib.ProgramFollows())
    return
  }
  if t, ok = matchYank(tokens, lib.ID); !ok {
    report(listing, "ID", t)
    sync(tokens, lib.ProgramFollows())
    return
  }
  if t, ok = matchYank(tokens, lib.OPEN_PAREN); !ok {
    report(listing, "(", t)
    sync(tokens, lib.ProgramFollows())
    return
  }
  identifierList(listing, tokens)
  if t, ok = matchYank(tokens, lib.CLOSE_PAREN); !ok {
    report(listing, ")", t)
    sync(tokens, lib.ProgramFollows())
    return
  }
  if t, ok = matchYank(tokens, lib.SEMICOLON); !ok {
    report(listing, ";", t)
    sync(tokens, lib.ProgramFollows())
    return
  }
  programBody(listing, tokens)
  if t, ok = matchYank(tokens, lib.EOF); !ok {
    report(listing, "EOF", t)
    sync(tokens, lib.ProgramFollows())
    return
  }
}

func Parse(listing *list.List, tokens *list.List, symbols map[string]*lib.Token) {
  program(listing, tokens)
}