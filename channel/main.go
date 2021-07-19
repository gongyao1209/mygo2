package main

import (
	"fmt"
	"gocode/mygo2/channel/cancelsignal"
	"sync"
	"time"
)

//主要是阻塞
func main() {
	str := "2021-10-30"
	fmt.Println(preMonthLastDay(str))
	return
	//如果说 信道channel 是个盒子，往里面放东西的人 是 生产者，从里面拿东西的人作为消费者
	go TestChannel2() //
	//TestChannel1()
	//TestChannel2() //只有当这个的时候会报错，和main方法一个goroutine，所以是阻塞主进程了，所以才报错

	/**
	 *	这个为什么会出错？
	 *  从信道中获取数据会阻塞 main程，因为信道没有被生产者关闭，而且使用for range从信道中读数据没有超时限制，所里这里就一直阻塞
	 */
	//TestChannel2Error()
	//return

	// 改进方法 1
	// 信道的生产者 关闭信道，还是使用 for range 从信道中读取数据
	//TestChannel2()

	// 改进方法 2
	// 开辟新的 goroutine 使用 for range 从信道里面读取数据，阻塞也是阻塞读取信道的go程 对于main程并不会阻塞。但是这样是不安全的
	TestChannel2_1()
	//return
	//TestChannel2_1Error() //TODO 为什么这个也会阻塞，sync.WaitGroup
	return
	// 改进方法3
	// 使用 for 来读取信道里面的数据
	TestChannel2_2()

	//所以针对 main程的阻塞，1、信道是要在生产者中关闭 2、进行超时限制，防止阻塞
}

func TestChannel1() {
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

func preMonthLastDay(str string) string {
	ti, err := time.Parse("2006-01-02", str)
	if err != nil {
		return ""
	}
	ti2 := ti.AddDate(0, 0, -1 * ti.Day())
	return ti2.Format("2006-01-02")
}

func TestChannel2() {
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

func TestChannel2Error() {
	ch := make(chan int)
	defer close(ch)

	go func(ch chan int) {
		//defer close(ch)
		for i := 0; i < 10; i++ {
			ch <- i + 1
		}
	}(ch)

	for c := range ch { //这里可能引起阻塞
		fmt.Println("ch value : ", c)
	}
}

func TestChannel2_1() {
	ch := make(chan int)

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer func() {
			wg.Done()
			close(ch)
		}()
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for c := range ch {
			fmt.Println("ch value : ", c)
		}
	}()

	wg.Wait()
	//time.Sleep(1 * time.Second)
}

func TestChannel2_1Error() {
	ch := make(chan int)
	defer close(ch)

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer func() {
			wg.Done()
		}()
		for i := 0; i < 10; i++ {
			//ch <- i
		}
	}()

	wg.Add(1) //这里增加了 wg
	go func() {
		defer wg.Done()
		//select { // select 不会发生阻塞
		//case c1 :=<- ch:
		//	fmt.Println("ch value : ", c1)
		//default:
		//	time.Sleep(1 * time.Second)
		////case time.After(1 * time.Second):
		////	fmt.Println("ch value : ", c1)
		//}
		for c := range ch { //所以这里 go程阻塞 会引起 main程阻塞
			fmt.Println("ch value : ", c)
		}
	}()

	wg.Wait()
	time.Sleep(1 * time.Second)
}

func TestChannel2_2() {
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
		case c, ok := <-ch:
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

type mychan struct {
	Ch      chan int
	Close   bool
	timeout time.Duration
}

//信道的生产者消费者
func Producer() *mychan {
	ch := make(chan int, 1)
	my_ch := mychan{
		Ch:      ch,
		Close:   false,
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

func Consumer(mychan *mychan) {
	for {
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

func main1() {
	cancelsignal.NormalCancel3()
	//publish.Test()
	//locktest.Test1()
}
