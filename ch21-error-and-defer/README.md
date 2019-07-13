## 错误处理与defer语句

错误处理：
```go
type error interface {
    Error() string
}

if value, err := Function1(param1); err != nil {
    fmt.Printf("An error occured in Function1 with parameter %v", param1)
    return err
}
// 未发生错误，继续执行
```

defer语句：
```go
func ReadFile(filename string) (string, error) {
    file, err := os.Open(filename)
    defer file.Close()
    defer fmt.Println(1)
    defer fmt.Println(2)
    if err != nil {
        return "", err
    }
    bt, err := ioutil.ReadAll(file)
    if err != nil {
        return "", err
    }
    return string(bt), nil
}
```

> `defer`的内容在当前函数结束之后执行（如果有return，就是在return之前执行）;多个`defer`函数是一个栈的形式，先进后出。
