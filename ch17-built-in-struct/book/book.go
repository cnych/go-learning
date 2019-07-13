package book

import (
	"fmt"
	"reflect"
)

type book struct {
	Id      int    "this is a book id"
	Title   string "this is a book title"
	Author  string "this is a book author"
	Subject string "this is a book subject"
}

type techBook struct {
	Cat string
	// Title string
	int  // 匿名字段
	book // 匿名字段，内嵌结构体
}

type A struct{ a int }
type B struct{ a, b int }
type C struct {
	A
	B
}

// c.a c.A.a c.B.a

func NewBook(id int, title, author, subject string) *book {
	return &book{id, title, author, subject}
}

func (b *book) String() string {
	return fmt.Sprintf("id=%d,title=%s,author=%s,subject=%s\n",
		b.Id, b.Title, b.Author, b.Subject)
}

func RefTag(b book, idx int) {
	bType := reflect.TypeOf(b)
	ixField := bType.Field(idx)
	fmt.Printf("%v\n", ixField.Tag)
}

func InitTechBook() {
	bk := NewBook(1000, "Golang", "James", "About golang book")
	tb := new(techBook)
	tb.Cat = "tech"
	tb.int = 10
	tb.book = *bk

	fmt.Println("techBook Cat=", tb.Cat)
	fmt.Println("techBook int=", tb.int)
	fmt.Println("techBook Title=", tb.Title)
	fmt.Println("techBook book=", tb.book.String())

	fmt.Println(tb.String())
}
