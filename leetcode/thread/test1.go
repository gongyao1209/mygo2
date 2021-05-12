package main

import (
	"fmt"
	"time"
)

type myChannel struct {
	ch chan interface{}
	isClose bool
}

func NewChannel() *myChannel {
	return &myChannel{
		ch: make(chan interface{}),
		isClose: false,
	}
}

func (c *myChannel)close()  {
	close(c.ch)
	c.isClose = true
}

func main()  {

	f := new(Foo)

	fi_ch := NewChannel()
	se_ch := NewChannel()
	th_ch := NewChannel()

	go func(foo Foo, ch *myChannel) {
		for true {
			if ch.isClose {
				return
			}
			select {
			case v := <- ch.ch:
				if v != nil {
					foo.first()
				}
			case <- time.After(time.Second):
				return
			}
		}
	}(*f, fi_ch)

	go func(foo Foo, ch *myChannel) {
		for true {
			if ch.isClose {
				return
			}
			select {
			case v :=<- ch.ch:
				if v != nil {
					foo.secend()
				}
			case <- time.After(time.Second):
				return
			}
		}
	}(*f, se_ch)

	go func(foo Foo, ch *myChannel) {
		for true {
			if ch.isClose {
				return
			}
			select {
			case v :=<- ch.ch:
				if v != nil {
					foo.thread()
				}
			case <- time.After(time.Second):
				return
			}
		}
	}(*f, th_ch)


	//fi_ch.ch <- 1
	arr := [3]int{2, 3, 1}
	for _, v := range arr {
		if v == 1 {
			fmt.Println(1)
			fi_ch.ch <- 1
			time.Sleep(100*time.Millisecond)
		} else if v == 2 {
			fmt.Println(2)
			se_ch.ch <- 1
			time.Sleep(100*time.Millisecond)
		} else {
			fmt.Println(3)
			th_ch.ch <- 1
			time.Sleep(100*time.Millisecond)
		}
	}

	fi_ch.close()
	se_ch.close()
	th_ch.close()

	time.Sleep(3 * time.Second)

	return
}

type Foo struct {
}

func (f Foo)first()  {
	fmt.Println("first")
}
func (f Foo)secend()  {
	fmt.Println("secend")
}
func (f Foo)thread()  {
	fmt.Println("thread")
}