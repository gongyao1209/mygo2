package publish

// 利用这个模型 做出我们的推送系统，哇咔咔
import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

//信道作为客户
type client chan interface{}

type publisher struct {
	Client map[string]PushMsg
}

func (p *publisher) Publish(msg string) {
	if len(p.Client) == 0 {
		return
	}
	for _, ch := range p.Client {
		ch <- msg
	}
}

func (p *publisher) CloseAll() {
	if len(p.Client) == 0 {
		return
	}

	for k, ch := range p.Client {
		delete(p.Client, k)
		close(ch)
	}
}

func (p *publisher) CloseName(name string) {
	if cli, ok := p.Client[name]; ok {
		close(cli)
		delete(p.Client, name)
	}
}

func Test() {
	//fmt.Println("1, ", runtime.NumGoroutine())
	//ch1 := get_driver("gongyao")
	//fmt.Println("2, ", runtime.NumGoroutine())
	//ch2 := get_driver("yaoke")
	//fmt.Println("3, ", runtime.NumGoroutine())

	g1 := NewGroup()

	fmt.Println("1, ", runtime.NumGoroutine())
	//ch2 := make(chan interface{}, 10)

	p := publisher{Client: make(map[string]PushMsg)}

	p.Client["a"] = g1.Msg
	//p.Client["b"] = ch2

	var wg sync.WaitGroup
	for i := 1; i <= 2; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			p.Publish(string(i))
		}()
		fmt.Println("2, ", runtime.NumGoroutine())
	}

	wg.Wait()
	//go func() {
	//	p.Publish("gongyao1")
	//	p.Publish("gongyao2")
	//	p.Publish("gongyao3")
	//
	//	p.CloseAll()
	//p.CloseName("b")
	//}()
	time.Sleep(1 * time.Second)

	p.Publish("gongyao4")
	fmt.Println("4, ", runtime.NumGoroutine())

	p.CloseAll()

	time.Sleep(1 * time.Second)

	fmt.Println("5, ", runtime.NumGoroutine())
	return
	//go func() {
	//	for c := range ch1 {
	//		fmt.Println("Ch1 : ", c)
	//	}
	//}()
	//
	//go func() {
	//	for c := range ch2 {
	//		fmt.Println("Ch2 : ", c)
	//	}
	//}()

	//ok1 := true
	//ok2 := true
	//
	//for {
	//	if !ok1 && !ok2 {
	//		goto End
	//	}
	//
	//	var c interface{}
	//	select {
	//	case c, ok1 = <- ch1:
	//		if ok1 {
	//			fmt.Println("ch1 : ", c)
	//		}
	//	case c, ok2 = <- ch2:
	//		if ok2 {
	//			fmt.Println("ch2 : ", c)
	//		}
	//	case <-time.After(1 * time.Second):
	//		goto End
	//	}
	//}
	//
	//End:
	//	fmt.Println("End")
}

// 获取节点 这才是真正的订阅发布者模式 哈哈
func get_driver(name string) chan interface{} {
	ch := make(chan interface{}, 10)
	go func() {
		for {
			select {
			case c, ok := <-ch:
				if !ok {
					goto End
				}
				fmt.Println(name, ":", c)
			default:
				time.Sleep(10 * time.Microsecond)
			}
		}

	End:
		fmt.Println("End")
	}()
	return ch
}
