package main

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
)

// Color is type that receives a RGBA color
type Color struct {
	Red   int
	Green int
	Blue  int
	Alpha int
}

var m = make(map[string]Color)

func add(name string) (string, Color) {
	var color Color
	var ok bool
	if color, ok = m[name]; ok != true {
		color = Color{rand.Intn(255), rand.Intn(255), rand.Intn(255), rand.Intn(255)}
		m[name] = color
	}
	return name, color
}

func normalizeUpperAndLowerLetters(name string) string {
	return strings.ToUpper(name)
}

func insertNames() {
	var name string
	for name != "END" {
		fmt.Println("Type a name: ")
		_, err := fmt.Scan(&name)
		if err != nil {
			fmt.Println("Error: ", err)
		}
		name = normalizeUpperAndLowerLetters(name)
		if name != "END" {
			_, color := add(name)
			fmt.Printf("%s -> (%d %d %d %d)\n", name, color.Red, color.Green, color.Blue, color.Alpha)
		}
	}
}

func main() {
	dir, _ := os.Getwd()
	fmt.Println(dir)
	fmt.Println(filepath.Join(dir, "..", "database"))

	insertNames()

}
