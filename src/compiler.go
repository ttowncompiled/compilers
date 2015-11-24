package main

import (
  "bufio"
  "compiler"
  "container/list"
  "fmt"
  "io"
  "os"
  "path/filepath"
)

type Line struct {
  number int
  value string
  errors *list.List
}

func check(e error) {
  if e != nil {
    panic(e)
  }
}

func OutputListingFile(listing *list.List) {
  for e := listing.Front(); e != nil; e = e.Next() {
    l := e.Value.(Line)
    fmt.Println(l.number, l.value)
    for e1 := l.errors.Front(); e1 != nil; e1 = e1.Next() {
      fmt.Println(l.number, e1.Value)
    }
    fmt.Print("\n")
  }
}

func main() {
  filepath, e0 := filepath.Abs(os.Args[1])
  check(e0)
  file, e1 := os.Open(filepath)
  check(e1)
  defer file.Close()
  reader := bufio.NewReader(file)
  listing := list.New()
  line, _, e2 := reader.ReadLine()
  lineNumber := 1
  for e2 != io.EOF {
    check(e2)
    l := Line{lineNumber, string(line[:]), list.New()}
    listing.PushBack(l)
    line, _, e2 = reader.ReadLine()
    lineNumber++
  }
  compiler.Tokenize(listing)
  OutputListingFile(listing)
}