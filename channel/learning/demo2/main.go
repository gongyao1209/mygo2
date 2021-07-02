package main

import (
	"fmt"
	"os"
	"time"
)

// 管道 进行协程间通信
func main()  {
	test1()

	ch := make(chan interface{}, 5)

	// 多路复用
	select {
	case c :=<-ch:
		fmt.Println(c)
	case <-time.After(1 * time.Second):
		fmt.Println("超时")
	}
}

// 开启协程
func test1()  {
	ch := make(chan interface{})
	zusech := make(chan bool)

	go test2(ch, zusech)

	go test3(ch)

	<- zusech // 阻塞
	//time.Sleep(2 * time.Second)
	os.Exit(0)
}

func test2(ch chan interface{}, suze chan bool)  {
	for i := 0; i < 1000; i++ {
		ch <- i
	}

	suze <- true
}

func test3(ch chan interface{})  {
	for c := range ch {
		fmt.Println(c)
	}
}