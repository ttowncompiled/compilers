package compiler

import (
  "container/list"
  "fmt"
  "lib"
  "strconv"
)

func checkAddGreenNode(listing *list.List, t lib.Token, stack *list.List, lex string, ttype int) {
  for stackNode := stack.Front(); stackNode != nil; stackNode = stackNode.Next() {
    if stackNode.Value.(lib.GreenNode).Lexeme == lex {
      e := listing.Front()
      for i := 1; i < t.LineNumber; i++ {
        e = e.Next()
      }
      line := e.Value.(lib.Line)
      line.Errors.PushBack(lib.Error{"SCOPE_ERR: " + lex + " HAS ALREADY BEEN DECLARED WITHIN THE CURRENT SCOPE", &t})
      return
    }
    for childNode := stackNode.Value.(lib.GreenNode).Children.Front(); childNode != nil; childNode = childNode.Next() {
      if childNode.Value.(lib.Node).Lex() == lex {
        e := listing.Front()
        for i := 1; i < t.LineNumber; i++ {
          e = e.Next()
        }
        line := e.Value.(lib.Line)
        line.Errors.PushBack(lib.Error{"SCOPE_ERR: " + lex + " HAS ALREADY BEEN DECLARED WITHIN THE CURRENT SCOPE", &t})
        return
      }
    }
  }
  greenNode := lib.GreenNode{lex, ttype, list.New()}
  peek := stack.Front().Value.(lib.GreenNode)
  peek.Children.PushBack(greenNode)
  stack.PushFront(greenNode)
}

func checkAddBlueNode(listing *list.List, t lib.Token, stack *list.List, lex string, ttype int) {
  peek := stack.Front()
  if peek.Value.(lib.GreenNode).Lexeme == lex {
    e := listing.Front()
    for i := 1; i < t.LineNumber; i++ {
      e = e.Next()
    }
    line := e.Value.(lib.Line)
    line.Errors.PushBack(lib.Error{"SCOPE_ERR: " + lex + " HAS ALREADY BEEN DECLARED WITHIN THE CURRENT SCOPE", &t})
    return
  }
  for childNode := peek.Value.(lib.GreenNode).Children.Front(); childNode != nil; childNode.Next() {
    if childNode.Value.(lib.Node).Lex() == lex {
      e := listing.Front()
      for i := 1; i < t.LineNumber; i++ {
        e = e.Next()
      }
      line := e.Value.(lib.Line)
      line.Errors.PushBack(lib.Error{"SCOPE_ERR: " + lex + " HAS ALREADY BEEN DECLARED WITHIN THE CURRENT SCOPE", &t})
      return
    }
  }
  blueNode := lib.BlueNode{lex, ttype}
  peek.Value.(lib.GreenNode).Children.PushBack(blueNode)
}

func popGreenNode(symbols map[string]*lib.Symbol, stack *list.List) {
  if stack.Len() == 0 {
    return
  }
  top := stack.Front()
  stack.Remove(top)
  greenNode := top.Value.(lib.GreenNode)
  symbols[greenNode.Lexeme].Decoration = *(symbols[greenNode.Lexeme].Decoration.PrevTypeD())
  for child := greenNode.Children.Front(); child != nil; child = child.Next() {
    childNode := child.Value.(lib.Node)
    symbols[childNode.Lex()].Decoration = *(symbols[childNode.Lex()].Decoration.PrevTypeD())
  }
}

func addType(lex string, ttype lib.TypeD, symbols map[string]*lib.Symbol) {
  if ttype.TypeD() == lib.ARRAY {
    decoration := ttype.(lib.ArrayD)
    decoration.Prev = &(symbols[lex].Decoration)
    symbols[lex].Decoration = decoration
    fmt.Println(lex, lib.Annotate(symbols[lex].Decoration.TypeD()), "of", lib.Annotate(symbols[lex].Decoration.(lib.ArrayD).Val.TypeD()))
  } else if ttype.TypeD() == lib.FUNCTION { 
    decoration := ttype.(lib.FunctionD)
    decoration.Prev = &(symbols[lex].Decoration)
    symbols[lex].Decoration = decoration
    fmt.Println(lex, lib.Annotate(symbols[lex].Decoration.TypeD()), "to", lib.Annotate(symbols[lex].Decoration.(lib.FunctionD).Return.TypeD()))
  } else {
    decoration := ttype.(lib.Decoration)
    decoration.Prev = &(symbols[lex].Decoration)
    symbols[lex].Decoration = decoration
    fmt.Println(lex, lib.Annotate(symbols[lex].Decoration.TypeD()))
  }
}

func modifyType(lex string, ttype lib.TypeD, symbols map[string]*lib.Symbol) {
  if ttype.TypeD() == lib.ARRAY {
    decoration := ttype.(lib.ArrayD)
    decoration.Prev = symbols[lex].Decoration.PrevTypeD()
    symbols[lex].Decoration = decoration
    fmt.Println(lex, lib.Annotate(symbols[lex].Decoration.TypeD()), "of", lib.Annotate(symbols[lex].Decoration.(lib.ArrayD).Val.TypeD()))
  } else if ttype.TypeD() == lib.FUNCTION { 
    decoration := ttype.(lib.FunctionD)
    decoration.Prev = symbols[lex].Decoration.PrevTypeD()
    symbols[lex].Decoration = decoration
    fmt.Println(lex, lib.Annotate(symbols[lex].Decoration.TypeD()), "to", lib.Annotate(symbols[lex].Decoration.(lib.FunctionD).Return.TypeD()))
  } else {
    decoration := ttype.(lib.ArrayD)
    decoration.Prev = symbols[lex].Decoration.PrevTypeD()
    symbols[lex].Decoration = decoration
    fmt.Println(lex, lib.Annotate(symbols[lex].Decoration.TypeD()))
  }
}

func getType(lex string, symbols map[string]*lib.Symbol) lib.TypeD {
  fmt.Println("get", lex, lib.Annotate(symbols[lex].Decoration.TypeD()))
  return symbols[lex].Decoration
}

func checkTypeAndReport(listing *list.List, t lib.Token, left int, right int) bool {
  fmt.Println("check", lib.Annotate(left), lib.Annotate(right))
  if left == lib.ERR || right == lib.ERR {
    return false
  }
  if left != right {
    e := listing.Front()
    for i := 1; i < t.LineNumber; i++ {
      e = e.Next()
    }
    line := e.Value.(lib.Line)
    line.Errors.PushBack(lib.Error{"TYPE_ERR: EXPECTED TYPE: " + lib.Annotate(left) + " RECEIVED TYPE: " + lib.Annotate(right), &t})
    return false
  }
  return true
}

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

func identifierListPrime(listing *list.List, tokens *list.List, symbols map[string]*lib.Symbol, stack *list.List) {
  t, ok := matchYank(tokens, lib.COMMA)
  if ok {
    if t, ok = matchYank(tokens, lib.ID); !ok {
      report(listing, "ID", t)
      sync(tokens, lib.IdentifierListPrimeFollows())
      return
    }
    addType(t.Lexeme, lib.Decoration{lib.PARG, nil}, symbols)
    checkAddBlueNode(listing, t, stack, t.Lexeme, lib.PARG)
    identifierListPrime(listing, tokens, symbols, stack)
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

func identifierList(listing *list.List, tokens *list.List, symbols map[string]*lib.Symbol, stack *list.List) {
  t, ok := matchYank(tokens, lib.ID)
  if !ok {
    report(listing, "ID", t)
    sync(tokens, lib.IdentifierListFollows())
    return
  }
  addType(t.Lexeme, lib.Decoration{lib.PARG, nil}, symbols)
  checkAddBlueNode(listing, t, stack, t.Lexeme, lib.PARG)
  identifierListPrime(listing, tokens, symbols, stack)
}

func standardType(listing *list.List, tokens *list.List, symbols map[string]*lib.Symbol) lib.TypeD {
  t, ok := matchYank(tokens, lib.INTEGER)
  if ok {
    return lib.Decoration{lib.INTEGER, nil}
  }
  t, ok = matchYank(tokens, lib.REAL)
  if !ok {
    report(listing, "integer OR real", t)
    sync(tokens, lib.StandardTypeFollows())
    return lib.Decoration{lib.ERR, nil}
  }
  return lib.Decoration{lib.REAL, nil}
}

func type_(listing *list.List, tokens *list.List, symbols map[string]*lib.Symbol) lib.TypeD {
  t, ok := match(tokens, lib.INTEGER)
  if !ok {
    t, ok = match(tokens, lib.REAL)
  }
  if ok {
    return standardType(listing, tokens, symbols)
  }
  t, ok = matchYank(tokens, lib.ARRAY)
  if !ok {
    report(listing, "integer OR real OR array", t)
    sync(tokens, lib.TypeFollows())
    return lib.Decoration{lib.ERR, nil}
  }
  if t, ok = matchYank(tokens, lib.OPEN_BRACKET); !ok {
    report(listing, "[", t)
    sync(tokens, lib.TypeFollows())
    return lib.Decoration{lib.ERR, nil}
  }
  if t, ok = matchYank(tokens, lib.NUM); !ok {
    report(listing, "NUM", t)
    sync(tokens, lib.TypeFollows())
    return lib.Decoration{lib.ERR, nil}
  }
  num1, _ := strconv.ParseInt(t.Lexeme, 64, 10)
  if t, ok = matchYank(tokens, lib.RANGE); !ok {
    report(listing, "..", t)
    sync(tokens, lib.TypeFollows())
    return lib.Decoration{lib.ERR, nil}
  }
  if t, ok = matchYank(tokens, lib.NUM); !ok {
    report(listing, "NUM", t)
    sync(tokens, lib.TypeFollows())
    return lib.Decoration{lib.ERR, nil}
  }
  num2, _ := strconv.ParseInt(t.Lexeme, 64, 10)
  if t, ok = matchYank(tokens, lib.CLOSE_BRACKET); !ok {
    report(listing, "]", t)
    sync(tokens, lib.TypeFollows())
    return lib.Decoration{lib.ERR, nil}
  }
  if t, ok = matchYank(tokens, lib.OF); !ok {
    report(listing, "of", t)
    sync(tokens, lib.TypeFollows())
    return lib.Decoration{lib.ERR, nil}
  }
  valType := standardType(listing, tokens, symbols)
  return lib.ArrayD{int(num2 - num1), valType, nil}
}

func declarationsPrime(listing *list.List, tokens *list.List, symbols map[string]*lib.Symbol, stack *list.List, addresses *list.List, loc int) {
  t, ok := matchYank(tokens, lib.VAR)
  if ok {
    if t, ok = matchYank(tokens, lib.ID); !ok {
      report(listing, "ID", t)
      sync(tokens, lib.DeclarationsPrimeFollows())
      return
    }
    id := t
    checkAddBlueNode(listing, id, stack, id.Lexeme, lib.VAR)
    if t, ok = matchYank(tokens, lib.COLON); !ok {
      report(listing, ":", t)
      sync(tokens, lib.DeclarationsPrimeFollows())
      return
    }
    ttype := type_(listing, tokens, symbols)
    addType(id.Lexeme, ttype, symbols)
    addresses.PushBack(lib.Address{id.Lexeme, loc})
    loc += computeMemory(ttype)
    if t, ok = matchYank(tokens, lib.SEMICOLON); !ok {
      report(listing, ";", t)
      sync(tokens, lib.DeclarationsPrimeFollows())
      return
    }
    declarationsPrime(listing, tokens, symbols, stack, addresses, loc)
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

func declarations(listing *list.List, tokens *list.List, symbols map[string]*lib.Symbol, stack *list.List, addresses *list.List, loc int) {
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
  id := t
  checkAddBlueNode(listing, id, stack, id.Lexeme, lib.VAR)
  if t, ok = matchYank(tokens, lib.COLON); !ok {
    report(listing, ":", t)
    sync(tokens, lib.DeclarationsFollows())
    return
  }
  ttype := type_(listing, tokens, symbols)
  addType(id.Lexeme, ttype, symbols)
  addresses.PushBack(lib.Address{id.Lexeme, loc})
  loc += computeMemory(ttype)
  if t, ok = matchYank(tokens, lib.SEMICOLON); !ok {
    report(listing, ";", t)
    sync(tokens, lib.DeclarationsFollows())
    return
  }
  declarationsPrime(listing, tokens, symbols, stack, addresses, loc)
}

func parameterListPrime(listing *list.List, tokens *list.List, symbols map[string]*lib.Symbol, stack *list.List) *list.List {
  t, ok := matchYank(tokens, lib.SEMICOLON)
  if ok {
    if t, ok = matchYank(tokens, lib.ID); !ok {
      report(listing, "ID", t)
      sync(tokens, lib.ParameterListPrimeFollows())
      l := list.New()
      l.PushFront(lib.Decoration{lib.ERR, nil})
      return l
    }
    id := t
    checkAddBlueNode(listing, id, stack, id.Lexeme, lib.FARG)
    if t, ok = matchYank(tokens, lib.COLON); !ok {
      report(listing, ":", t)
      sync(tokens, lib.ParameterListPrimeFollows())
      l := list.New()
      l.PushFront(lib.Decoration{lib.ERR, nil})
      return l
    }
    ttype := type_(listing, tokens, symbols)
    addType(id.Lexeme, ttype, symbols)
    plist := parameterListPrime(listing, tokens, symbols, stack)
    plist.PushFront(ttype)
    return plist
  }
  if t, ok = match(tokens, lib.CLOSE_PAREN); !ok {
    report(listing, "; OR )", t)
    sync(tokens, lib.ParameterListPrimeFollows())
    l := list.New()
    l.PushFront(lib.Decoration{lib.ERR, nil})
    return l
  }
  // epsilon production
  l := list.New()
  l.PushFront(lib.Decoration{lib.VOID, nil})
  return l
}

func parameterList(listing *list.List, tokens *list.List, symbols map[string]*lib.Symbol, stack *list.List) *list.List {
  t, ok := matchYank(tokens, lib.ID)
  if !ok {
    report(listing, "ID", t)
    sync(tokens, lib.ParameterListFollows())
    l := list.New()
    l.PushFront(lib.Decoration{lib.ERR, nil})
    return l
  }
  id := t
  checkAddBlueNode(listing, id, stack, id.Lexeme, lib.FARG)
  if t, ok = matchYank(tokens, lib.COLON); !ok {
    report(listing, ":", t)
    sync(tokens, lib.ParameterListFollows())
    l := list.New()
    l.PushFront(lib.Decoration{lib.ERR, nil})
    return l
  }
  ttype := type_(listing, tokens, symbols)
  addType(id.Lexeme, ttype, symbols)
  plist := parameterListPrime(listing, tokens, symbols, stack)
  plist.PushFront(ttype)
  return plist
}

func arguments(listing *list.List, tokens *list.List, symbols map[string]*lib.Symbol, stack *list.List) *list.List {
  t, ok := matchYank(tokens, lib.OPEN_PAREN)
  if !ok {
    report(listing, "(", t)
    sync(tokens, lib.ArgumentsFollows())
    l := list.New()
    l.PushFront(lib.Decoration{lib.ERR, nil})
    return l
  }
  plist := parameterList(listing, tokens, symbols, stack)
  if t, ok = matchYank(tokens, lib.CLOSE_PAREN); !ok {
    report(listing, ")", t)
    sync(tokens, lib.ArgumentsFollows())
    l := list.New()
    l.PushFront(lib.Decoration{lib.ERR, nil})
    return l
  }
  return plist
}

func subprogramHeadPrime(listing *list.List, tokens *list.List, symbols map[string]*lib.Symbol, stack *list.List) lib.TypeD {
  t, ok := match(tokens, lib.OPEN_PAREN)
  if ok {
    plist := arguments(listing, tokens, symbols, stack)
    if t, ok = matchYank(tokens, lib.COLON); !ok {
      report(listing, ":", t)
      sync(tokens, lib.SubprogramHeadPrimeFollows())
      return lib.Decoration{lib.ERR, nil}
    }
    ttype := standardType(listing, tokens, symbols)
    if t, ok = matchYank(tokens, lib.SEMICOLON); !ok {
      report(listing, ";", t)
      sync(tokens, lib.SubprogramHeadPrimeFollows())
      return lib.Decoration{lib.ERR, nil}
    }
    return lib.FunctionD{plist, ttype, nil}
  }
  t, ok = matchYank(tokens, lib.COLON)
  if !ok {
    report(listing, "( OR :", t)
    sync(tokens, lib.SubprogramHeadPrimeFollows())
    return lib.Decoration{lib.ERR, nil}
  }
  ttype := standardType(listing, tokens, symbols)
  if t, ok = matchYank(tokens, lib.SEMICOLON); !ok {
    report(listing, ";", t)
    sync(tokens, lib.SubprogramHeadPrimeFollows())
    return lib.Decoration{lib.ERR, nil}
  }
  l := list.New()
  l.PushFront(lib.Decoration{lib.VOID, nil})
  return lib.FunctionD{l, ttype, nil}
}

func subprogramHead(listing *list.List, tokens *list.List, symbols map[string]*lib.Symbol, stack *list.List) {
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
  l := list.New()
  l.PushFront(lib.Decoration{lib.VOID, nil})
  addType(t.Lexeme, lib.FunctionD{l, lib.Decoration{lib.VOID, nil}, nil}, symbols)
  checkAddGreenNode(listing, t, stack, t.Lexeme, lib.FUNCTION)
  ttype := subprogramHeadPrime(listing, tokens, symbols, stack)
  fmt.Println(t.Lexeme)
  modifyType(t.Lexeme, ttype, symbols)
}

func statementPrime(listing *list.List, tokens *list.List, symbols map[string]*lib.Symbol) lib.TypeD {
  t, ok := matchYank(tokens, lib.ELSE)
  if ok {
    return statement(listing, tokens, symbols)
  }
  t, ok = match(tokens, lib.SEMICOLON)
  if !ok {
    t, ok = match(tokens, lib.END)
    if !ok {
      report(listing, "else OR ; OR end", t)
      sync(tokens, lib.StatementPrimeFollows())
      return lib.Decoration{lib.ERR, nil}
    }
  }
  // epsilon production
  return lib.Decoration{lib.VOID, nil}
}

func expressionListPrime(listing *list.List, tokens *list.List, symbols map[string]*lib.Symbol, params *list.List, idx int) *list.List {
  t, ok := matchYank(tokens, lib.COMMA)
  if ok {
    etype := expression(listing, tokens, symbols)
    if idx >= params.Len() {
      // err too many arguments
      e := listing.Front()
      for i := 1; i < t.LineNumber; i++ {
        e = e.Next()
      }
      line := e.Value.(lib.Line)
      line.Errors.PushBack(lib.Error{"TYPE_ERR: TOO MANY ARGUMENTS", &t})
      l := list.New()
      l.PushFront(lib.Decoration{lib.ERR, nil})
      return expressionListPrime(listing, tokens, symbols, l, 0)
    }
    e := params.Front()
    for i := 0; i < idx; i++ {
      e = e.Next()
    }
    inType := e.Value.(lib.TypeD)
    if !checkTypeAndReport(listing, t, inType.TypeD(), etype.TypeD()) {
      l := list.New()
      l.PushFront(lib.Decoration{lib.ERR, nil})
      return expressionListPrime(listing, tokens, symbols, l, 0)
    }
    return expressionListPrime(listing, tokens, symbols, params, idx+1)
  }
  t, ok = match(tokens, lib.CLOSE_PAREN)
  if !ok {
    report(listing, ", OR )", t)
    sync(tokens, lib.ExpressionListPrimeFollows())
    l := list.New()
    l.PushFront(lib.Decoration{lib.ERR, nil})
    return l
  }
  // epsilon production
  if idx+1 < params.Len() && idx != 0 {
    // err not enough arguments
    e := listing.Front()
    for i := 1; i < t.LineNumber; i++ {
      e = e.Next()
    }
    line := e.Value.(lib.Line)
    line.Errors.PushBack(lib.Error{"TYPE_ERR: NOT ENOUGH ARGUMENTS", &t})
    l := list.New()
    l.PushFront(lib.Decoration{lib.ERR, nil})
    return l
  }
  return params
}

func expressionList(listing *list.List, tokens *list.List, symbols map[string]*lib.Symbol, params *list.List, idx int) *list.List {
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
    l := list.New()
    l.PushFront(lib.Decoration{lib.ERR, nil})
    return l
  }
  etype := expression(listing, tokens, symbols)
  if idx >= params.Len() {
    // err too many arguments
    e := listing.Front()
    for i := 1; i < t.LineNumber; i++ {
      e = e.Next()
    }
    line := e.Value.(lib.Line)
    line.Errors.PushBack(lib.Error{"TYPE_ERR: TOO MANY ARGUMENTS", &t})
    l := list.New()
    l.PushFront(lib.Decoration{lib.ERR, nil})
    return expressionListPrime(listing, tokens, symbols, l, 0)
  }
  e := params.Front()
  for i := 0; i < idx; i++ {
    e = e.Next()
  }
  inType := e.Value.(lib.TypeD)
  if !checkTypeAndReport(listing, t, inType.TypeD(), etype.TypeD()) {
    l := list.New()
    l.PushFront(lib.Decoration{lib.ERR, nil})
    return expressionListPrime(listing, tokens, symbols, l, 0)
  }
  return expressionListPrime(listing, tokens, symbols, params, idx+1)
}

func factorPrime(listing *list.List, tokens *list.List, symbols map[string]*lib.Symbol, in lib.TypeD) lib.TypeD {
  t, ok := matchYank(tokens, lib.OPEN_PAREN)
  if ok {
    if !checkTypeAndReport(listing, t, in.TypeD(), lib.FUNCTION) {
      l := list.New()
      l.PushFront(lib.Decoration{lib.ERR, nil})
      expressionList(listing, tokens, symbols, l, 0)
      if t, ok = matchYank(tokens, lib.CLOSE_PAREN); !ok {
        report(listing, ")", t)
        sync(tokens, lib.FactorPrimeFollows())
        return lib.Decoration{lib.ERR, nil}
      }
      return lib.Decoration{lib.ERR, nil}
    } else {
      expressionList(listing, tokens, symbols, in.(lib.FunctionD).Params, 0)
      if t, ok = matchYank(tokens, lib.CLOSE_PAREN); !ok {
        report(listing, ")", t)
        sync(tokens, lib.FactorPrimeFollows())
        return lib.Decoration{lib.ERR, nil}
      }
      return in.(lib.FunctionD).Return
    }
  }
  t, ok = matchYank(tokens, lib.OPEN_BRACKET)
  if ok {
    etype := expression(listing, tokens, symbols)
    if t, ok = matchYank(tokens, lib.CLOSE_BRACKET); !ok {
      report(listing, "]", t)
      sync(tokens, lib.FactorPrimeFollows())
      return lib.Decoration{lib.ERR, nil}
    }
    flag := false
    if !checkTypeAndReport(listing, t, etype.TypeD(), lib.INTEGER) {
      flag = true
    }
    if !checkTypeAndReport(listing, t, in.TypeD(), lib.ARRAY) {
      flag = true
    }
    if flag {
      return lib.Decoration{lib.ERR, nil}
    }
    return in.(lib.ArrayD).Val
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
    return lib.Decoration{lib.ERR, nil}
  }
  // epsilon production
  return in
}

func factor(listing *list.List, tokens *list.List, symbols map[string]*lib.Symbol) lib.TypeD {
  t, ok := matchYank(tokens, lib.ID)
  if ok {
    idType := getType(t.Lexeme, symbols)
    return factorPrime(listing, tokens, symbols, idType)
  }
  t, ok = matchYank(tokens, lib.NUM)
  if ok {
    if t.Attr == lib.INT || t.Attr == lib.INTEGER {
      return lib.Decoration{lib.INTEGER, nil}
    } else if t.Attr == lib.REAL || t.Attr == lib.LONG_REAL {
      return lib.Decoration{lib.REAL, nil}
    }
    return lib.Decoration{lib.ERR, nil}
  }
  t, ok = matchYank(tokens, lib.OPEN_PAREN)
  if ok {
    etype := expression(listing, tokens, symbols)
    if t, ok = matchYank(tokens, lib.CLOSE_PAREN); !ok {
      report(listing, ")", t)
      sync(tokens, lib.FactorFollows())
      return lib.Decoration{lib.ERR, nil}
    }
    return etype
  }
  t, ok = matchYank(tokens, lib.NOT)
  if !ok {
    report(listing, "ID OR NUM OR ( OR not", t)
    sync(tokens, lib.FactorFollows())
    return lib.Decoration{lib.ERR, nil}
  }
  ftype := factor(listing, tokens, symbols)
  if !checkTypeAndReport(listing, t, ftype.TypeD(), lib.BOOLEAN) {
    return lib.Decoration{lib.ERR, nil}
  }
  return lib.Decoration{lib.BOOLEAN, nil}
}

func termPrime(listing *list.List, tokens *list.List, symbols map[string]*lib.Symbol, in lib.TypeD) lib.TypeD {
  t, ok := matchYank(tokens, lib.MULOP)
  if ok {
    ftype := factor(listing, tokens, symbols)
    if !checkTypeAndReport(listing, t, in.TypeD(), ftype.TypeD()) {
      return termPrime(listing, tokens, symbols, lib.Decoration{lib.ERR, nil})
    }
    return termPrime(listing, tokens, symbols, ftype)
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
    return lib.Decoration{lib.ERR, nil}
  }
  // epsilon production
  return in
}

func term(listing *list.List, tokens *list.List, symbols map[string]*lib.Symbol) lib.TypeD {
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
    return lib.Decoration{lib.ERR, nil}
  }
  ftype := factor(listing, tokens, symbols)
  return termPrime(listing, tokens, symbols, ftype)
}

func simpleExpressionPrime(listing *list.List, tokens *list.List, symbols map[string]*lib.Symbol, in lib.TypeD) lib.TypeD {
  t, ok := matchYank(tokens, lib.ADDOP)
  if ok {
    ttype := term(listing, tokens, symbols)
    if !checkTypeAndReport(listing, t, in.TypeD(), ttype.TypeD()) {
      return simpleExpressionPrime(listing, tokens, symbols, lib.Decoration{lib.ERR, nil})
    }
    return simpleExpressionPrime(listing, tokens, symbols, ttype)
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
    return lib.Decoration{lib.ERR, nil}
  }
  // epsilon production
  return in
}

func sign(listing *list.List, tokens *list.List, symbols map[string]*lib.Symbol) {
  t, ok := matchYank(tokens, lib.ADDOP)
  if !ok || (t.Attr != lib.PLUS && t.Attr != lib.MINUS) {
    report(listing, "+ OR -", t)
    sync(tokens, lib.SignFollows())
    return
  }
}

func simpleExpression(listing *list.List, tokens *list.List, symbols map[string]*lib.Symbol) lib.TypeD {
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
    ttype := term(listing, tokens, symbols)
    return simpleExpressionPrime(listing, tokens, symbols, ttype)
  }
  t, ok = match(tokens, lib.ADDOP)
  if !ok || (t.Type == lib.ADDOP && t.Attr != lib.PLUS && t.Attr != lib.MINUS) {
    report(listing, "ID OR NUM OR ( OR not OR + OR -", t)
    sync(tokens, lib.SimpleExpressionFollows())
    return lib.Decoration{lib.ERR, nil}
  }
  sign(listing, tokens, symbols)
  ttype := term(listing, tokens, symbols)
  return simpleExpressionPrime(listing, tokens, symbols, ttype)
}

func expressionPrime(listing *list.List, tokens *list.List, symbols map[string]*lib.Symbol, in lib.TypeD) lib.TypeD {
  t, ok := matchYank(tokens, lib.RELOP)
  if ok {
    stype := simpleExpression(listing, tokens, symbols)
    if !checkTypeAndReport(listing, t, in.TypeD(), stype.TypeD()) {
      return lib.Decoration{lib.ERR, nil}
    }
    return lib.Decoration{lib.BOOLEAN, nil}
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
    return lib.Decoration{lib.ERR, nil}
  }
  // epsilon production
  return in
} 

func expression(listing *list.List, tokens *list.List, symbols map[string]*lib.Symbol) lib.TypeD {
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
      return lib.Decoration{lib.ERR, nil}
    }
  }
  stype := simpleExpression(listing, tokens, symbols)
  return expressionPrime(listing, tokens, symbols, stype)
}

func variablePrime(listing *list.List, tokens *list.List, symbols map[string]*lib.Symbol, in lib.TypeD) lib.TypeD {
  t, ok := matchYank(tokens, lib.OPEN_BRACKET)
  if ok {
    etype := expression(listing, tokens, symbols)
    if t, ok = matchYank(tokens, lib.CLOSE_BRACKET); !ok {
      report(listing, "]", t)
      sync(tokens, lib.VariablePrimeFollows())
      return lib.Decoration{lib.ERR, nil}
    }
    flag := false
    if checkTypeAndReport(listing, t, etype.TypeD(), lib.INTEGER) {
      flag = true
    }
    if !checkTypeAndReport(listing, t, in.TypeD(), lib.ARRAY) {
      flag = true
    }
    if flag {
      return lib.Decoration{lib.ERR, nil}
    }
    return in.(lib.ArrayD).Val
  }
  t, ok = match(tokens, lib.ASSIGNOP)
  if !ok {
    report(listing, "[ OR :=", t)
    sync(tokens, lib.VariablePrimeFollows())
    return lib.Decoration{lib.ERR, nil}
  }
  // epsilon production
  return in
}

func variable(listing *list.List, tokens *list.List, symbols map[string]*lib.Symbol) lib.TypeD {
  t, ok := matchYank(tokens, lib.ID)
  if !ok {
    report(listing, "ID", t)
    sync(tokens, lib.VariableFollows())
    return lib.Decoration{lib.ERR, nil}
  }
  idType := getType(t.Lexeme, symbols)
  return variablePrime(listing, tokens, symbols, idType)
}

func statement(listing *list.List, tokens *list.List, symbols map[string]*lib.Symbol) lib.TypeD {
  t, ok := match(tokens, lib.ID)
  if ok {
    vtype := variable(listing, tokens, symbols)
    if t, ok = matchYank(tokens, lib.ASSIGNOP); !ok {
      report(listing, ":=", t)
      sync(tokens, lib.StatementFollows())
      return lib.Decoration{lib.ERR, nil}
    }
    etype := expression(listing, tokens, symbols)
    if !checkTypeAndReport(listing, t, vtype.TypeD(), etype.TypeD()) {
      return lib.Decoration{lib.ERR, nil}
    }
    return lib.Decoration{lib.VOID, nil}
  }
  t, ok = match(tokens, lib.BEGIN)
  if ok {
    return compoundStatement(listing, tokens, symbols)
  }
  t, ok = matchYank(tokens, lib.IF)
  if ok {
    etype := expression(listing, tokens, symbols)
    if t, ok = matchYank(tokens, lib.THEN); !ok {
      report(listing, "then", t)
      sync(tokens, lib.StatementFollows())
      return lib.Decoration{lib.ERR, nil}
    }
    stype1 := statement(listing, tokens, symbols)
    stype2 := statementPrime(listing, tokens, symbols)
    if !checkTypeAndReport(listing, t, etype.TypeD(), lib.BOOLEAN) {
      return lib.Decoration{lib.ERR, nil}
    }
    if !checkTypeAndReport(listing, t, stype1.TypeD(), stype2.TypeD()) {
      return lib.Decoration{lib.ERR, nil}
    } 
    return lib.Decoration{lib.VOID, nil}
  }
  t, ok = matchYank(tokens, lib.WHILE)
  if !ok {
    report(listing, "ID OR begin OR if OR while", t)
    sync(tokens, lib.StatementFollows())
    return lib.Decoration{lib.ERR, nil}
  }
  etype := expression(listing, tokens, symbols)
  if t, ok = matchYank(tokens, lib.DO); !ok {
    report(listing, "do", t)
    sync(tokens, lib.StatementFollows())
    return lib.Decoration{lib.ERR, nil}
  }
  stype := statement(listing, tokens, symbols)
  if !checkTypeAndReport(listing, t, etype.TypeD(), lib.BOOLEAN) {
    return lib.Decoration{lib.ERR, nil}
  }
  return stype
}

func statementListPrime(listing *list.List, tokens *list.List, symbols map[string]*lib.Symbol) lib.TypeD {
  t, ok := matchYank(tokens, lib.SEMICOLON)
  if ok {
    stype := statement(listing, tokens, symbols)
    sltype := statementListPrime(listing, tokens, symbols)
    if !checkTypeAndReport(listing, t, stype.TypeD(), sltype.TypeD()) {
      return lib.Decoration{lib.ERR, nil}
    }
    return lib.Decoration{lib.VOID, nil}
  }
  t, ok = match(tokens, lib.END)
  if !ok {
    report(listing, "; OR end", t)
    sync(tokens, lib.StatementListPrimeFollows())
    return lib.Decoration{lib.ERR, nil}
  }
  // epsilon production
  return lib.Decoration{lib.VOID, nil}
}

func statementList(listing *list.List, tokens *list.List, symbols map[string]*lib.Symbol) lib.TypeD {
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
    return lib.Decoration{lib.ERR, nil}
  }
  stype := statement(listing, tokens, symbols)
  sltype := statementListPrime(listing, tokens, symbols)
  if !checkTypeAndReport(listing, t, stype.TypeD(), sltype.TypeD()) {
    return lib.Decoration{lib.ERR, nil}
  }
  return lib.Decoration{lib.VOID, nil}
}

func optionalStatements(listing *list.List, tokens *list.List, symbols map[string]*lib.Symbol) lib.TypeD {
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
    return lib.Decoration{lib.ERR, nil}
  }
  return statementList(listing, tokens, symbols)
}

func compoundStatementPrime(listing *list.List, tokens *list.List, symbols map[string]*lib.Symbol) lib.TypeD {
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
    otype := optionalStatements(listing, tokens, symbols)
    if t, ok = matchYank(tokens, lib.END); !ok {
      report(listing, "end", t)
      sync(tokens, lib.CompoundStatementPrimeFollows())
      return lib.Decoration{lib.ERR, nil}
    }
    return otype
  }
  t, ok = matchYank(tokens, lib.END)
  if !ok {
    report(listing, "ID OR begin OR if OR while OR end", t)
    sync(tokens, lib.CompoundStatementPrimeFollows())
    return lib.Decoration{lib.ERR, nil}
  }
  return lib.Decoration{lib.VOID, nil}
}

func compoundStatement(listing *list.List, tokens *list.List, symbols map[string]*lib.Symbol) lib.TypeD {
  t, ok := matchYank(tokens, lib.BEGIN)
  if !ok {
    report(listing, "begin", t)
    sync(tokens, lib.CompoundStatementFollows())
    return lib.Decoration{lib.ERR, nil}
  }
  return compoundStatementPrime(listing, tokens, symbols)
}

func subprogramSubbody(listing *list.List, tokens *list.List, symbols map[string]*lib.Symbol, stack *list.List, addresses *list.List, loc int) {
  t, ok := match(tokens, lib.FUNCTION)
  if ok {
    subprogramDeclarations(listing, tokens, symbols, stack, addresses, loc)
    compoundStatement(listing, tokens, symbols)
    return
  }
  t, ok = match(tokens, lib.BEGIN)
  if !ok {
    report(listing, "function OR begin", t)
    sync(tokens, lib.SubprogramSubbodyFollows())
    return
  }
  compoundStatement(listing, tokens, symbols)
}

func subprogramBody(listing *list.List, tokens *list.List, symbols map[string]*lib.Symbol, stack *list.List, addresses *list.List, loc int) {
  t, ok := match(tokens, lib.VAR)
  if ok {
    declarations(listing, tokens, symbols, stack, addresses, loc)
    subprogramSubbody(listing, tokens, symbols, stack, addresses, loc)
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
  subprogramSubbody(listing, tokens, symbols, stack, addresses, loc)
}

func subprogramDeclaration(listing *list.List, tokens *list.List, symbols map[string]*lib.Symbol, stack *list.List, addresses *list.List, loc int) {
  t, ok := match(tokens, lib.FUNCTION)
  if !ok {
    report(listing, "function", t)
    sync(tokens, lib.SubprogramDeclarationFollows())
    return
  }
  subprogramHead(listing, tokens, symbols, stack)
  loc = 0
  subprogramBody(listing, tokens, symbols, stack, addresses, loc)
}

func subprogramDeclarationsPrime(listing *list.List, tokens *list.List, symbols map[string]*lib.Symbol, stack *list.List, addresses *list.List, loc int) {
  t, ok := match(tokens, lib.FUNCTION)
  if ok {
    subprogramDeclaration(listing, tokens, symbols, stack, addresses, loc)
    popGreenNode(symbols, stack)
    if t, ok = matchYank(tokens, lib.SEMICOLON); !ok {
      report(listing, ";", t)
      sync(tokens, lib.SubprogramDeclarationsPrimeFollows())
      return
    }
    subprogramDeclarationsPrime(listing, tokens, symbols, stack, addresses, loc)
    return
  }
  if t, ok = match(tokens, lib.BEGIN); !ok {
    report(listing, "function OR begin", t)
    sync(tokens, lib.SubprogramDeclarationsPrimeFollows())
    return
  }
  // epsilon production
}

func subprogramDeclarations(listing *list.List, tokens *list.List, symbols map[string]*lib.Symbol, stack *list.List, addresses *list.List, loc int) {
  t, ok := match(tokens, lib.FUNCTION)
  if !ok {
    report(listing, "function", t)
    sync(tokens, lib.SubprogramDeclarationsFollows())
    return
  }
  subprogramDeclaration(listing, tokens, symbols, stack, addresses, loc)
  popGreenNode(symbols, stack)
  if t, ok = matchYank(tokens, lib.SEMICOLON); !ok {
    report(listing, ";", t)
    sync(tokens, lib.SubprogramDeclarationsFollows())
    return
  }
  subprogramDeclarationsPrime(listing, tokens, symbols, stack, addresses, loc)
}

func programSubbody(listing *list.List, tokens *list.List, symbols map[string]*lib.Symbol, stack *list.List, addresses *list.List, loc int) {
  t, ok := match(tokens, lib.FUNCTION)
  if ok {
    subprogramDeclarations(listing, tokens, symbols, stack, addresses, loc)
    compoundStatement(listing, tokens, symbols)
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
  compoundStatement(listing, tokens, symbols)
  if t, ok = matchYank(tokens, lib.PERIOD); !ok {
    report(listing, ".", t)
    sync(tokens, lib.ProgramSubbodyFollows())
    return
  }
}

func programBody(listing *list.List, tokens *list.List, symbols map[string]*lib.Symbol, stack *list.List, addresses *list.List, loc int) {
  t, ok := match(tokens, lib.VAR)
  if ok {
    declarations(listing, tokens, symbols, stack, addresses, loc)
    programSubbody(listing, tokens, symbols, stack, addresses, loc)
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
  programSubbody(listing, tokens, symbols, stack, addresses, loc)
}

func program(listing *list.List, tokens *list.List, symbols map[string]*lib.Symbol, stack *list.List, addresses *list.List, loc int) {
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
  id := t
  addType(id.Lexeme, lib.Decoration{lib.PROGRAM, nil}, symbols)
  checkAddGreenNode(listing, id, stack, id.Lexeme, lib.PROGRAM)
  if t, ok = matchYank(tokens, lib.OPEN_PAREN); !ok {
    report(listing, "(", t)
    sync(tokens, lib.ProgramFollows())
    return
  }
  identifierList(listing, tokens, symbols, stack)
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
  loc = 0
  programBody(listing, tokens, symbols, stack, addresses, loc)
  if t, ok = matchYank(tokens, lib.EOF); !ok {
    report(listing, "EOF", t)
    sync(tokens, lib.ProgramFollows())
    return
  }
  popGreenNode(symbols, stack)
}

func Parse(listing *list.List, tokens *list.List, symbols map[string]*lib.Symbol, stack *list.List, addresses *list.List, loc int) {
  program(listing, tokens, symbols, stack, addresses, loc)
}