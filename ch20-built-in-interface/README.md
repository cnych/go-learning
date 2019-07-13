## 常用系统接口

接口方法定义：

* 指针方法可以通过指针调用
* 值方法可以通过值调用
* 接收者是值的方法可以通过指针调用，因为指针会首先被解引用
* 接收者是指针的方法不可以通过值调用，因为存储在接口中的值没有地址

常用的系统接口：
```go
type Stringer interface {
    String() string
}

type Interface interface {
    Len() int
    Less(i, j int) bool
    Swap(i, j int)
}

type Reader interface {
    Read(p []byte) (n int, err error)
}
type Writer interface {
    Write(p []byte) (n int, err error)
}
```
