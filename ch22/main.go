package main

import "fmt"

func tryPanic() int {
	// defer func() {
	// 	if e := recover(); e != nil {
	// 		fmt.Printf("catch panic in recover function: %v\n", e)
	// 	}
	// }()

	fmt.Println("first line in tryPanic function")
	panic("call a panic")
	// a := 0
	// b := 100
	// r := b / a
	// fmt.Println(r)
	fmt.Println("second line in tryPanic function")
	return 0
}

func tryPanic2() {
	// defer func() {
	// 	if e := recover(); e != nil {
	// 		fmt.Printf("catch panic in recover function: %v\n", e)
	// 	}
	// }()

	fmt.Println("first line in tryPanic2 function")
	panic("call a panic2")
	// a := 0
	// b := 100
	// r := b / a
	// fmt.Println(r)
	fmt.Println("second line in tryPanic2 function")
}

// type panicFunc func()

func protect(g func() int) {
	defer func() {
		if e := recover(); e != nil {
			fmt.Printf("catch panic in recover function: %v\n", e)
		}
	}()
	g()
}

func main() {
	fmt.Println("start call tryPanic")
	// tryPanic()
	protect(tryPanic)
	fmt.Println("end call tryPanic")
}
