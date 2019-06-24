package main

import "fmt"

// 全局变量
var g int = 100

// 形式参数
func sum(a, b int) int {
	// c 局部变量
	var c = 10
	g := 0
	s := a + b + c + g
	return s
}

// 值传递
func funcValRef(a int) {
	a = 1000
	fmt.Printf("in func inner: %d\n", a)
}

// 值传递
func funcValRefPtr(a *int) {
	*a = 1000
	fmt.Printf("in func inner: %d\n", *a)
}

func main() {
	// 局部变量
	var a, b, c int

	g := 200

	a = 10
	b = 20
	c = a + b + g
	fmt.Println(a, b, c)

	fmt.Println("===========")

	var p int = 100
	funcValRef(p)
	fmt.Println(p)

	funcValRefPtr(&p)
	fmt.Println(p)
}
