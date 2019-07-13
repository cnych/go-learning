package main

import (
	"fmt"
	"sort"
)

func main() {
	var m map[string]int //m==nil
	fmt.Println(m, m == nil)
	m = make(map[string]int)
	m["a"] = 10
	fmt.Println(m)

	m1 := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
	}
	fmt.Println(m1, m1["b"], m1["d"], m1["e"])
	if val, ok := m1["e"]; ok { // 判断key是否存在
		fmt.Printf("key e in map m1 and value %d\n", val)
	} else {
		fmt.Println("key e not in map m1")
	}
	fmt.Println("===========")

	for key := range m1 {
		fmt.Printf("key=%s val=%d\n", key, m1[key])
	}
	fmt.Println("===========")

	for key, val := range m1 {
		fmt.Printf("key=%s val=%d\n", key, val)
	}
	fmt.Println("===========")

	// delete(m1, "b")

	// for key, val := range m1 {
	// 	fmt.Printf("key=%s val=%d\n", key, val)
	// }

	var keys []string
	for key := range m1 {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, k := range keys {
		fmt.Printf("key=%s, val=%d\n", k, m1[k])
	}

}
