package package2

import (
	"context"
	"fmt"
	"math/rand"
	"runtime"
	"strconv"
	"sync"
	"time"
)

// 一直在运行的函数
func boring(msg string) {
	for i := 0; ; i++ {
		fmt.Println(msg, i)
		//time.Sleep(time.Second) //睡😴一秒
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}

func Test() { //主方法一个进程
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

func foo() {
	for i := 0; i <= 5; i++ {
		fmt.Println("I am runing, ", i)
	}
	syn <- 1 //
}
func Test2() {
	go foo()
	i := <-syn
	fmt.Println(i)
}

//信道 一个读一个写才能畅通无阻。可以使用信道进行交流和同步
func boring1(msg string, ch chan string) {
	for i := 0; ; i++ {
		ch <- fmt.Sprintf("Runing: %s; I: %d", msg, i)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}
func Test3() {
	c := make(chan string, 1)
	go boring1("boring", c)
	go boring1("wocao", c)

	for i := 0; i < 5; i++ {
		fmt.Println("Main func: ", <-c)
	}

	fmt.Println("I am leaving")
}

// 生产者模式
//生成器(Generator)
func boring2(msg string) <-chan string {
	ch := make(chan string)
	go func() {
		for i := 0; ; i++ {
			ch <- fmt.Sprintf("I am boring2, %s, %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return ch
}

func Test4() { // 和 Test8 类似，Test4执行完之后，boring2 中的 goroutine并没有退出，只是因为 信道给阻塞了
	ch := boring2("boring")

	for i := 0; i < 5; i++ {
		fmt.Println("test4: ", <-ch)
	}

	fmt.Println("You are boring, I am leaving")
}

func Test5() {
	gongyao := boring2("gongyao")
	yaoke := boring2("yaoke")

	for i := 0; i < 5; i++ {
		fmt.Println(<-gongyao)
		fmt.Println(<-yaoke)
	}

	fmt.Println("You are boring, I am leaving")
}

func fanIn(ch1, ch2 <-chan string) chan string {
	c := make(chan string)

	go func() {
		for { //一直要循环
			c <- <-ch1
		}
	}()

	go func() {
		for {
			c <- <-ch2
		}

	}()

	return c
}

func Test6() {
	c := fanIn(boring2("gongyao"), boring2("yaoke"))

	for i := 0; i < 5; i++ {
		fmt.Println(<-c)
	}
	fmt.Println("You are boring, I am leaving")
}

func Test7() {
	c := boring2("Joe")

	for v := range c { //遍历 即 从 c中吐数据
		fmt.Println(v)
	}
	//for {
	//	select {
	//	case s := <-c:
	//		fmt.Println(s)
	//	case <-time.After(1 * time.Second):
	//		fmt.Println("You're too slow.")
	//		return
	//	}
	//}
}

func boring3(msg string, quit chan bool) chan string {
	ch := make(chan string)

	go func() { //重新启动一个 goroutine 来做一些事情，就是为了简单的并发
		for i := 0; ; i++ {
			select {
			case ch <- fmt.Sprintf("Boring3, msg: %s, i: %d", msg, i):
				time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
			case <-quit:
				return //退出 goroutine
			}
		}
	}()

	return ch //返回一个信道的，函数里面 go func 的，基本都是生产者模式
}
func Test8() {
	quit := make(chan bool)
	c := boring3("gongyao", quit)

	go func(ch chan string, quit chan bool) {
		for i := 0; i < 100; i++ {
			select {
			case a := <-ch:
				fmt.Println(a)
			case <-quit:
				return
			}
		}
	}(c, quit)

	//go func() { //闭包之内不要调用别的变量
	//	for i:= 0; i < 100; i++ {
	//		select {
	//		case a :=<-c:
	//			fmt.Println(a)
	//		case <-quit:
	//			return
	//		}
	//	}
	//}()
	time.Sleep(5 * time.Second)
	quit <- true //退出 goroutine
}

func Test10() {

	ctx, cancel := context.WithCancel(context.Background())

	c := func(ctx2 context.Context) chan int {
		ch := make(chan int)

		go func(ch chan int) {
			for i := 0; ; i++ {
				select {
				case ch <- i:
					fmt.Println("go: ", i)
				case <-ctx2.Done():
					return
				}
			}
		}(ch)

		return ch
	}(ctx)

	for v := range c {
		fmt.Println(v)
		if v == 5 { //为什么 不会把
			cancel()
			break
		}
	}
}

func f(left, right chan int) {
	left <- <-right
}

func Test9() { //菊花链 嗷嗷待哺的信道
	const n = 100
	leftmost := make(chan int, 1)
	right := leftmost
	left := leftmost
	for i := 0; i < n; i++ {
		right = make(chan int, 1)
		go f(left, right)
		left = right
	}
	go func(c chan int) {
		c <- 2
	}(right)
	fmt.Println(<-leftmost)
}

//------------------------------------------GO 并发模式------------------------------------------------

//生产者
func Producer(factor int, ch chan int) {
	for i := 1; ; i++ {
		ch <- factor * i
	}
}

//消费者
func Consumer(ch chan int) {
	for value := range ch {
		fmt.Println(value)
	}
}

//不稳定的退出
func Test1010_01() {
	ch := make(chan int, 64)

	go func(ch chan int) {
		Producer(2, ch) //生产者之间互不影响
	}(ch)

	go func(ch chan int) {
		Producer(3, ch)
	}(ch)

	go Consumer(ch) //生产者 和 消费者 也互不影响

	time.Sleep(1 * time.Second)
}

//使用 quit 退出
func Test1010_02() {
	ch := make(chan int, 64)
	quit := make(chan bool)

	go func(ch chan int, quit chan bool) {
		for i := 0; ; i++ {
			select {
			case <-quit:
				return
			case ch <- 10 * i:
			}
		}
	}(ch, quit)

	go func(ch chan int, quit chan bool) {
		for i := 0; ; i++ {
			select {
			case <-quit:
				return
			case ch <- 3 * i:
			}
		}
	}(ch, quit)

	go func(ch chan int, quit chan bool) {
		for value := range ch {
			if value > 377262 {
				quit <- true
				return
			}

			fmt.Println(value)
		}
	}(ch, quit)

	time.Sleep(2 * time.Second)
	//quit<-true
}

//安全的退出 goroutine
func Test1010_03() {
	ch := make(chan int, 64)
	ctx, cancel := context.WithCancel(context.Background())

	go func(ch chan int, ctx2 context.Context) {
		for i := 0; ; i++ {
			select {
			case <-ctx2.Done():
				return
			case ch <- 10 * i:
			}
		}
	}(ch, ctx)

	go func(ch chan int, ctx2 context.Context) {
		for i := 0; ; i++ {
			select {
			case <-ctx2.Done():
				return
			case ch <- 3 * i:
			}
		}
	}(ch, ctx)

	go func(ch chan int) {
		for value := range ch {
			if value > 377262 {
				cancel()
				return
			}

			fmt.Println(value)
		}
	}(ch)

	time.Sleep(2 * time.Second)
}

type Students struct {
	Age int
	m   sync.RWMutex
}

func (s *Students) addAge(num int, p string) {
	s.m.Lock()
	//defer s.m.Unlock()
	s.m.RLock() //锁🔐
	defer s.m.RUnlock()

	for i := 0; i < 100; i++ {
		fmt.Println(i, p)

		s.Age += num
	}
}

func Test1012_01() {
	stu := &Students{Age: 0}

	go stu.addAge(10, "--------")
	go stu.addAge(10000, "********")
	go stu.addAge(100000000, "&&&&&&&&")

	time.Sleep(1 * time.Second)
	fmt.Println(stu)
}

//----其他并发
// 控制并发数据
func Test1012_02() {
	ch := make(chan bool, 3)
	var wg sync.WaitGroup
	fmt.Println("go number1 : ", runtime.NumGoroutine())
	fmt.Println(runtime.NumGoroutine())

	for i := 0; i < 10; i++ {
		wg.Add(1)
		ch <- true

		go func(ch chan bool, int2 int) {
			defer func() {
				wg.Done()
			}()
			fmt.Println(i)
			<-ch
		}(ch, i)
	}

	fmt.Println("go number2 : ", runtime.NumGoroutine())
	wg.Wait()
}

// 返回生成自然数序列的管道: 2, 3, 4, ...
func GenerateNatural() chan int {
	ch := make(chan int, 1)
	go func() {
		for i := 2; ; i++ {
			ch <- i
		}
	}()
	return ch
}

// 管道过滤器: 删除能被素数整除的数
func PrimeFilter(in <-chan int, prime int) chan int {
	out := make(chan int, 1)
	go func() {
		for {
			if i := <-in; i%prime != 0 {
				fmt.Println("-----", i)
				out <- i
			}
		}
	}()
	return out
}

func Test1012_03() {
	ch := GenerateNatural() // 自然数序列: 2, 3, 4, ...
	for i := 0; i < 10; i++ {

		time.Sleep(1 * time.Second)
		fmt.Printf("ch_pre %p\n", &ch)

		prime := <-ch // 新出现的素数
		fmt.Printf("%v: %v\n", i+1, prime)
		ch1 := PrimeFilter(ch, prime) // 基于新素数构造的过滤器
		fmt.Printf("ch1 %p\n", &ch1)
		ch = ch1
		fmt.Printf("ch_cur %p\n", &ch)
	}
}

func ToInt(str interface{}) (int, bool) {

	var r int

	switch str.(type) {
	case int:
		var ok bool
		if r, ok = str.(int); !ok {
			return 0, false
		}
		break
	case string:
		var err error
		r, err = strconv.Atoi(str.(string))
		if err != nil {
			fmt.Println(err)
			return 0, false
		}
		break
	default:
		return 0, false
	}

	return r, true
}

func ToStr(n interface{}) (string, bool) {
	switch n.(type) {
	case string:
		if r, ok := n.(string); ok {
			return r, true
		}
		return "", false
	case int:
		if t, ok := n.(int); ok {
			return strconv.Itoa(t), true
		}
		return "", false
	default:
		return "", false
	}
}
