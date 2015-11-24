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

func MatchId(l string, index int) (int, string) {
  if !unicode.IsLetter(rune(l[index])) {
    return index, ""
  }
  i := index
  for i < len(l) && (unicode.IsLetter(rune(l[i])) || unicode.IsDigit(rune(l[i]))) {
    i++
  }
  return i, l[index:i]
}

func MatchLongReal(l string, index int) (int, string) {
  if !unicode.IsDigit(rune(l[index])) {
    return index, ""
  }
  i := index
  for i < len(l) && unicode.IsDigit(rune(l[i])) {
    i++
  }
  if i >= len(l) || (string(l[i]) != "." && string(l[i]) != "E") {
    return index, ""
  }
  if string(l[i]) == "." {
    i++
    if i >= len(l) || !unicode.IsDigit(rune(l[i])) {
      return index, ""
    }
    for i < len(l) && unicode.IsDigit(rune(l[i])) {
      i++
    }
    if i >= len(l) || string(l[i]) != "E" {
      return index, ""
    }
  }
  if string(l[i]) == "E" {
    i++
    if i <= len(l) && (string(l[i]) == "+" || string(l[i]) == "-") {
      i++
    }
    if i >= len(l) || !unicode.IsDigit(rune(l[i])) {
      return index, ""
    }
    for i < len(l) && unicode.IsDigit(rune(l[i])) {
      i++
    }
  }
  return i, l[index:i]
}

func MatchReal(l string, index int) (int, string) {
  if !unicode.IsDigit(rune(l[index])) {
    return index, ""
  }
  i := index
  for i < len(l) && unicode.IsDigit(rune(l[i])) {
    i++
  }
  if i >= len(l) || string(l[i]) != "." {
    return index, ""
  }
  if string(l[i]) == "." {
    i++
    if i >= len(l) || !unicode.IsDigit(rune(l[i])) {
      return index, ""
    }
    for i < len(l) && unicode.IsDigit(rune(l[i])) {
      i++
    }
  }
  return i, l[index:i]
}

func MatchInt(l string, index int) (int, string) {
  if !unicode.IsDigit(rune(l[index])) {
    return index, ""
  }
  i := index
  for i < len(l) && unicode.IsDigit(rune(l[i])) {
    i++
  }
  return i, l[index:i]
}

func MatchRelop(l string, index int) (int, string) {
  i := index
  if string(l[i]) == "=" {
    i++
    return i, l[index:i]
  }
  if string(l[i]) == "<" {
    i++
    if string(l[i]) == "=" || string(l[i]) == ">" {
      i++
    }
    return i, l[index:i]
  }
  if string(l[i]) == ">" {
    i++
    if string(l[i]) == "=" {
      i++
    }
    return i, l[index:i]
  }
  return index, ""
}

func MatchAddop(l string, index int) (int, string) {
  i := index
  if string(l[i]) == "+" || string(l[i]) == "-" {
    i++
    return i, l[index:i]
  }
  return index, ""
}

func MatchMulop(l string, index int) (int, string) {
  i := index
  if string(l[i]) == "*" || string(l[i]) == "/" {
    i++
    return i, l[index:i]
  }
  return index, ""
}

func MatchAssignop(l string, index int) (int, string) {
  i := index
  if string(l[i]) == ":" && string(l[i+1]) == "=" {
    i += 2
    return i, l[index:i]
  }
  return index, ""
}

func CatchAll(l string, index int) (int, string) {
  i := index
  c := string(l[i])
  if c == "." {
    i++
    if i < len(l) && string(l[i]) == "." {
      i++
    }
    return i, l[index:i]
  }
  if c == "[" || c == "]" || c == "(" || c == ")" || c == ";" || c == ":" || c == "," {
    i++
    return i, l[index:i]
  }
  return index, ""
}

func TokenizeLine(line lib.Line, tokens *list.List) {
  i := 0
  for i < len(line.Value) {
    i = MatchWhitespace(line.Value, i)
    if idx, lex := MatchId(line.Value, i); lex != "" {
      i = idx
      tokens.PushBack(lib.Token{line.Number, lex})
      continue
    }
    if idx, lex := MatchLongReal(line.Value, i); lex != "" {
      i = idx
      tokens.PushBack(lib.Token{line.Number, lex})
      continue
    }
    if idx, lex := MatchReal(line.Value, i); lex != "" {
      i = idx
      tokens.PushBack(lib.Token{line.Number, lex})
      continue
    }
    if idx, lex := MatchInt(line.Value, i); lex != "" {
      i = idx
      tokens.PushBack(lib.Token{line.Number, lex})
      continue
    }
    if idx, lex := MatchRelop(line.Value, i); lex != "" {
      i = idx
      tokens.PushBack(lib.Token{line.Number, lex})
      continue
    }
    if idx, lex := MatchAddop(line.Value, i); lex != "" {
      i = idx
      tokens.PushBack(lib.Token{line.Number, lex})
      continue
    }
    if idx, lex := MatchMulop(line.Value, i); lex != "" {
      i = idx
      tokens.PushBack(lib.Token{line.Number, lex})
      continue
    }
    if idx, lex := MatchAssignop(line.Value, i); lex != "" {
      i = idx
      tokens.PushBack(lib.Token{line.Number, lex})
      continue
    }
    if idx, lex := CatchAll(line.Value, i); lex != "" {
      i = idx
      tokens.PushBack(lib.Token{line.Number, lex})
      continue
    }
    i++
  }
}

func Tokenize(listing *list.List) {
  tokens := list.New()
  for e := listing.Front(); e != nil; e = e.Next() {
    line := e.Value.(lib.Line)
    TokenizeLine(line, tokens)
  }
}