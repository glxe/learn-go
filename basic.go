package main

import (
	"fmt"
	"math"
)

const (
	cpp = iota
	java
	python
	golang
)
const (
	b = 1 << (10 * iota)
	kb
	mb
	gb
	tb
	pb
)

func variable() {
}

func triangle() {
	var a, b int = 3, 4
	var c int

	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}

func main() {
	fmt.Println("Hello World!")

	triangle()

	fmt.Println(cpp, java, python, golang)
	fmt.Println(b, kb, mb, gb, tb, pb)
}
