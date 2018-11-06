package main

import (
	"fmt"
	"runtime"
	"time"
)

func main()  {

	//go TestChannel2() //
	TestChannel1()
	TestChannel2() //只有当这个的时候会报错，和main方法一个goroutine，所以是阻塞主进程了，所以才报错
}

func TestChannel1()  {
	ch := make(chan int)
	defer close(ch)

	go func() {
		for c := range ch {
			fmt.Println("ch value : ", c)
		}
	}()

	for i := 0; i < 10; i++ {
		ch <- i
	}

	time.Sleep(1 * time.Second)
}

func TestChannel2()  {
	ch := make(chan int)

	go func() {
		defer close(ch)
		for i := 0; i < 10; i++ {
			ch <- i + 1
		}
	}()

	for c := range ch { //这里可能引起阻塞
		fmt.Println("ch value : ", c)
	}
}

func TestChannel2_1()  {
	ch := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()

	go func() {
		for c := range ch {
			fmt.Println("ch value : ", c)
		}
	}()

	time.Sleep(1 * time.Second)
}

func TestChannel2_2()  {
	ch := make(chan int, 1)

	go func() {
		defer close(ch)
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()

	s := make([]int, 0, 10)

	for {
		select {
		case c, ok :=<- ch:
			if !ok {
				fmt.Println("CLOSE")
				goto End
			}
			s = append(s, c)
			fmt.Println("ch value : ", c)
		case <-time.After(1 * time.Second):
			goto End
		}
	}

	End:
		fmt.Println("end")

	fmt.Println(s)
}

func main1()  {
	var c = make(chan int, 1)

	go func() {
		//time.Sleep(time.Second * 3)
		c <- 1
		c <- 2
		c <- 3
		close(c)
	}()

	fmt.Println("count goroutine: ", runtime.NumGoroutine())

	for  {
		select {
		case v := <- c:
			fmt.Println(v)
			if v == 0 {
				goto Gongyao
			}
		case <-time.After(1 * time.Second):
			goto Gongyao
		}
	}

	Gongyao:
		fmt.Println("gongyao")

	return
}

func rangeArr()  {
	a := [5]int{1, 2, 3, 4, 5}
	var r [5]int
	for key, value := range a {
		if key == 0 {
			a[1] = 12
			a[2] = 13
		}

		r[key] = value
	}

	fmt.Println(r)
}

func getArray() (a [10000000]string) {
	for i := 0; i < 10000000; i++ {
		a[i] = "My Name Is Gongyao"
	}
	return
}

func rangeSlice()  {
	a := []int{1, 2, 3, 4, 5}
	var r [5]int

	for k, v := range a { // 这里也是副本，但是 由于slice的结构不同
		if k == 0 {
			a[1] = 12
			a[2] = 13
		}

		r[k] = v
	}

	fmt.Println(r)
}

func rangeMap()  {
	m := map[string]interface{} {
		"name":"wangliang",
		"age":26,
		"home":"shanxi",
		"town":"taiyuan",
	}

	counter := 0
	for k, v := range m{ //map 是无序的 map也是传的指针
		fmt.Println(k, " ", v)
		if counter == 0 {
			delete(m, "age") //因为map是无序的，所以删除了 可能输出也可能不输出
		}
		if counter == 1 {
			m["controy"] = "china" //同理，可能输出可能不输出
		}

		counter ++
	}
}

type mychan struct {
	Ch chan int
	Close bool
	timeout time.Duration
}

func rangeChannel()  {
	a := make(chan int, 10)
	ch := mychan{Ch:a, Close:false}

	//ch := make(chan int, 1)
	//is_close := false

	go func(ch1 *mychan) {
		defer func() {
			close(ch1.Ch)
			ch1.Close = true
		}()

		for i := 0; i <= 100; i++ {
			ch1.Ch <- i
		}
	}(&ch)

	for  {
		if ch.Close {
			fmt.Println("close")
			return
		}

		select {
		case c := <-ch.Ch:
			fmt.Println(c)
		case <-time.After(1 * time.Second):
			fmt.Println("timeout")
			return
		}
	}
}


//信道的生产者消费者
func Producer() *mychan {
	ch := make(chan int, 1)
	my_ch := mychan{
		Ch:ch,
		Close:false,
		timeout: 600 * time.Millisecond,
	}

	go func() {
		defer func() {
			close(my_ch.Ch)
			my_ch.Close = true
		}()

		for i := 0; i <= 20; i++ {
			time.Sleep(10 * time.Millisecond)
			my_ch.Ch <- i
		}
	}()

	return &my_ch
}

func Consumer(mychan *mychan)  {
	for  {
		if mychan.Close {
			goto end
		}

		for c := range mychan.Ch {
			fmt.Println("consumer: ", c)
		}
		//select {
		//case c :=<- mychan.Ch:
		//	fmt.Println("consumer: ", c)
		//case <-time.After(mychan.timeout):
		//	goto end
		//}
	}

	end:
		fmt.Println("End")
}

func TestPC()  {
	mych := Producer()
	Consumer(mych)
}