package main

import "fmt"

type Vertex struct {
  Lat, Long float64
}

var m = map[string]Vertex {
  "algobet Labs": {21.00001, 120.00001},
  "Bell Labs"   : {40.68433, -74.39967},
  "Google"      : {37.42202, -122.08408},
}

func main() {
  fmt.Println(m)
}
