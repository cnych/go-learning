## 切片的定义

生成切片：
```go
var s0 []int // s0 = []int{1, 2, 3}
s1 := make([]int, 5, 5)

arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
s2 := arr[2:6]
s3 := arr[2:]
s4 := arr[:6]
s5 := arr[:]
s6 := s4[3:5]
```

> 初始化切片，[]表示是切片类型，不需要指定大小；也可以通过make关键字进行初始化，`make([]T, length, capacity)`；指定数组的索引初始化切片，是数组的引用；切片之上还可以继续指定切片。

切片的数据结构：
```go
type slice struct {
    array unsafe.Pointer
    len   int
    cap   int
}
```
![go slice struct](https://bxdc-static.oss-cn-beijing.aliyuncs.com/images/go-slice-struct.jpg) (图片来源于网络)
