package main

import (
	"ch20-built-in-interface/example"
	"fmt"
)

type Appender interface {
	Append(int)
}

type Lener interface {
	Len() int
}

type List []int

func (l *List) Len() int {
	return len(*l)
}

func (l *List) Append(val int) {
	*l = append(*l, val)
}

func Counter(a Appender, start, end int) {
	for i := start; i <= end; i++ {
		a.Append(i)
	}
}

func IsLarge(l Lener) bool {
	return l.Len() > 50
}

func main() {
	// var lst List
	// Counter(lst, 1, 10)
	// IsLarge(lst)

	// pointer
	plst := new(List)
	Counter(plst, 1, 10)
	fmt.Println(plst)

	// IsLarge(plst)

	course := new(example.Course)
	course.Title = "golang"
	course.SubTitle = "golang实战课程"
	fmt.Println(course)
	fmt.Printf("%v\n", course)

	// 定义一个slice
	data := []int{23, 50, 78, 11, 19, 60, 0, -100, 1000}
	// 强制转换成IntArray类型
	ia := example.IntArray(data)
	// 调用排序函数
	example.Sort(ia)
	fmt.Println(ia)
}
