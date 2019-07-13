## 条件语句

if 语句：
```go
if 布尔表达式 {
    /* 在布尔表达式为 true 时执行 */
}


if 布尔表达式 {
    /* 在布尔表达式为 true 时执行 */
} else if 另外一个布尔表达式 {
    /* 在布尔表达式为 true 时执行 */
} else {
    /* 在布尔表达式为 false 时执行 */
}
```

switch...case 语句：
```go
switch {
case true:
        fmt.Println("1、case 条件语句为 false”)
        fallthrough
case false:
        fmt.Println("2、case 条件语句为 false")
case true:
        fmt.Println("3、case 条件语句为 true")
case true:
        fmt.Println("4、case 条件语句为 true")
default:
        fmt.Println(“5、默认 case”)
}
```

> 匹配到的 case 后面不需要加 break，相当于默认就有 break；默认情况下 case 匹配成功后就不会执行其他 case，如果我们需要执行后面的 case，可以使用 `fallthrough`, `fallthrough` 不会判断下一条 case 的表达式结果是否为 true。
