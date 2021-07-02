package main

import (
	"fmt"
	"os"
	"time"
)

// 如何优雅的关闭
type myChannel struct {
	Ch chan interface{}
	IsClose bool
}
func (c *myChannel)close()  {
	close(c.Ch)
	c.IsClose = true
}
func getMyChannel() *myChannel {
	return &myChannel{
		Ch:      make(chan interface{}),
		IsClose: false,
	}
}
func getMyNChannel() *myChannel {
	return &myChannel{
		Ch:      make(chan interface{}, 10),
		IsClose: false,
	}
}

// 关闭
func main()  {
	//test1()
	test2()
}

// 管道谁来关闭
func test1()  {
	ch := getMyChannel()

	go func(c *myChannel) {
		for i := 0; i < 1000; i++ {
			c.Ch <- i
		}
	}(ch)

	go func(c *myChannel) {
		for true {
			select {
			case v :=<- c.Ch:
				fmt.Println(v)
			case <-time.After(2*time.Second):
				fmt.Println("time out")
			}
		}
	}(ch)

	ch.close()

	time.Sleep(1 * time.Second)

	os.Exit(0)
}

//
func test2()  {
	//ch := getMyChannel() // 无缓冲通道
	ch := getMyNChannel() // 有缓冲通道

	go func(c *myChannel) {
		for i := 0; i < 1000; i++ {
			c.Ch <- i
		}
		c.close() // 生产者 关闭通道
	}(ch)

	go func(c *myChannel) {
		for true {
			if c.IsClose {
				break
			}

			select {
			case v :=<- c.Ch:
				fmt.Println(v)
			case <-time.After(2*time.Second):
				fmt.Println("time out")
			}
		}
	}(ch)

	time.Sleep(1 * time.Second)

	os.Exit(0)
}