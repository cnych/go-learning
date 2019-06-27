package main

import "fmt"

func changeEle(arr []int) {
	arr[0] = 1000
}

func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	s1 := arr[2:6]
	fmt.Println(s1)

	s2 := arr[:6]
	fmt.Println(s2)

	s3 := arr[2:]
	fmt.Println(s3)

	s4 := arr[:]
	fmt.Println(s4)

	s5 := make([]int, 5, 10)
	fmt.Printf("s5=%v len=%d cap=%d\n", s5, len(s5), cap(s5))

	s6 := arr[2:6]
	fmt.Printf("arr=%v s6=%v len=%d cap=%d\n", arr, s6, len(s6), cap(s6))

	s7 := s6[3:6]
	fmt.Printf("s7=%v len=%d cap=%d\n", s7, len(s7), cap(s7))

	changeEle(s7)
	fmt.Println(s7, s6, arr)
}
