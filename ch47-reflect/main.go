package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type User struct {
	Name string `json:"name" bson:"b_name"`
	Age int `json:"age" bson:"b_age"`
}

func (u User) String(prefix string)  {
	fmt.Printf("Prefix=%s, Name=%s, Age=%d\n", prefix, u.Name, u.Age)
}

func (u User) Print() {
	fmt.Println("Hello Reflect")
}

//Int Bool\Int8 Int32  Func Chan Map Array Slice String  Struct

func main() {
	//reflectBaisc()
	//reflectLoopStruct()
	//reflectChangeFiledValue()
	//reflectDynamicCallMethod()
	reflectStructTag()
}

// 反射的基本用法
func reflectBaisc() {
	u := User{"cnych", 30}

	//<Value,Type>
	t := reflect.TypeOf(u)
	fmt.Printf("TypeOf(u)=%v\n", t)

	v := reflect.ValueOf(u)
	t0 := v.Type()
	fmt.Printf("reflect.Value(u)=%v, reflect.Type(u)=%v\n", v, t0)

	fmt.Printf("%T, %v\n", u, u)

	fmt.Println("==============")
	// reflect.Value转成原始数据
	u1 := v.Interface().(User)
	fmt.Println(u1, reflect.TypeOf(u1))

	fmt.Println("==============")
	// 获取底层的数据类型
	fmt.Println(t0.Kind())
}

// 循环获取结构体的熟悉和方法
func reflectLoopStruct() {
	u := User{"张三", 20}

	t := reflect.TypeOf(u)
	for i :=0; i < t.NumField(); i++ {
		f := t.Field(i)
		fmt.Printf("fieldIndex: %d, fieldName: %s\n", f.Index, f.Name)
	}

	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		fmt.Printf("methodIndex: %d, methodName: %s\n", m.Index, m.Name)
	}
}

// 修改字段的值
func reflectChangeFiledValue()  {
	 x := 2
	 x = 50
	 fmt.Println(x)
	 //获得x的reflect.Value
	 v := reflect.ValueOf(&x)
	 v.Elem().SetInt(100)
	 fmt.Println(x)
}

// 动态调用方法
func reflectDynamicCallMethod() {
	u := User{"优点知识", 20}

	v := reflect.ValueOf(u)

	printM := v.MethodByName("String")
	if printM.IsValid() {
		args := []reflect.Value{reflect.ValueOf("PrintPrefix")}
		fmt.Println(printM.Call(args))
	}
}

// struct tag
func reflectStructTag() {
	var u User
	h := `{"name": "cnych", "age": 20}`
	err := json.Unmarshal([]byte(h), &u)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(u, u.Name, u.Age)
	}

	t := reflect.TypeOf(u)
	for i := 0; i < t.NumField(); i ++ {
		f := t.Field(i)
		fmt.Println(f.Tag, f.Tag.Get("json"), f.Tag.Get("bson"))
	}
}