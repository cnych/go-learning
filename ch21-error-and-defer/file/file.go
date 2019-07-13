package file

import (
	"fmt"
	"io/ioutil"
	"os"
)

func ReadFile(filename string) (string, error) {
	file, err := os.Open(filename)
	defer fmt.Println("first defer function")
	defer file.Close()
	if err != nil {
		panic(err)
	}
	defer fmt.Println("second defer function")
	bts, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}
	return string(bts), nil
}
