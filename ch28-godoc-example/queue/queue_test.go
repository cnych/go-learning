package queue

import "fmt"

func ExampleQueue_Push() {
	q := Queue{1}
	q.Push(2)
	q.Push("a")
	q.Push(true)

	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())

	// output:
	// 1
	// 2
	// a
	// true
}
