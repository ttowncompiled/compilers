package compiler

import (
  "container/list"
  "lib"
)

func match(tokens *list.List, expectedType int) (lib.Token, bool) {
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

func identifierList(listing *list.List, tokens *list.List) {
  
}

func program(listing *list.List, tokens *list.List) {
  follows := list.New()
  t, ok := match(tokens, lib.PROGRAM)
  if !ok {
    report(listing, "program", t)
    sync(tokens, follows)
    return
  }
  if t, ok = match(tokens, lib.ID); !ok {
    report(listing, "ID", t)
    sync(tokens, follows)
    return
  }
  if t, ok = match(tokens, lib.OPEN_PAREN); !ok {
    report(listing, "(", t)
    sync(tokens, follows)
    return
  }
  identifierList(listing, tokens)
  if t, ok = match(tokens, lib.CLOSE_PAREN); !ok {
    report(listing, ")", t)
    sync(tokens, follows)
    return
  }
  if t, ok = match(tokens, lib.SEMICOLON); !ok {
    report(listing, ";", t)
    sync(tokens, follows)
    return
  }
}

func Parse(listing *list.List, tokens *list.List, symbols map[string]*lib.Token) {
  program(listing, tokens)
}