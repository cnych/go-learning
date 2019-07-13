package main

import (
	"fmt"
)

func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s1 := arr[2:6]
	fmt.Println(s1, len(s1), cap(s1), arr)
	s1 = append(s1, 10, 20, 30, 40, 50)
	fmt.Println(s1, len(s1), cap(s1), arr)

	var s2 []int // nil
	fmt.Println(s2, s2 == nil, len(s2), cap(s2))
	for i := 0; i < 10; i++ {
		s2 = append(s2, i)
		fmt.Println(s2, len(s2), cap(s2))
	}

	// 8<1024 16
	// 1025: 1025+1025/4+1025/4
	// 如果原有的切片容量小于1024，那么扩容就是2倍的容量进行扩充
	// 如果>=1024，在原有基础上增加1/4的容量
	// s3 := make([]int, 1023)
	// fmt.Println(s3, len(s3), cap(s3)) // 1023 1024
	// s3 = append(s3, 200)
	// fmt.Println(s3, len(s3), cap(s3)) // 2048 1024+1024/4=1280

	s4 := []int{1, 2, 3, 4}
	s5 := make([]int, 10)
	fmt.Println(s4, s5)
	copy(s5, s4)
	fmt.Println(s4, s5)

	s6 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("before delete slice s6=%v\n", s6)
	s6 = append(s6[:4], s6[5:]...)
	fmt.Printf("after delete slice s6=%v\n", s6)

	s6 = s6[1:]
	fmt.Printf("after delete first element slice s6=%v\n", s6)

	s6 = s6[:len(s6)-1]
	fmt.Printf("after delete tail element slice s6=%v\n", s6)
}
