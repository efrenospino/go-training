package main

import (
	"fmt"
	"strings"
)

func areIsomorphics(word1, word2 string) (ans bool) {
	wordt1, wordt2 := word1, word2
	w1 := strings.Split(word1, "")
	w2 := strings.Split(word2, "")
	for i := 0; i < len(w1); i++ {
		word1 = strings.Replace(word1, w1[i], w2[i], -1)
	}
	fmt.Print("Are ", wordt1, " and ", wordt2, " isomorphics?: ")
	if word1 == word2 {
		ans = true
	}
	return ans
}

func main() {
	fmt.Println(areIsomorphics("foo", "bar"))
	fmt.Println(areIsomorphics("egg", "add"))
	fmt.Println(areIsomorphics("paper", "title"))
}
