## panic和recover

```go
func tryPanic() {
    defer func() {
        if e := recover(); e != nil {
            fmt.Printf("catch panic from recover: %s\n", e)
        }
    }()
    panic("call a panic")
}
```

> panic停止当前函数的执行，一直向上返回，执行每一层的defer；如果没有遇见recover，则退出程序；recover只能在defer调用中使用，获取panic的值。
