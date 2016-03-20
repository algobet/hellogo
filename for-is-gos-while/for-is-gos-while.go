package main

import "fmt"

func main() {
  sum := 1
  for sum<1000 { // 'for' in Go is like 'while' in C
    sum += sum
  }
  fmt.Println(sum)
}
