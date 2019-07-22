package main

import (
	"fmt"
	"sync"
)

type safeInt struct {
	a int
	lock sync.Mutex
}

func (si *safeInt) increase() {
	si.lock.Lock()
	defer si.lock.Unlock()
	si.a++
}

func (si *safeInt) get() int {
	si.lock.Lock()
	defer si.lock.Unlock()
	return si.a
}

func main() {
	var si safeInt
	si.increase()

	var wg sync.WaitGroup

	for i:=0;i<100000;i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			si.increase()
		}()
	}

	wg.Wait()
	//time.Sleep(time.Microsecond)
	fmt.Println(si.get())
	//var a int
	//a++
	//go func() {
	//	a++
	//}()
	//time.Sleep(time.Millisecond)
	//fmt.Println(a)
}
