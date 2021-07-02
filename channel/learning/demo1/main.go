package main

import (
	"fmt"
	"os"
	"runtime"
	"time"
)

// 通道的概念
func main()  {
	test2()

	ch := make(chan interface{}, 0)
	//ch := make(chan interface{})

	go func() {
		time.Sleep(2 * time.Second)
		ch <- 1
		close(ch)
	}()

	time.Sleep(3 * time.Second)

	fmt.Println(<-ch)
}

// main函数是不是协程
func test1()  {
	fmt.Println(runtime.NumGoroutine())
	os.Exit(0)
}

// 开启协程
func test2()  {
	arr := make([]int, 0)
	for i := 0; i < 1000; i++ {
		arr = append(arr, i)
	}

	ch := make(chan interface{})

	go test3(arr, ch)

	<- ch // 阻塞
	//time.Sleep(2 * time.Second)
	os.Exit(0)
}

func test3(arr []int, ch chan interface{})  {
	for i, _ := range arr {
		fmt.Println(arr[i])
	}

	ch <- 1
}