## 常量、枚举、变量定义与别名

常量：
```go
const s string = "Hello"
const a, b = 3, 4
const (
    s1  = "golang"
    c   = 5
    MAX = 10
)
```

> 不指定类型的常量，则它的类型是不确定的,可以当作各种类型使用。

枚举：
```go
const (
    Monday = 1 + iota
    Tuesday
    Wednesday
    Thursday
    Friday
    Staurday
    Sunday
)
```

> 枚举是一种特殊的常量，可以通过iota快速设置连续的值

类型定义与别名：
```
type MyInt1 int
type MyInt2 = int

var i int = 1
var i1 MyInt1 = MyInt1(i)
var i2 MyInt2 = i
```

> 类型定义基于类型创建的一个新类型，主要提高代码可读性；类型别名基于类型创建的一个别名，和原类型完全一样，主要用于包兼容；类型定义是一个新的类型了，所以类型转换的时候必须强制类型转换。