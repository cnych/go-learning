package main

import (
	"fmt"
	"time"
)


func getData(i int, c chan int) {
	for n := range c {
		fmt.Printf("get data %d from channel %d\n", n, i)
		//n := <-c
		//if n,ok:= <-c; ok {
		//	fmt.Printf("get data %d from channel %d\n", n, i)
		//} else {
		//	break
		//}
	}
}

func createChannel(i int) chan<- int {
	c := make(chan int)
	go getData(i, c)
	return c
}

func channelDemo() {
	//var c chan int // c=nil
	var channels [10]chan<- int
	for i:=0;i<10;i++ {
		channels[i] = createChannel(i)
	}

	for i:=0;i<10;i++ {
		channels[i] <- i
	}

	for i:=0;i<10;i++ {
		channels[i] <- i + 100
	}

	for i:=0;i<10;i++ {
		close(channels[i])
	}

	//channels[0] <- 1111

	//c := createChannel()
	//go getData(c)
	//c <- 10  // 发送者
	//c <- 20
	time.Sleep(time.Millisecond)
}

func main() {
	channelDemo()
}
