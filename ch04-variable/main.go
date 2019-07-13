package main

import "fmt"

var c1, c2 = 6, "golang"

// d1, d2 := 7, "go"

var (
	d1 = 7
	d2 = "go"
)

func varInitValue() {
	var a, b int = 3, 4
	fmt.Println(a, b)
	var s = "abcd"
	fmt.Println(s)

	var s1, s2, s3 = 5, "hello", true
	fmt.Println(s1, s2, s3)

	s4, s5, s6 := 5, "world", false
	fmt.Println(s4, s5, s6)
}

func main() {
	var a int
	var b string
	fmt.Printf("%d %q\n", a, b)
	// var a int = 3
	// a = 3
	varInitValue()

	fmt.Println(c1, c2)
	fmt.Println(d1, d2)
}
