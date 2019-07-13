## Go Modules

自从 Go 官方从去年推出 1.11 之后，增加新的依赖管理模块并且更加易于管理项目中所需要的模块。模块是存储在文件树中的 Go 包的集合，其根目录中包含 go.mod 文件。 go.mod 文件定义了模块的模块路径，它也是用于根目录的导入路径，以及它的依赖性要求。每个依赖性要求都被写为模块路径和特定语义版本。

* `export GO111MODULE=on`
* 项目可以放置在任何地方，不用在 GOPATH 下面
* `go mod init 模块名`
* `go mod tidy`（自动清理不需要的依赖，同时将依赖项更新到当前版本）
* `export GOPROXY=https://goproxy.io`（配置代理）
