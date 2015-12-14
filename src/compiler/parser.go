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
  if !ok {
    // epsilon production
    return
  }
  if t, ok = matchYank(tokens, lib.ID); !ok {
    report(listing, "ID", t)
    sync(tokens, lib.IdentifierListPrimeFollows())
    return
  }
  identifierListPrime(listing, tokens)
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
  if !ok {
    if t, ok = match(tokens, lib.FUNCTION); ok {
      // epsilon production
      return
    }
    if t, ok = match(tokens, lib.BEGIN); ok {
      // epsilon production
      return
    }
    report(listing, "var OR function OR begin", t)
    sync(tokens, lib.DeclarationsPrimeFollows())
    return
  }
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
  // parameterList(listing, tokens)
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

func subprogramDeclaration(listing *list.List, tokens *list.List) {
  t, ok := match(tokens, lib.FUNCTION)
  if !ok {
    report(listing, "function", t)
    sync(tokens, lib.SubprogramDeclarationFollows())
    return
  }
  subprogramHead(listing, tokens)
  // subprogramBody(listing, tokens)
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
    // compoundStatement(listing, tokens)
    if t, ok = matchYank(tokens, lib.PERIOD); !ok {
      report(listing, ".", t)
      sync(tokens, lib.ProgramSubbodyFollows())
      return
    }
  }
  t, ok = match(tokens, lib.BEGIN)
  if !ok {
    report(listing, "function OR begin", t)
    sync(tokens, lib.ProgramSubbodyFollows())
    return
  }
  // compoundStatement(listing, tokens)
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
}

func Parse(listing *list.List, tokens *list.List, symbols map[string]*lib.Token) {
  program(listing, tokens)
}