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

type Rword struct {
  Type int
  Attr int
}

type Token struct {
  LineNumber int
  Lexeme string
  Type int
  Attr int
}

type TypeD interface {
  TypeD() int
}

type Decoration struct {
  Type int
}

func (self *Decoration) TypeD() int {
  return self.Type
}

type ArrayD struct {
  Size int
  Val int
}

func (self *ArrayD) TypeD() int {
  return ARRAY
}

type FunctionD struct {
  Params *list.List
  Return int
}

func (self *FunctionD) TypeD() int {
  return FUNCTION
}

type Symbol struct {
  Token *Token
  Decoration *Decoration
}