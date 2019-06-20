package main

import "fmt"

func main() {
	a := 23
	b := 2

	if a < 20 { // true false
		fmt.Println("a < 20")
	} else if a == 20 { // false
		fmt.Println("a = 20")
	} else { // true
		fmt.Println("a > 20")
		if b > 5 { // false
			fmt.Println("b > 5")
		} else {
			fmt.Println("b <= 5")
		}
	}
	fmt.Println("------------")

	switch a {
	case 20:
		fmt.Println("a=20")
	case 21:
		fmt.Println("a=21")
	case 22, 23, 24:
		fmt.Println("a in [22, 23, 24]")
	default:
		fmt.Println("Unknown")
	}

	switch {
	case b == 2:
		fmt.Println("b = 2")
		fallthrough
	case a > 20:
		fmt.Println("a > 20")
		fallthrough
	case a == 20:
		fmt.Println("a=20")
	case a < 20: // true false
		fmt.Println("a < 20")
	default:
		fmt.Println("Unknown")
	}

	s := []int{1, 2, 3, 4, 5, 6}

	fmt.Println(s[1:4])
}
