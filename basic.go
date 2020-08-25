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
func printArray(arr [5]int) {
	for i := range arr {
		fmt.Println(arr[i])
	}
}

func updateSlice(s []int) {
	s[0] = 100
}

func pop(s []int) int {
	len := len(s)
	return s[len-1 : len][0]
}

func main() {
	fmt.Println("Hello World!")

	// triangle()

	// fmt.Println(cpp, java, python, golang)
	// fmt.Println(b, kb, mb, gb, tb, pb)
	// if res, err := eval(4, 6, "*"); err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(res)
	// }

	// fmt.Println(apply(pow, 3, 4))
	// // OR 写成匿名函数
	// fmt.Println(apply(func(a int, b int) int {
	// 	return int(math.Pow(
	// 		float64(a), float64(b)))
	// }, 3, 4))

	// fmt.Println(sum(1, 2, 3, 4, 5))
	// fmt.Println("--------------------------------------------------")

	// var arr1 [5]int
	// arr2 := [3]int{1, 2, 3}            // 如果是 := 方式赋值，则必须要制定数组里的值
	// arr3 := [...]int{9, 8, 7, 6, 5, 4} // 可以让编译器自己推断长度，但是必须加...
	// var gride [5][6]int                // 二维数组
	// fmt.Println(arr1, arr2, arr3)
	// fmt.Println(gride)

	// for i := 0; i < len(arr3); i++ {
	// 	fmt.Println(arr3[i])
	// }

	// for i := range arr3 {
	// 	fmt.Println(arr3[i])
	// }

	// for i, v := range arr3 {
	// 	fmt.Println(i, v)
	// }
	// for _, v := range arr3 {
	// 	fmt.Println(v)
	// }
	// fmt.Println("--------------------------------------------------")

	// printArray(arr1)
	// fmt.Println("--------------------------------------------------")

	// arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	// fmt.Println("arr[2:6] = ", arr[2:6])
	// fmt.Println("arr[2:] = ", arr[2:])
	// s1 := arr[2:6]
	// fmt.Println("arr[2:6] = ", s1)
	// s2 := arr[:]
	// fmt.Println("arr[:] = ", s2)
	// fmt.Println("--------------------------------------------------")

	// fmt.Println("after updateSlice (s1)")
	// updateSlice(s1)
	// fmt.Println(s1)
	// fmt.Println(arr)

	// fmt.Println("after updateSlice (s2)")
	// updateSlice(s2)
	// fmt.Println(s2)
	// fmt.Println(arr)

	// fmt.Println("--------------------------------------------------")

	// arr11 := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	// s11 := arr11[2:6]
	// s22 := s11[3:5]q
	// fmt.Printf("s11 = %v, len(s11) = %d, cap(s11) = %d\n",
	// 	s11, len(s11), cap(s11))
	// fmt.Printf("s22 = %v, len(s22) = %d, cap(s22) = %d\n",
	// 	s22, len(s22), cap(s22))

	// s33 := append(s22, 10)
	// s44 := append(s33, 11)
	// s55 := append(s44, 12)
	// fmt.Println("s33, s44, s55 =", s33, s44, s55)
	// fmt.Println("arr11 =", arr11)

	s66 := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(pop(s66[0:]))

	s77 := make([]int, 16, 32)
	copy(s77, s66[:])
	fmt.Println(s77)

	s77 = append(s77[:8], s77[9:]...)
	fmt.Println(s77)
}
