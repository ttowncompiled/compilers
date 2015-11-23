package main

import (
  "bufio"
  "container/list"
  "fmt"
  "io"
  "os"
  "path/filepath"
)

func check(e error) {
  if e != nil {
    panic(e)
  }
}

func write_listing_file(listing *list.List) {
  line_number := 0
  for e := listing.Front(); e != nil; e = e.Next() {
    line_number++
    fmt.Println(line_number, e.Value, "\n")
  }
}

func main() {
  filepath, err := filepath.Abs(os.Args[1])
  check(err)

  file, err := os.Open(filepath)
  check(err)
  
  reader := bufio.NewReader(file)
  listing_file := list.New()
  
  line, _, err := reader.ReadLine()
  for err != io.EOF {
    check(err)
    line, _, err = reader.ReadLine()
    listing_file.PushBack(string(line[:]))
  }
  
  write_listing_file(listing_file)
}