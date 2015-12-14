package compiler

import (
  "container/list"
  "lib"
  "strings"
  "unicode"
)

func MatchWhitespace(l string, index int) int {
  i := index
  for i < len(l) && unicode.IsSpace(rune(l[i])) {
    i++
  }
  return i
}

func MatchId(line lib.Line, index int, rwords map[string]int, symbols map[string]*lib.Token) (int, lib.Token) {
  l := line.Value
  if !unicode.IsLetter(rune(l[index])) {
    return index, lib.Token{}
  }
  i := index
  for i < len(l) && (unicode.IsLetter(rune(l[i])) || unicode.IsDigit(rune(l[i]))) {
    i++
  }
  lex := l[index:i]
  if val, ok := rwords[lex]; ok {
    return i, lib.Token{line.Number, lex, val, lib.NULL}
  }
  if (len(lex) > 10) {
    t := lib.Token{line.Number, lex, lib.LEXERR, lib.ID_TOO_LONG}
    line.Errors.PushBack(lib.Error{"LEXERR: EXTRA LONG ID: " + t.Lexeme, &t})
    return i, t
  }
  t := lib.Token{line.Number, lex, lib.ID, lib.NULL}
  if _, ok := symbols[lex]; !ok {
    symbols[lex] = &t
  }
  return i, t
}

func MatchLongReal(line lib.Line, index int) (int, lib.Token) {
  l := line.Value
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
  lexeme := l[index:i]
  if (len(lexeme) > 1 && string(lexeme[0]) == "0" && unicode.IsDigit(rune(lexeme[1]))) {
    t := lib.Token{line.Number, lexeme, lib.LEXERR, lib.LEADING_ZEROS}
    line.Errors.PushBack(lib.Error{"LEXERR: LEADING ZEROS: " + t.Lexeme, &t})
    return i, t
  }
  if (strings.Index(lexeme, ".") > 5) {
    t := lib.Token{line.Number, lexeme, lib.LEXERR, lib.XX_TOO_LONG}
    line.Errors.PushBack(lib.Error{"LEXERR: EXTRA LONG CHARACTERISTIC: " + lexeme, &t})
    return i, t
  }
  if (strings.Index(lexeme, "E") - strings.Index(lexeme, ".") > 6) {
    t := lib.Token{line.Number, lexeme, lib.LEXERR, lib.YY_TOO_LONG}
    line.Errors.PushBack(lib.Error{"LEXERR: EXTRA LONG FRACTIONAL PART: " + lexeme, &t})
    return i, t
  }
  if (strings.Index(lexeme, ".")+2 < strings.Index(lexeme, "E") && string(lexeme[strings.Index(lexeme, "E")-1]) == "0" && unicode.IsDigit(rune(lexeme[strings.Index(lexeme, "E")-2]))) {
    t := lib.Token{line.Number, lexeme, lib.LEXERR, lib.TRAILING_ZEROS}
    line.Errors.PushBack(lib.Error{"LEXERR: TRAILING ZEROS: " + t.Lexeme, &t})
    return i, t
  }
  if (len(lexeme) - strings.Index(lexeme, "E") > 3) {
    t := lib.Token{line.Number, lexeme, lib.LEXERR, lib.ZZ_TOO_LONG}
    line.Errors.PushBack(lib.Error{"LEXERR: EXTRA LONG EXPONENTIAL PART: " + lexeme, &t})
    return i, t
  }
  if (strings.Index(lexeme, "E")+2 < len(lexeme) && string(lexeme[strings.Index(lexeme, "E")+1]) == "0" && unicode.IsDigit(rune(lexeme[strings.Index(lexeme, "E")+2]))) {
    t := lib.Token{line.Number, lexeme, lib.LEXERR, lib.LEADING_ZEROS}
    line.Errors.PushBack(lib.Error{"LEXERR: LEADING ZEROS: " + t.Lexeme, &t})
    return i, t
  }
  return i, lib.Token{line.Number, lexeme, lib.NUM, lib.LONG_REAL}
}

func MatchReal(line lib.Line, index int) (int, lib.Token) {
  l := line.Value
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
  lexeme := l[index:i]
  if (len(lexeme) > 1 && string(lexeme[0]) == "0" && unicode.IsDigit(rune(lexeme[1]))) {
    t := lib.Token{line.Number, lexeme, lib.LEXERR, lib.LEADING_ZEROS}
    line.Errors.PushBack(lib.Error{"LEXERR: LEADING ZEROS: " + t.Lexeme, &t})
    return i, t
  }
  if (strings.Index(lexeme, ".") > 5) {
    t := lib.Token{line.Number, lexeme, lib.LEXERR, lib.XX_TOO_LONG}
    line.Errors.PushBack(lib.Error{"LEXERR: EXTRA LONG CHARACTERISTIC: " + lexeme, &t})
    return i, t
  }
  if (len(lexeme) - strings.Index(lexeme, ".") > 6) {
    t := lib.Token{line.Number, lexeme, lib.LEXERR, lib.YY_TOO_LONG}
    line.Errors.PushBack(lib.Error{"LEXERR: EXTRA LONG FRACTIONAL PART: " + lexeme, &t})
    return i, t
  }
  if (strings.Index(lexeme, ".")+2 < len(lexeme) && string(lexeme[len(lexeme)-1]) == "0" && unicode.IsDigit(rune(lexeme[len(lexeme)-2]))) {
    t := lib.Token{line.Number, lexeme, lib.LEXERR, lib.TRAILING_ZEROS}
    line.Errors.PushBack(lib.Error{"LEXERR: TRAILING ZEROS: " + t.Lexeme, &t})
    return i, t
  }
  return i, lib.Token{line.Number, lexeme, lib.NUM, lib.REAL}
}

func MatchInt(line lib.Line, index int) (int, lib.Token) {
  l := line.Value
  if !unicode.IsDigit(rune(l[index])) {
    return index, lib.Token{}
  }
  i := index
  for i < len(l) && unicode.IsDigit(rune(l[i])) {
    i++
  }
  lexeme := l[index:i]
  if (len(lexeme) > 1 && string(lexeme[0]) == "0" && unicode.IsDigit(rune(lexeme[1]))) {
    t := lib.Token{line.Number, lexeme, lib.LEXERR, lib.LEADING_ZEROS}
    line.Errors.PushBack(lib.Error{"LEXERR: LEADING ZEROS: " + t.Lexeme, &t})
    return i, t
  }
  if (len(lexeme) > 10) {
    t := lib.Token{line.Number, lexeme, lib.LEXERR, lib.XX_TOO_LONG}
    line.Errors.PushBack(lib.Error{"LEXERR: EXTRA LONG INTEGER: " + lexeme, &t})
    return i, t
  }
  return i, lib.Token{line.Number, lexeme, lib.NUM, lib.INT}
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

func TokenizeLine(line lib.Line, tokens *list.List, rwords map[string]int, symbols map[string]*lib.Token) {
  i := 0
  for i < len(line.Value) {
    i = MatchWhitespace(line.Value, i)
    if idx, t := MatchId(line, i, rwords, symbols); idx != i {
      i = idx
      tokens.PushBack(t)
      continue
    }
    if idx, t := MatchLongReal(line, i); idx != i {
      i = idx
      tokens.PushBack(t)
      continue
    }
    if idx, t := MatchReal(line, i); idx != i {
      i = idx
      tokens.PushBack(t)
      continue
    }
    if idx, t := MatchInt(line, i); idx != i {
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
    unrecognizedSymbol := lib.Token{line.Number, string(line.Value[i]), lib.LEXERR, lib.UNRECOGNIZED_SYMBOL}
    tokens.PushBack(unrecognizedSymbol)
    line.Errors.PushBack(lib.Error{"LEXERR: unrecognized symbol '" + unrecognizedSymbol.Lexeme + "'", &unrecognizedSymbol})
    i++
  }
}

func Tokenize(listing *list.List, rwords map[string]int) (*list.List, map[string]*lib.Token) {
  tokens := list.New()
  symbols := make(map[string]*lib.Token)
  for e := listing.Front(); e != nil; e = e.Next() {
    line := e.Value.(lib.Line)
    TokenizeLine(line, tokens, rwords, symbols)
  }
  eofToken := lib.Token{listing.Len(), "", lib.EOF, lib.NULL}
  tokens.PushBack(eofToken)
  return tokens, symbols
}