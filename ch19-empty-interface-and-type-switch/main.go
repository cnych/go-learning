package main

import "fmt"

type Phone interface {
	Call(string)
}

type Camera interface {
	Take() string
}

type SmartPhone interface {
	Phone
	Camera
	Download(string) string
}

func ListSmartPhoneFunction(sp SmartPhone) {
	if v, ok := sp.(*Iphone); ok {
		v.Call("10086")
		fmt.Println("sp.Take()", v.Take())
		fmt.Println("sp.Download()", v.Download("https://youdianzhishi.com"))
	} else {
		fmt.Println("Not MiPhone Pointer type")
	}
}

type MiPhone struct {
	Logo string
}

func (mp *MiPhone) Call(phone string) {
	fmt.Println("Call to phone:", phone)
}

func (mp *MiPhone) Take() string {
	return "logo.png"
}

func (mp *MiPhone) Download(url string) string {
	return fmt.Sprintf("visit url:%s", url)
}

type Iphone struct {
	Logo string
}

func (mp *Iphone) Call(phone string) {
	fmt.Println("Call to phone:", phone)
}

func (mp *Iphone) Take() string {
	return "logo.png"
}

func (mp *Iphone) Download(url string) string {
	return fmt.Sprintf("visit url:%s", url)
}

// 空接口，可以用interface{}表示任何的一个数据的类型
type Any interface{}

// type-switch
func GetAnyType(any interface{}) {
	switch t := any.(type) {
	case int:
		fmt.Println("any is int type")
	case string:
		fmt.Println("any is string type")
	case *MiPhone:
		fmt.Println("any is MiPhone pointer type")
	default:
		fmt.Printf("Unexpected type %T\n", t)
	}
}

type Queue []interface{}

func (q *Queue) Push(n interface{}) {
	*q = append(*q, n)
}

func (q *Queue) Pop() interface{} {
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}

func main() {
	mp := new(MiPhone)
	mp.Logo = "xiaomi"

	ip := new(Iphone)
	ip.Logo = "apple"

	ListSmartPhoneFunction(ip)

	var va1 Any
	va1 = 5
	fmt.Printf("va1 value: %v\n", va1)

	GetAnyType(va1)

	str := "ABCD"
	va1 = str
	fmt.Printf("va1 value: %v\n", va1)
	GetAnyType(va1)

	va1 = *mp
	fmt.Printf("va1 value: %v\n", va1)
	GetAnyType(va1)

	q := Queue{1, 2, 3}
	q.Push(4)
	q.Push(5)
	fmt.Println(q)
	q.Pop()
	fmt.Println(q)
	q.Push("a")
	q.Push("b")
	fmt.Println(q)

}
