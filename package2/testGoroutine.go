package package2

import (
	"fmt"
	"math/rand"
	"time"
)

// 一直在运行的函数
func boring(msg string)  {
	for i := 0; ; i++ {
		fmt.Println(msg, i)
		//time.Sleep(time.Second) //睡😴一秒
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}

func Test()  { //主方法一个进程
	//var wg sync.WaitGroup
	//wg.Add(1)
	go func() { //goroutine 新开一个进程
		//defer wg.Done()
		boring("boring")
	}()
	//wg.Wait()

	fmt.Println("listening ")
	time.Sleep(2 * time.Second)
	fmt.Println("You're boring; I'm leaving.")
}

//--信道保证 同步
var syn chan int = make(chan int)
func foo()  {
	for i:= 0; i <= 5; i++ {
		fmt.Println("I am runing, ", i)
	}
	syn <- 1 //
}
func Test2()  {
	go foo()
	i :=<-syn
	fmt.Println(i)
}


//信道 一个读一个写才能畅通无阻。可以使用信道进行交流和同步
func boring1(msg string, ch chan string)  {
	for i := 0; ; i++ {
		ch <- fmt.Sprintf("Runing: %s; I: %d", msg, i)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}
func Test3()  {
	c := make(chan string, 1)
	go boring1("boring", c)
	go boring1("wocao", c)


	for i := 0; i < 5; i++ {
		fmt.Println("Main func: ", <- c)
	}


	fmt.Println("I am leaving")
}


// 生产者模式
//生成器(Generator)
func boring2(msg string) chan string {
	ch := make(chan string)
	go func() {
		for i := 0; ; i++ {
			ch <- fmt.Sprintf("I am boring2, %s, %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return ch
}

func Test4()  { // 和 Test8 类似，Test4执行完之后，boring2 中的 goroutine并没有退出，只是因为 信道给阻塞了
	ch := boring2("boring")

	for i := 0; i < 5; i++ {
		fmt.Println("test4: ", <-ch)
	}

	fmt.Println("You are boring, I am leaving")
}

func Test5()  {
	gongyao := boring2("gongyao")
	yaoke := boring2("yaoke")

	for i := 0; i < 5; i++ {
		fmt.Println(<-gongyao)
		fmt.Println(<-yaoke)
	}

	fmt.Println("You are boring, I am leaving")
}

func fanIn(ch1, ch2 chan string) chan string {
	c := make(chan string)

	go func() {
		for  { //一直要循环
			c <- <- ch1
		}
	}()

	go func() {
		for  {
			c <- <- ch2
		}

	}()

	return c
}

func Test6()  {
	c := fanIn(boring2("gongyao"), boring2("yaoke"))

	for i := 0; i < 5; i++ {
		fmt.Println(<-c)
	}
	fmt.Println("You are boring, I am leaving")
}


func Test7() {
	c := boring2("Joe")
	for {
		select {
		case s := <-c:
			fmt.Println(s)
		case <-time.After(1 * time.Second):
			fmt.Println("You're too slow.")
			return
		}
	}
}


func boring3(msg string, quit chan bool) chan string {
	ch := make(chan string)

	go func() { //重新启动一个 goroutine 来做一些事情，就是为了简单的并发
		for i := 0; ; i++ {
			select {
			case ch <- fmt.Sprintf("Boring3, msg: %s, i: %d", msg, i):
				time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
			case <- quit:
				return //退出 goroutine
			}
		}
	}()

	return ch //返回一个信道的，函数里面 go func 的，基本都是生产者模式
}
func Test8()  {
	quit := make(chan bool)
	c := boring3("gongyao", quit)

	for i:= 0; i < 5; i++ {
		fmt.Println(<-c)
	}
	quit<-true
}



func f(left, right chan int) {
	left <- 1 + <-right
}

func Test9() { //菊花链 什么原理？
	const n = 100000
	leftmost := make(chan int)
	right := leftmost
	left := leftmost
	for i := 0; i < n; i++ {
		right = make(chan int)
		go f(left, right)
		left = right
	}
	go func(c chan int) {
		c <- 2
		}(right)
	fmt.Println(<-leftmost)
}