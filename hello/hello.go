/* hello.go - My first Golan program */

package main

import (
	"fmt"
)

type Vertex struct {
	x int
	y int
}

func main() {
	/* fmt.Printf("Hello, world\n")
	  i, j := 34, 20
	  p := &i
	  fmt.Println(*p)
	  *p = 22
	  fmt.Println(i)
	  fmt.Println(p)
	  p = &j         // point to j
		*p = *p / 4   // divide j through the pointer
		fmt.Println(j) // see the new value of j
	  for k := 0; k < 10; k++ {
	    fmt.Println(k)
	  }
	  v := Vertex{5, 8}
		p := &v
		p.x = 34
		fmt.Println(v) */

	var arr [2]int
	arr[0] = 1
	arr[1] = 2
	fmt.Println(arr[0])

	var s = []int{2, 3, 5, 7, 11, 13}
	fmt.Println("s ==", s)

}
