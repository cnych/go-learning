package main

import (
	"fmt"
	"github.com/cnych/stardust/stringsx"
	stringsV2 "github.com/cnych/stardust/v2/stringsx"
)

func main() {
	fmt.Println(stringsx.Hello("cnych"))

	if greet, err := stringsV2.Hello("cnych", "zh"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(greet)
	}

}

