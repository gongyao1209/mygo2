package cancelsignal

import (
	"context"
	"fmt"
	"time"
)

// ======================= 普通关闭 ===========================
// go程一直在进行，消耗资源
func longtimeCostFunc(c chan<- int) {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		fmt.Println("calculating...")
	}
	c <- 1
	fmt.Println("calculate finished")
}

func NormalCancel() {
	ch := make(chan int)

	go longtimeCostFunc(ch)

	select {
	case r := <-ch:
		fmt.Println("longtimeCostFunc return:", r)
	case <-time.After(3 * time.Second):
		fmt.Println("too long to wait for the result")
	}

	time.Sleep(20 * time.Second)

	return
}

// ======================使用 context ============================

func longtimeCostFunc2(ctx context.Context, c chan<- int) {
	defer close(c)

	for i := 0; i < 10; i++ {
		select {
		case <-ctx.Done():
			fmt.Println("calculate interrupted")
			return
		case <-time.After(time.Second):
			fmt.Println("calculating...")
		}
	}

	c <- 1
	fmt.Println("calculate finished")
}

func NormalCancel2() {
	ch := make(chan int)

	ctx, cancel := context.WithCancel(context.Background())

	go longtimeCostFunc2(ctx, ch)

	select {
	case r := <-ch:
		fmt.Println("longtimeCostFunc return:", r)
	case <-time.After(3 * time.Second):
		cancel()
		fmt.Println("too long to wait for the result")
	}

	time.Sleep(10 * time.Second)

	return
}

// =========================普通关闭信号==========================

func longtimeCostFunc3(is_cancel *bool, c chan<- int) {
	defer close(c)

	for i := 0; i < 10; i++ {
		if *is_cancel {
			fmt.Println("calculate interrupted")
			return
		}
		select {
		case <-time.After(time.Second):
			fmt.Println("calculating...")
		}
	}

	c <- 1
	fmt.Println("calculate finished")
}

func NormalCancel3() {
	ch := make(chan int)
	is_close := false

	go longtimeCostFunc3(&is_close, ch)

	select {
	case r := <-ch:
		fmt.Println("longtimeCostFunc return:", r)
	case <-time.After(3 * time.Second):
		is_close = true
		fmt.Println("too long to wait for the result")
	}

	time.Sleep(10 * time.Second)

	return
}

// ======

func longtimeCostFunc4(b chan bool, c chan<- int) {
	defer close(c)

	for i := 0; i < 10; i++ {
		select {
		case <-b:
			fmt.Println("calculate interrupted")
			return
		case <-time.After(time.Second):
			fmt.Println("calculating...")
		}
	}

	c <- 1
	fmt.Println("calculate finished")
}

func NormalCancel4() {
	ch := make(chan int)
	b := make(chan bool)

	go longtimeCostFunc4(b, ch)

	select {
	case r := <-ch:
		fmt.Println("longtimeCostFunc return:", r)
	case <-time.After(3 * time.Second):
		b <- true
		fmt.Println("too long to wait for the result")
	}

	time.Sleep(10 * time.Second)

	return
}
