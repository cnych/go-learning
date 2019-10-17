package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"regexp"
	"time"
)

func timeDemo() {
	//获取当前时间
	now := time.Now()
	fmt.Println(now)
	//格式化时间
	fmt.Println(now.Format("2006-01-02 15:04:05"))
	fmt.Println(now.Year(), now.Month(), now.Day(),
		now.Hour(), now.Minute(), now.Second())

	prev := time.Date(2010, 10,10,17,20,34,12340, time.UTC)
	fmt.Println(prev.Before(now))
	fmt.Println(prev.After(now))
	fmt.Println(prev.Equal(now))

	fmt.Println(now.Sub(prev))
}

func randDemo() {
	// 0<=x<=100 得随机数
	fmt.Println(rand.Intn(100))
}


type BaseResponse struct {
	Code int `json:"code"`
	Data ResponseData `json:"data"`
}
//{
//	"code": 1,
//	"data" {
//		"name": "xxx",
//		"age": 20
//	}
//}
type ResponseData struct {
	Name string `json:"name"`
	Age int `json:"age"`
}

func jsonDemo() {
	br := BaseResponse{
		Code: 1,
		Data: ResponseData{
			Name: "cnych",
			Age: 20,
		},
	}
	jsonBytes, _ := json.Marshal(&br)
	fmt.Println(string(jsonBytes))

	var br2 BaseResponse
	_ = json.Unmarshal(jsonBytes, &br2)
	fmt.Println(br2.Code, br2.Data.Name, br2.Data.Age)
}

func regexDemo() {
	input := "My email is icnych@gmail.com xxx@abc.com yyy@bcd.com.cn"
	exp, _ := regexp.Compile("([a-zA-Z0-9]+)@[a-zA-Z0-9]+.[a-zA-Z0-9]+")
	fmt.Println(exp.FindString(input))
	fmt.Println(exp.FindAllString(input, -1))

	//fmt.Println(exp.FindAllStringSubmatch(input, -1))
	for _, subMatch := range exp.FindAllStringSubmatch(input, -1) {
		fmt.Println(subMatch[1])
	}

	exp2 := regexp.MustCompile("([a-zA-Z0-9]+)@[a-zA-Z0-9]+.[a-zA-Z0-9]+")
	fmt.Println(exp2.FindAllString(input, -1))
}

// init函数是最先执行得
func init() {
	rand.Seed(time.Now().Unix())
}

func main() {
	timeDemo()

	fmt.Println("============")

	randDemo()
	
	fmt.Println("============")

	jsonDemo()

	fmt.Println("============")
	regexDemo()

}
