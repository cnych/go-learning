package main

import (
	"fmt"
)

func changeEle(arr *[3]int) {
	arr[0] = 1000
}

func main() {
	var arr1 [5]int
	arr1[1] = 10
	fmt.Println(arr1, arr1[0], arr1[1], arr1[4])

	arr2 := [3]int{1, 2, 3}
	fmt.Println(arr2)

	arr3 := [...]int{4, 5, 6, 7, 8, 10}
	fmt.Println(len(arr3), arr3)

	for i := 0; i < len(arr3); i++ {
		fmt.Println(arr3[i])
	}

	for i, num := range arr3 {
		fmt.Println(i, num)
	}

	for _, num := range arr3 {
		fmt.Println(num)
	}

	for i := range arr3 {
		fmt.Println(i)
	}

	// var grid [4][6][5]int
	arr4 := [3][4]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
	}
	fmt.Println(arr4)

	arr5 := [3]int{1, 2, 3}
	changeEle(&arr5)
	fmt.Println(arr5)
}
