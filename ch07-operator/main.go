package main

import "fmt"

func operatorDemo() {
	var a int = 21
	var b int = 10
	var c int

	c = a + b
	// a += b  // a = a + b
	fmt.Printf("a+b=%d\n", c)

	c = a - b
	// a -= b  // a = a-b
	fmt.Printf("a-b=%d\n", c)

	c = a * b
	fmt.Printf("a*b=%d\n", c)

	c = a / b
	fmt.Printf("a/b=%d\n", c)

	c = a % b
	fmt.Printf("a%%b=%d\n", c)

	a++ // a = a + 1
	fmt.Printf("a=%d\n", a)

	a = 21

	a-- // a = a - 1
	fmt.Printf("a=%d\n", a)
}

func relationDemo() {
	var a int = 21
	var b int = 10

	if a == b {
		fmt.Println("a == b")
	} else {
		fmt.Println("a != b")
	}

	if a > b { // true
		fmt.Println("a > b")
	} else {
		fmt.Println("a <= b")
	}

	// a = 5
	if a < b { // false
		fmt.Println("a < b")
	} else {
		fmt.Println("a >= b")
	}
}

func logicDemo() {
	var a bool = true  // 真
	var b bool = false // 假

	if !(a && b) { // 如果a和b都是真  // false
		fmt.Println("xxzvzxc")
	} else {
		fmt.Println("a和b有假")
	}

	if a || b { // 如果a或者b中有一个为真 // true
		fmt.Println("a或者b至少有一个真")
	} else {
		fmt.Println("a和b都为假")
	}

	if !b { // a 为真，非真即为假
		fmt.Println("b为假")
	} else {
		fmt.Println("b为真")
	}
}

func byteComputeDemo() {
	var a uint = 60 // 60=111100
	// fmt.Printf("%b\n", a)
	var b uint = 13 // 13=1101
	// fmt.Printf("%b\n", b)
	var c uint

	// a=111100
	// b=001101
	c = a & b // 12
	fmt.Printf("a&b=%d\n", c)
	//    c = 1100

	c = a | b // 111101=61
	// a=111100  0
	// b=001101  1
	fmt.Printf("a|b=%d\n", c)

	c = a ^ b // 110001=1*1 + 0 + 16+32= 49
	// a=111100
	// b=001101
	fmt.Printf("a^b=%d\n", c)

	c = a << 2 // 60 * 4 = 240
	// a << n 左移n位=乘以2的n次方
	fmt.Printf("a<<2=%d\n", c)
	// <<

	// >>
	c = a >> 2 // 60 / 4 = 15
	// a >> n 右移n位=除以2的n次方

	a >>= 2 // a = a >> 2
	fmt.Printf("a>>2=%d\n", a)
}

func main() {
	operatorDemo()
	fmt.Println("---------------")
	relationDemo()
	fmt.Println("---------------")
	logicDemo()
	fmt.Println("---------------")
	byteComputeDemo()
}
