package lib

import "container/list"

func ProgramFollows() *list.List {
  return list.New()
}

func IdentifierListFollows() *list.List {
  follows := list.New()
  follows.PushBack(CLOSE_PAREN)
  return follows
}

func IdentifierListPrimeFollows() *list.List {
  follows := list.New()
  follows.PushBack(CLOSE_PAREN)
  return follows
}