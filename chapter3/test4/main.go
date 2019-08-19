package main

import (
	"fmt"
	"sync"
	"time"
)

type T interface {}

// 检查 信道是否关闭

func main()  {

	TestSafe()
}

// --------------------------------------------------

func IsChClosed(ch <-chan T) bool { // 检查信道是否关闭
	select {
	case <-ch: // 这里是为什么呢？
		return true
	default:
	}

	return false
}

// --------------------------------------------------

func SafeSend(ch chan T, value T) (close bool) { // 安全的往信道里面放数据
	defer func() {
		if r := recover(); r != nil { // 抓取异常
			fmt.Println(r)
			close = true
		}
	}()

	ch <- value // panic if ch is closed
	return false
}

func TestSafe()  {
	ch := make(chan T)

	//c := <- ch // 直接阻塞
	//fmt.Println(c)
	go func() {
		i := 1

	End:
		for true {
			var t T = i
			clo := SafeSend(ch, t)
			if clo {
				break End
			}

			i ++
		}
	}()

	go func() {
		for c := range ch {
			fmt.Println("jieguo: ", c)
		}
	}()

	//fmt.Println("1, ", IsChClosed(ch))
	time.Sleep(1 * time.Millisecond)

	close(ch)

	time.Sleep(1 * time.Second)
	//fmt.Println("2, ",IsChClosed(ch))
}

// --------------------------------------------------

type MyChannel struct {
	ch chan T
	sy sync.Once
}

func NewMyChannel() *MyChannel {
	return &MyChannel{
		ch: make(chan T),
	}
}

func (ch *MyChannel)SafeClose() { // 使用 sync.Once 来关闭
	ch.sy.Do(func() {
			close(ch.ch)
		})
}

// --------------------------------------------------
type MyChannel2 struct {
	c		chan T
	closed	bool
	m		sync.Mutex // 使用锁 来保证单线工作
}

func NewMyChannel2() *MyChannel2 {
	return &MyChannel2{
		c: make(chan T),
	}
}

func (myCh *MyChannel2)SaleClose() {
	myCh.m.Lock() // 锁定
	defer myCh.m.Unlock() // 解锁

	if !myCh.closed {
		close(myCh.c)
		myCh.closed = true
	}
}

func (myCh *MyChannel2)IsClosed() bool {
	myCh.m.Lock() // 锁定
	defer myCh.m.Unlock() // 解锁

	return myCh.closed
}

// --------------------------------------------------

