package main

import "fmt"

type Vertex struct {
  Lat, Long float64
}

var m map[string]Vertex

func main() {
  m = make(map[string]Vertex)
  m["algobet Labs"] = Vertex {21.00001, 120.00001}
  fmt.Println(m["algobet Labs"])
}
