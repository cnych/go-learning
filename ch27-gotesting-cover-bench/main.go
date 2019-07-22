package main

import (
	"ch27-gotesting-cover-bench/calculate"
	"fmt"
)

func main() {
	for i := 0; i <= 100; i++ {
		fmt.Printf("%d is even? %v\n", i, calculate.Even(i))
	}
}
