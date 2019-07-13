package main

import (
	"ch17-built-in-struct/book"
	"fmt"
	"strconv"
)

type Integer int

func (it Integer) String() string {
	return strconv.Itoa(int(it))
}

func main() {
	bk := book.NewBook(1000, "Golang", "James", "About Golang Book")
	fmt.Println(bk.String())

	book.RefTag(*bk, 2)

	book.InitTechBook()

	it := Integer(100)
	str := it.String()
	fmt.Printf("it=%s and type=%T\n", str, str)
}
