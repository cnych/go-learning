package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"path"
)

func sendHttpRequestDemo() {
	//User-Agent: Mozilla/5.0 (iPhone; CPU iPhone OS 11_0 like Mac OS X) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Mobile/15A372 Safari/604.1
	req, err := http.NewRequest(http.MethodGet, "http://www.baidu.com", nil)
	if err != nil {
		panic(err)
	}
	req.Header.Add("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 11_0 like Mac OS X) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Mobile/15A372 Safari/604.1")
	req.Header.Add("Host", "www.baidu.com")

	resp, err := http.DefaultClient.Do(req)
	//resp, err := http.Get("http://www.baidu.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}

type User struct {
	Name string `json:"name"`
	Age int `json:"age"`
}
//{"name": "cnych", "age": 20}

func main() {
	//sendHttpRequestDemo()
	// xxx.com xxxx.com/course xxx.com/course/123
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		io.WriteString(writer, "Hello Golang\n")
	})
	http.HandleFunc("/json", func(writer http.ResponseWriter, request *http.Request) {
		user := User{
			Name: "cnych",
			Age: 30,
		}

		userJSON, err := json.Marshal(user)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		writer.Header().Set("Content-Type", "application/json")
		writer.Write(userJSON)
	})
	http.HandleFunc("/image", func(writer http.ResponseWriter, request *http.Request) {
		image := path.Join("images", "golang.png")
		http.ServeFile(writer, request, image)
	})
	http.HandleFunc("/html", func(writer http.ResponseWriter, request *http.Request) {
		user := User{
			Name: "cnych",
			Age: 30,
		}
		htmlFile := path.Join("templates", "index.html")
		tmpl, err := template.ParseFiles(htmlFile)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		if err := tmpl.Execute(writer, user); err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
		}

	})
	log.Fatal(http.ListenAndServe(":1234", nil))
}
