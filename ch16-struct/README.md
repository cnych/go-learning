## 结构体

```go
type Book struct {
    id      int
    title   string
    author  string
    subject string
}

book := new(Book)

func (book *Book) String() string {
    return fmt.Sprintf("id=%d,title=%s,author=%s,subject=%s",
        book.id, book.title,
        book.author, book.subject)
}
```

> go语言不支持继承和多态，但是有接口的概念；使用`new`函数给一个新的结构体变量分配内存，它返回指向已分配内存的指针；给结构体定义的函数叫方法，前面需要带上结构体指针的声明。
