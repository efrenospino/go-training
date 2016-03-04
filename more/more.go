package main

import (
  "fmt"
  "strings"
)

func main() {
  /*game := [][]string{
    []string{"_", "_", "_"},
    []string{"_", "_", "_"},
    []string{"_", "_", "_"},
  }

  printBoard(game)
  */
  a := make([]int, 4, 5)
  printSlice("a", a)
  a = append(a, 0, 0)
  printSlice("a", a)
}

func printBoard(s [][]string)  {
  for i := 0; i < len(s); i++ {
    fmt.Printf("%s\n", strings.Join(s[i], " "))
  }
}

func printSlice(s string, x []int)  {
  fmt.Printf("%s len=%d cap=%d %s=%v\n", s, len(x), cap(x), s, x)
  if len(x) == 0 && cap(x) == 0 {
    fmt.Println("Is nil!")
  }
}
