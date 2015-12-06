package lib

import "container/list"

type Line struct {
  Number int
  Value string
  Errors *list.List
}

type Error struct {
  Reason string
  Value *Token
}

type Token struct {
  LineNumber int
  Lexeme string
  Type int
  Attr int
}