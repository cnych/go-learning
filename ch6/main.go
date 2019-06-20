package main

import (
	"fmt"
	"math"
	// "github.com/cnych/lib2/T1"
)

// package:github.com/cnych/lib2
// type T1 = lib2.T1

// 可导出的MAX常量
const (
	s1  = "golang" // private
	MAX = 10       // public export
)

func typeDefAndAlias() {
	type MyInt1 int   // 基于int定义的一个新的类型，这个用得比较多一点，（别名）
	type MyInt2 = int // MyInt2和int就是完全一样的 type alias

	var i int = 1
	var i1 MyInt1 = MyInt1(i)
	var i2 MyInt2 = i
	fmt.Println(i1, i2)
}

func enumDemo() {
	const (
		Monday = 1 + iota // 1
		Tuesday
		Wednesday
		Thursday
		Friday
		Staurday
		Sunday
	)
	fmt.Println(Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Staurday)
}

func constDemo() {
	const s string = "Hello"
	const a, b = 3, 4
	fmt.Println(s, a, b)

	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(c)

	// s1 = "World"
	fmt.Println(s1, MAX)
}

func main() {
	constDemo()
	enumDemo()
	typeDefAndAlias()
}
