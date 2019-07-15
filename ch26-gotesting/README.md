## 单元测试：

* 测试程序必须属于被测试的包，文件名必须是这种形式`*_test.go`
* 测试程序不会被普通的 Go 编译器编译。
* 测试文件中必须导入`"testing"`包， `TestZzz`开头的全局函数
*表格驱动测试

```
func TestAdd(t *testing.T) {
	inputs:=[]struct{a,b,c int}{
		{1,2,3},
		{4,5,9},
		{10,20,30},
		{3001,4002,7003},
	}
	for _, data := range inputs {
		if result := Add(data.a, data.b); result != data.c {
			t.Errorf("Add(%d,%d) expected result=%d, actual result=%d\n",
				data.a, data.b, data.c, result)
		}
	}
}
```