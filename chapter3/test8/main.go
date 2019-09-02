package main

import (
	"fmt"
	"sync"
	"time"
)

func main()  {
	mf := myFilter{
		ch1: make(chan string),
		ch2: make(chan string),
		ch3: make(chan string),
		clo: make(chan struct{}),
	}

	go func() {
		//defer func() {
		//	close(mf.clo)
		//}()
		for i := 0; i < 5; i++ {
			mf.ch1 <- "aaa"
		}

		time.Sleep(1 * time.Second)
		//
		close(mf.clo)
	}()

	a := mf.guolv()

	fmt.Println(a)
}


type myFilter struct {
	ch1 chan string
	ch2 chan string
	ch3 chan string
	clo chan struct{}
}

func (f *myFilter)guolv() []string {

	arr := make([]string, 1)

	wg := sync.WaitGroup{}
	wg.Add(3)

	// 第一层
	go func() {
		defer wg.Done()

		for true {
			select {
			case <- f.clo:
				return
			default:

			}
			select {
			case <- f.clo:
				return
			case v :=<-f.ch1:
				f.ch2 <- v
			}
		}
	}()

	// 第二层
	go func() {
		defer wg.Done()

		for true {
			select {
			case <- f.clo:
				return
			default:

			}
			select {
			case <- f.clo:
				return
			case v :=<-f.ch2:
				f.ch3 <- v
			}
		}
	}()

	go func() {
		defer wg.Done()

		for true {
			select {
			case <- f.clo:
				return
			default:

			}
			select {
			case <- f.clo:
				return
			case v :=<-f.ch3:
				arr = append(arr, v)
			}
		}
	}()

	time.Sleep(1 * time.Second)
	wg.Wait()

	return arr
}