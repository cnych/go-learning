## 接口和duck typing

duck typing：

* 当一只鸟走起来像鸭子、游泳起来像鸭子、叫起来也像鸭子，那么这只鸟就可以被称为鸭子
* 我们并不关心对象是什么类型，到底是不是鸭子，只关心行为（描述事物的外部行为，而非内部结构）

接口：
```go
type Phone interface {
    Call()
}

type MiPhone struct {
}

func (mp *MiPhone) Call() {
    fmt.Println("I am iPhone, I can call you!")
}
```

> 接口把所有的具有共性的方法定义在一起；任何其他类型只要实现了这些方法就是实现了这个接口。
