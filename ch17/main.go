package main

import (
	"ch17/book"
	"fmt"
)

func main() {

	bk := book.NewBook(1000, "Golang", "James", "About Golang Book")
	fmt.Println(bk.String())

	book.RefTag(*bk, 2)

	book.InitTechBook()
}
