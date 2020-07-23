package main

import "fmt"

func printSlice(s []int) {
	fmt.Printf("%v, len(s) = %d, cap(s) = %d\n", s, len(s), cap(s))
}
func main() {
	fmt.Println("Creating slice")
	var s []int

	for i := 0; i < 100; i++ {
		printSlice(s)
		s = append(s, 2*i+1)
	}
	/*
		[], len(s) = 0, cap(s) = 0
		[1], len(s) = 1, cap(s) = 1
		[1 3], len(s) = 2, cap(s) = 2
		[1 3 5], len(s) = 3, cap(s) = 4
		[1 3 5 7], len(s) = 4, cap(s) = 4
		[1 3 5 7 9], len(s) = 5, cap(s) = 8
		[1 3 5 7 9 11], len(s) = 6, cap(s) = 8
		[1 3 5 7 9 11 13], len(s) = 7, cap(s) = 8
		[1 3 5 7 9 11 13 15], len(s) = 8, cap(s) = 8
		[1 3 5 7 9 11 13 15 17], len(s) = 9, cap(s) = 16
		[1 3 5 7 9 11 13 15 17 19], len(s) = 10, cap(s) = 16
		[1 3 5 7 9 11 13 15 17 19 21], len(s) = 11, cap(s) = 16
		[1 3 5 7 9 11 13 15 17 19 21 23], len(s) = 12, cap(s) = 16
	*/

	fmt.Println(s)

	s1 := []int{2, 4, 6, 8}
	printSlice(s1)
	// [2 4 6 8], len(s) = 4, cap(s) = 4

	s2 := make([]int, 16)
	s3 := make([]int, 10, 32)
	printSlice(s2)
	printSlice(s3)
	// [0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0], len(s) = 16, cap(s) = 16
	// [0 0 0 0 0 0 0 0 0 0], len(s) = 10, cap(s) = 32
	fmt.Println("copying slice")
	copy(s2, s1)
	printSlice(s2)
	// [2 4 6 8 0 0 0 0 0 0 0 0 0 0 0 0], len(s) = 16, cap(s) = 16
	fmt.Println("deleting elements from slice")
	s2 = append(s2[:3], s2[4:]...) // php is ...s2
	fmt.Println(s2)
	// [2 4 6 0 0 0 0 0 0 0 0 0 0 0 0]

	fmt.Println("Popping from front")
	front := s2[0]
	s2 = s2[1:]

	fmt.Println("Popping from back")
	tail := s2[len(s2)-1]
	s2 = s2[:len(s2)-1]
	fmt.Println(front, tail)
	printSlice(s2)
}
