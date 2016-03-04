package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	word := strings.Fields(s)
	count := len(word)
  m := make(map[string]int)
  for i := 0; i < count; i++ {
    m[word[i]]++
  }
	return m
}

func main() {
	wc.Test(WordCount)
}
