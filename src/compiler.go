package main

import (
  "bufio"
  "compiler"
  "container/list"
  "fmt"
  "io"
  "lib"
  "os"
  "path/filepath"
)

func check(e error) {
  if e != nil {
    panic(e)
  }
}

func OutputListingFile(listing *list.List) {
  for e := listing.Front(); e != nil; e = e.Next() {
    l := e.Value.(lib.Line)
    fmt.Println(l.Number, l.Value)
    for e1 := l.Errors.Front(); e1 != nil; e1 = e1.Next() {
      fmt.Println(l.Number, e1.Value)
    }
    fmt.Print("\n")
  }
}

func OutputTokenFile(tokens *list.List) {
  for e := tokens.Front(); e != nil; e = e.Next() {
    t := e.Value.(lib.Token)
    fmt.Println(t.Lexeme, t.Type, t.Attr)
  }
  fmt.Print("\n")
}

func main() {
  fpath, e0 := filepath.Abs(os.Args[1])
  check(e0)
  f, e1 := os.Open(fpath)
  check(e1)
  defer f.Close()
  reader := bufio.NewReader(f)
  listing := list.New()
  line, _, e2 := reader.ReadLine()
  lineNumber := 1
  for e2 != io.EOF {
    check(e2)
    l := lib.Line{lineNumber, string(line[:]), list.New()}
    listing.PushBack(l)
    line, _, e2 = reader.ReadLine()
    lineNumber++
  }
  tokens := compiler.Tokenize(listing)
  OutputListingFile(listing)
  OutputTokenFile(tokens)
}