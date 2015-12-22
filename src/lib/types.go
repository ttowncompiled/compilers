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
  PrevTypeD() *TypeD
}

type Decoration struct {
  Type int
  Prev *TypeD
}

func (self Decoration) TypeD() int {
  return self.Type
}

func (self Decoration) PrevTypeD() *TypeD {
  return self.Prev
}

type ArrayD struct {
  Size int
  Val int
  Prev *TypeD
}

func (self ArrayD) TypeD() int {
  return ARRAY
}

func (self ArrayD) PrevTypeD() *TypeD {
  return self.Prev
}

type FunctionD struct {
  Params *list.List
  Return int
  Prev *TypeD
}

func (self FunctionD) TypeD() int {
  return FUNCTION
}

func (self FunctionD) PrevTypeD() *TypeD {
  return self.Prev
}

type Symbol struct {
  Token *Token
  Decoration TypeD
}