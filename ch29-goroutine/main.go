package main

import (
	"fmt"
	"time"
)

func printGroutine(i int) {
	fmt.Printf("I am from groutine %d", i)
}

func main() {
	var a [10]int
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Printf("I am from groutine %d\n", i)
		}(i)
		//go func(i int) {
		//	for {
		//		a[i]++
		//		//fmt.Printf("I am from groutine %d\n", i)
		//		runtime.Gosched()
		//	}
		//}(i)
	}
	time.Sleep(time.Millisecond)
	fmt.Println(a)

}

