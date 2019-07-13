## 空接口和类型判断

类型断言和判断：
```go
if v, ok := varI.(T); ok {
    // varI is type T
    Process(v)
}
// varI is not of type T

switch t := varI.(type) {
case *xxx:
    fmt.Printf("Type XXX %T with value %v\n", t, t)
default:
    fmt.Printf("Unexpected type %T\n", t)
}
```

> varI 必须是一个接口变量；应该总是使用该方式进行类型断言；接口变量的类型也可以使用 type-swtich 来检测

空接口：
```go
type Any interface {}
```

> 不包含任何方法的空接口，类似于Java中的Object基类，可以表示任何类型。
