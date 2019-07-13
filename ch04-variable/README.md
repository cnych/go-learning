## 变量

vscode go 插件配置：

* "cmd+shift+p": “go:install/update tools” 安装插件
* "go.toolsGopath": “指定tools包的路径，不指定默认在 GOPATH”
```json
{
    "go.gopath": "${workspaceFolder}",
    "go.inferGopath": true,
    "go.autocompleteUnimportedPackages": true,
    "go.gocodePackageLookupMode": "go",
    "go.gotoSymbol.includeImports": true,
    "go.useCodeSnippetsOnFunctionSuggest": true,
    "go.useCodeSnippetsOnFunctionSuggestWithoutType": true,
    "go.docsTool": "gogetdoc"
}
```

变量声明：
```go
var a int、var a, b int = 1, 2
var a, b, c = 1, "go", true
a, b, c := 1, "go", true
var (
    a = 1
    b = "go"
    c = true
)
```

> `:=`只能在函数内使用，不能用在函数外