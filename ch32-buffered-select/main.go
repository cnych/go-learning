package main

import (
	"fmt"
	"time"
)

func bufferedChannelDemo() {
	c := make(chan int, 3)
	go func() {
		for {
			fmt.Println(<-c)
		}
	}()
	c <- 10
	c <- 20
	c <- 30
	c <- 40
}

func selectDemo() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	timeout := make(chan bool)

	go func() {
		time.Sleep(2 * time.Millisecond)
		timeout <- true
	}()

	go func() {
		ch1 <- 10
		ch2 <- 20
	}()

	// 可以使用select语句和default case来实现非阻塞
	for {
		select {
		case <-ch1:
			fmt.Println("receive data from ch1")
		case <-ch2:
			fmt.Println("receive data from ch2")
		case <- timeout:
			fmt.Println("timeout")
			return
		default:
			fmt.Println("no data")
		}
	}
}

func checkChannelFull(c chan int) {
	select {
	case c <- 2:
	default:
		fmt.Println("channel is full")
	}
}

func main() {
	//bufferedChannelDemo()
	//selectDemo()

	c := make(chan int, 2)
	c <- 1
	c <- 2
	checkChannelFull(c)

	time.Sleep(time.Millisecond)
	//<- time.After(time.Millisecond)
}
