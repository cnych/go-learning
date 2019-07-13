## 内嵌结构体

带标签的结构体：
```go
type Book struct { // 标签(tag)
    id      int    "书籍编号"
    title   string "书籍标题"
    author  string "书籍作者"
    subject string "书籍主题"
}
```

> 结构体中的字段除了有名字和类型外，还可以有一个可选的标签（tag）

匿名字段和内嵌结构体：
```go
type TechBook struct {
    cat  string
    int  // 匿名字段
    Book // 匿名字段
}
```

> 结构体可以包含一个或多个匿名（或内嵌）字段；匿名类型的可见方法也同样被内嵌，这在效果上等同于继承。
