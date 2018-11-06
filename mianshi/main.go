package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main()  {
	//c := make(chan int)
	//
	//go func() {
	//	time.Sleep(1 * time.Second)
	//	//c <- 1
	//	fmt.Println(<- c)
	//}()
	//
	//fmt.Println(1)
	//c <- 1
	//fmt.Println(2)
	rangeMap()
	//rangeSlice()
	//TestPC()
	return
	rangeChannel()
	return
	ch := make(chan int, 1)
	defer close(ch)

	var wg sync.WaitGroup

	fmt.Println("goroutine num1: ", runtime.NumGoroutine())

	var m sync.RWMutex

	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer func() {
				wg.Done()
				m.Unlock()
			}()

			m.Lock()
			//ch <- i
			//fmt.Println("add to channel, ", i)

			select {
			case ch <- i:
				fmt.Println("add to channel, ", i)
			default:
				fmt.Println("close channel, ", i)
			}
		}(i)
	}
	fmt.Println("goroutine num2: ", runtime.NumGoroutine())

	for {
		select {
		case c := <-ch:
			fmt.Println(c)
		case <-time.After(1 * time.Second): //超时
			goto Gongyao
		}
	}

	Gongyao:
		fmt.Println("end")
	//for c := range ch {
	//	fmt.Println(c)
	//}
	//fmt.Println(<- ch) //当一个goroutine正在运行的时候，chan才有用
	wg.Wait()

	fmt.Println("goroutine num3: ", runtime.NumGoroutine())

	return
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