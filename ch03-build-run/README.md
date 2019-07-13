## 构建与运行

第一个golang程序：
```go
package main

import "fmt"

func main() {
    fmt.Print("1.Hello World\n")
    fmt.Println("2.Hello World")
    fmt.Println("3.Hello World")
}
```

构建运行 golang 程序方法：

* go build：编译指定的源文件或代码包以及依赖包
* go install：安装自身包和依赖包
* go run：编译并运行 go 程序
* vscode 插件：`Code Runner`
* GoLand IDE
