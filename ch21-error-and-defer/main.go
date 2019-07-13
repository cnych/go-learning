package main

import (
	"ch21-error-and-defer/file"
	"errors"
	"fmt"
)

type Divide struct {
	dividee int
	divider int
}

// error interface
func (d *Divide) Error() string {
	strFormat := `
	Cannot proceed, the divider is zero.
	dividee: %d
	divider: 0
	`
	return fmt.Sprintf(strFormat, d.dividee)
}

func computeDiv(d *Divide) (result int, err error) {
	if d.divider == 0 {
		err = d
	} else {
		result = d.dividee / d.divider
	}
	return
}

func main() {
	err := errors.New("a new err object")
	fmt.Printf("%v\n", err)

	err = fmt.Errorf("a fmt error format object: %s", err.Error())
	fmt.Printf("%v\n", err)

	de := Divide{100, 0}
	if result, err := computeDiv(&de); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(result)
	}

	if content, err := file.ReadFile("abc1.txt"); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(content)
	}

}
