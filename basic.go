package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
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

func eval(a, b int, op string) (int, error) {
	var result int
	var err error
	switch op {
	case "+":
		result, err = a+b, nil
	case "-":
		result, err = a-b, nil
	case "*":
		result, err = a*b, nil
	case "/":
		result, err = a/b, nil
	default:
		result, err = 0, fmt.Errorf("unsupported operation: %s", op)
	}
	return result, err
}

func apply(op func(int, int) int, a, b int) int {

	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling function %s with args "+
		"(%d, %d) ", opName, a, b)
	return op(a, b)
}

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

func sum(numbers ...int) int {
	s := 0
	for i := range numbers {
		s += numbers[i]
	}
	return s
}

func main() {
	fmt.Println("Hello World!")

	triangle()

	fmt.Println(cpp, java, python, golang)
	fmt.Println(b, kb, mb, gb, tb, pb)
	if res, err := eval(4, 6, "*"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}

	fmt.Println(apply(pow, 3, 4))
	// OR 写成匿名函数
	fmt.Println(apply(func(a int, b int) int {
		return int(math.Pow(
			float64(a), float64(b)))
	}, 3, 4))

	fmt.Println(sum(1, 2, 3, 4, 5))
}
