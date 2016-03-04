package main

import "fmt"

type Vertex struct {
  Lat, Long float64
}

//var m map[string]Vertex

func main() {
  m := make(map[string]Vertex)
  m["Bell Labs"] = Vertex{40.68895, -63.68347,}

  m1 := map[string]Vertex{
    "Bell Labs": Vertex{
      86.74574, -63.346466,
    },
    "Google": Vertex{
      36.634563, -35.564355,
    },
  }

  fmt.Println(m["Bell Labs"])
  fmt.Println(m1["Google"])

  isIn, ok := m1["Google"]
  fmt.Println("The value:", isIn, "Present?", ok)
  fmt.Println(m1)
  delete(m1, "Google")
  isIn, ok = m1["Google"]
  fmt.Println(m1)
  fmt.Println("The value:", isIn, "Present?", ok)
}
