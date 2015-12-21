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
  SetPrevTypeD(prevType TypeD)
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

func (self Decoration) SetPrevTypeD(prevType TypeD) {
  self.Prev = &prevType
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

func (self ArrayD) SetPrevTypeD(prevType TypeD) {
  self.Prev = &prevType
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

func (self FunctionD) SetPrevTypeD(prevType TypeD) {
  self.Prev = &prevType
}

type Symbol struct {
  Token *Token
  Decoration TypeD
}