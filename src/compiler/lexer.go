package compiler

import (
  "container/list"
  "lib"
  "unicode"
)

func MatchWhitespace(l string, index int) int {
  i := index
  for i < len(l) && unicode.IsSpace(rune(l[i])) {
    i++
  }
  return i
}

func MatchId(l string, index int, ln int) (int, lib.Token) {
  if !unicode.IsLetter(rune(l[index])) {
    return index, lib.Token{}
  }
  i := index
  for i < len(l) && (unicode.IsLetter(rune(l[i])) || unicode.IsDigit(rune(l[i]))) {
    i++
  }
  return i, lib.Token{ln, l[index:i], lib.ID, lib.NULL}
}

func MatchLongReal(l string, index int, ln int) (int, lib.Token) {
  if !unicode.IsDigit(rune(l[index])) {
    return index, lib.Token{}
  }
  i := index
  for i < len(l) && unicode.IsDigit(rune(l[i])) {
    i++
  }
  if i >= len(l) || (string(l[i]) != "." && string(l[i]) != "E") {
    return index, lib.Token{}
  }
  if string(l[i]) == "." {
    i++
    if i >= len(l) || !unicode.IsDigit(rune(l[i])) {
      return index, lib.Token{}
    }
    for i < len(l) && unicode.IsDigit(rune(l[i])) {
      i++
    }
    if i >= len(l) || string(l[i]) != "E" {
      return index, lib.Token{}
    }
  }
  if string(l[i]) == "E" {
    i++
    if i <= len(l) && (string(l[i]) == "+" || string(l[i]) == "-") {
      i++
    }
    if i >= len(l) || !unicode.IsDigit(rune(l[i])) {
      return index, lib.Token{}
    }
    for i < len(l) && unicode.IsDigit(rune(l[i])) {
      i++
    }
  }
  return i, lib.Token{ln, l[index:i], lib.NUM, lib.LONG_REAL}
}

func MatchReal(l string, index int, ln int) (int, lib.Token) {
  if !unicode.IsDigit(rune(l[index])) {
    return index, lib.Token{}
  }
  i := index
  for i < len(l) && unicode.IsDigit(rune(l[i])) {
    i++
  }
  if i >= len(l) || string(l[i]) != "." {
    return index, lib.Token{}
  }
  if string(l[i]) == "." {
    i++
    if i >= len(l) || !unicode.IsDigit(rune(l[i])) {
      return index, lib.Token{}
    }
    for i < len(l) && unicode.IsDigit(rune(l[i])) {
      i++
    }
  }
  return i, lib.Token{ln, l[index:i], lib.NUM, lib.REAL}
}

func MatchInt(l string, index int, ln int) (int, lib.Token) {
  if !unicode.IsDigit(rune(l[index])) {
    return index, lib.Token{}
  }
  i := index
  for i < len(l) && unicode.IsDigit(rune(l[i])) {
    i++
  }
  return i, lib.Token{ln, l[index:i], lib.NUM, lib.INT}
}

func MatchRelop(l string, index int, ln int) (int, lib.Token) {
  if string(l[index]) == "=" {
    return index+1, lib.Token{ln, l[index:index+1], lib.RELOP, lib.EQ}
  }
  if string(l[index]) == "<" {
    if string(l[index+1]) == "=" {
      return index+2, lib.Token{ln, l[index:index+2], lib.RELOP, lib.LEQ}
    }
    if string(l[index+1]) == ">" {
      return index+2, lib.Token{ln, l[index:index+2], lib.RELOP, lib.NEQ}
    }
    return index+1, lib.Token{ln, l[index:index+1], lib.RELOP, lib.LT}
  }
  if string(l[index]) == ">" {
    if string(l[index+1]) == "=" {
      return index+2, lib.Token{ln, l[index:index+2], lib.RELOP, lib.GEQ}
    }
    return index+1, lib.Token{ln, l[index:index+1], lib.RELOP, lib.GT}
  }
  return index, lib.Token{}
}

func MatchAddop(l string, index int, ln int) (int, lib.Token) {
  if string(l[index]) == "+" {
    return index+1, lib.Token{ln, l[index:index+1], lib.ADDOP, lib.PLUS}
  }
  if string(l[index]) == "-" {
    return index+1, lib.Token{ln, l[index:index+1], lib.ADDOP, lib.MINUS}
  }
  return index, lib.Token{}
}

func MatchMulop(l string, index int, ln int) (int, lib.Token) {
  if string(l[index]) == "*" {
    return index+1, lib.Token{ln, l[index:index+1], lib.MULOP, lib.ASTERISK}
  }
  if string(l[index]) == "/" {
    return index+1, lib.Token{ln, l[index:index+1], lib.MULOP, lib.SLASH}
  }
  return index, lib.Token{}
}

func MatchAssignop(l string, index int, ln int) (int, lib.Token) {
  if string(l[index]) == ":" && string(l[index+1]) == "=" {
    return index+2, lib.Token{ln, l[index:index+2], lib.ASSIGNOP, lib.NULL}
  }
  return index, lib.Token{}
}

func CatchAll(l string, index int, ln int) (int, lib.Token) {
  c := string(l[index])
  if c == "." {
    if index+1 < len(l) && string(l[index+1]) == "." {
      return index+2, lib.Token{ln, l[index:index+2], lib.RANGE, lib.NULL}
    }
    return index+1, lib.Token{ln, l[index:index+1], lib.PERIOD, lib.NULL}
  }
  if c == "[" {
    return index+1, lib.Token{ln, l[index:index+1], lib.OPEN_BRACKET, lib.NULL}
  }
  if c == "]" {
    return index+1, lib.Token{ln, l[index:index+1], lib.CLOSE_BRACKET, lib.NULL}
  }
  if c == "(" {
    return index+1, lib.Token{ln, l[index:index+1], lib.OPEN_PAREN, lib.NULL}
  }
  if c == ")" {
    return index+1, lib.Token{ln, l[index:index+1], lib.CLOSE_PAREN, lib.NULL}
  }
  if c == ";" {
    return index+1, lib.Token{ln, l[index:index+1], lib.SEMICOLON, lib.NULL}
  }
  if c == ":" {
    return index+1, lib.Token{ln, l[index:index+1], lib.COLON, lib.NULL}
  }
  if c == "," {
    return index+1, lib.Token{ln, l[index:index+1], lib.COMMA, lib.NULL}
  }
  return index, lib.Token{}
}

func TokenizeLine(line lib.Line, tokens *list.List) {
  i := 0
  for i < len(line.Value) {
    i = MatchWhitespace(line.Value, i)
    if idx, t := MatchId(line.Value, i, line.Number); idx != i {
      i = idx
      tokens.PushBack(t)
      continue
    }
    if idx, t := MatchLongReal(line.Value, i, line.Number); idx != i {
      i = idx
      tokens.PushBack(t)
      continue
    }
    if idx, t := MatchReal(line.Value, i, line.Number); idx != i {
      i = idx
      tokens.PushBack(t)
      continue
    }
    if idx, t := MatchInt(line.Value, i, line.Number); idx != i {
      i = idx
      tokens.PushBack(t)
      continue
    }
    if idx, t := MatchRelop(line.Value, i, line.Number); idx != i {
      i = idx
      tokens.PushBack(t)
      continue
    }
    if idx, t := MatchAddop(line.Value, i, line.Number); idx != i {
      i = idx
      tokens.PushBack(t)
      continue
    }
    if idx, t := MatchMulop(line.Value, i, line.Number); idx != i {
      i = idx
      tokens.PushBack(t)
      continue
    }
    if idx, t := MatchAssignop(line.Value, i, line.Number); idx != i {
      i = idx
      tokens.PushBack(t)
      continue
    }
    if idx, t := CatchAll(line.Value, i, line.Number); idx != i {
      i = idx
      tokens.PushBack(t)
      continue
    }
    i++
  }
}

func Tokenize(listing *list.List) *list.List {
  tokens := list.New()
  for e := listing.Front(); e != nil; e = e.Next() {
    line := e.Value.(lib.Line)
    TokenizeLine(line, tokens)
  }
  return tokens
}