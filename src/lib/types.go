package lib

import "container/list"

type Line struct {
  Number int
  Value string
  Errors *list.List
}

type Token struct {
  LineNumber int
  Lexeme string
}