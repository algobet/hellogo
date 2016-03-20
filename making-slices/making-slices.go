package main

import "fmt"

func main() {
	a := make([]int, 5)     // len(a)=5
	printSlice("a", a)

	b := make([]int, 0, 5)  // len(b)=0, cap(b)=5
	printSlice("b", b)

	c := b[:2]              // len(c)=3, cap(c)=3
	printSlice("c", c)

	d := c[2:5]             // len(d)=4, cap(c)=
	printSlice("d", d)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
