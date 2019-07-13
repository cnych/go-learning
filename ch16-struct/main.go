package main

import "fmt"

type Book struct {
	id      int
	title   string
	author  string
	subject string
}

// new New 开头，将结构体编程小写（私有的）可以来强制使用工厂函数来实例化
func NewBook(id int, title, author, subject string) *Book {
	return &Book{id, title, author, subject}
}

func (book *Book) String() string {
	return fmt.Sprintf("id=%d,title=%s,author=%s,subject=%s\n",
		book.id, book.title, book.author, book.subject)
}

func (book *Book) GetTitle() string {
	return book.title
}

func (book *Book) SetTitle(title string) {
	book.title = title
}

func printBook(book *Book) {
	fmt.Printf("id=%d,title=%s,author=%s,subject=%s\n",
		book.id, book.title, book.author, book.subject)
	book.id = 1000
}

func main() {
	var book1 *Book
	book1 = new(Book)
	book1.id = 1001
	book1.title = "go in action"
	book1.author = "james"
	book1.subject = "about golang"
	fmt.Println(book1)

	fmt.Println(book1.String())

	book2 := Book{
		id:      1002,
		title:   "python in action",
		author:  "jordan",
		subject: "about python",
	}
	// book2 := Book{
	// 	1002,
	// 	"python in action",
	// 	"jordan",
	// 	"about python",
	// }
	fmt.Println(book2)

	fmt.Println("book2.title=", book2.title)

	printBook(&book2)
	fmt.Println(book2)
	fmt.Println(book2.String())

	book3 := NewBook(1004, "Java", "gsl", "Java in action")
	fmt.Println(book3.String())
}
