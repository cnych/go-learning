## 循环语句

```go
sum := 0
for i := 1; i <= 100; i++ {
    sum += i
}

n := 0
for n < 10 {
    n++
}

for {
    fmt.Println("dead loop")
}
```

> `for` 语句不需要括号包裹起来；没有初始值，相当于 while 循环；没有初始值，没有循环条件，表示死循环

循环控制语句：

* break语句：用于中断当前 for 循环
* continue语句：跳过当前循环的剩余语句，然后继续进行下一轮循环
* goto语句：将控制转移到被标记的语句
