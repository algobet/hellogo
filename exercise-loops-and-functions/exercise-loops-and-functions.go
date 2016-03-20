package main

import (
  "fmt"
)

// 牛顿法实现开方函数(循环10次)
// z = z - (z*z - x)/2/z
func Sqrt(x float64) float64 {
  z := float64(1)
  for i := 0; i<10; i++ {
    z -= (z*z - x)/2/z
  }
  return z
}

func main() {
  fmt.Println(Sqrt(2))
}
