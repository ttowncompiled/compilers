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
  "regexp"
  "strconv"
)

func check(e error) {
  if e != nil {
    panic(e)
  }
}

func ReadSourceFile(fpath string) *list.List {
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
  return listing
}

func ReadReservedWordFile(rpath string) map[string]int {
  f, e1 := os.Open(rpath)
  check(e1)
  defer f.Close()
  reader := bufio.NewReader(f)
  m := make(map[string]int)
  line, _, e2 := reader.ReadLine()
  for e2 != io.EOF {
    check(e2)
    l := regexp.MustCompile(" ").Split(string(line), -1)
    value, e3 := strconv.ParseInt(l[1], 10, 64)
    check(e3)
    m[l[0]] = int(value)
    line, _, e2 = reader.ReadLine()
  }
  return m
}

func OutputListingFile(listing *list.List) {
  for e := listing.Front(); e != nil; e = e.Next() {
    l := e.Value.(lib.Line)
    fmt.Println(l.Number, l.Value)
    for e1 := l.Errors.Front(); e1 != nil; e1 = e1.Next() {
      fmt.Println(l.Number, e1.Value.(lib.Error).Reason)
    }
    fmt.Print("\n")
  }
}

func OutputTokenFile(tokens *list.List, symbols map[string]*lib.Token) {
  fmt.Printf("%-20s %-20s %-20s\n", "Lexeme", "Type", "Attribute")
  fmt.Println("--------------------------------------------------------------")
  for e := tokens.Front(); e != nil; e = e.Next() {
    t := e.Value.(lib.Token)
    if t.Type == lib.ID {
      address := symbols[t.Lexeme]
      fmt.Printf("%-20s %-2d %-17s %-20p\n", t.Lexeme, t.Type, lib.Annotate(t.Type), address)
    } else {
      fmt.Printf("%-20s %-2d %-17s %-2d %-17s\n", t.Lexeme, t.Type, lib.Annotate(t.Type), t.Attr, lib.Annotate(t.Attr))
    }
  }
  fmt.Print("\n")
}

func OutputSymbolFile(symbols map[string]*lib.Token) {
  fmt.Printf("%-20s %-20s\n", "Lexeme", "Token")
  fmt.Println("-----------------------------------------")
  for k, v := range symbols {
    fmt.Printf("%-20s %-20p\n", k, v)
  }
  fmt.Print("\n")
}

func main() {
  if (len(os.Args) < 3) {
    fmt.Println("usage: go run compiler.go <pascal-file> <reserved-words-file>")
    os.Exit(1)
  }
  fpath, e0 := filepath.Abs(os.Args[1])
  check(e0)
  listing := ReadSourceFile(fpath)
  rpath, e1 := filepath.Abs(os.Args[2])
  check(e1)
  rwords := ReadReservedWordFile(rpath)
  tokens, symbols := compiler.Tokenize(listing, rwords)
  OutputListingFile(listing)
  OutputTokenFile(tokens, symbols)
  OutputSymbolFile(symbols)
}