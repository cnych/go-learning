## 函数

函数的定义：
```go
func operate(a, b int, op string) int

func swap(a, b int) (x, y int)

func compute(op func(int, int) int, a, b int) int

func sum(nums ...int) int
```

函数的特性：

* 函数返回值类型写在最后面
* 函数可以返回多个值，也可以给返回值命名，一般和error结合使用
* 函数式编程，函数也可以作为参数传递给其他函数
* go语言中没有默认参数，可选参数，但是可以使用可变参数列表
