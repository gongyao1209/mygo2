package main

import (
	"fmt"
	"runtime"
	"strconv"
	"sync"
	"time"
)

// 举个例子 推送模型
func main()  {
	p := newDriverPool()

	fmt.Println("协程数量1", runtime.NumGoroutine())
	for i := 0; i < 100; i++ {
		name := "gongyao" + strconv.Itoa(i)
		p.addDriver(name)

	}
	fmt.Println("协程数量2", runtime.NumGoroutine())
	p.sendMsg("hello")
	//p.sendMsg("hello2")

	time.Sleep(2 * time.Second)
	p.close()
	time.Sleep(1 * time.Second)

	fmt.Println("协程数量3", runtime.NumGoroutine())
}

// 一个人一个通道
type driver struct {
	Ch chan interface{}
	IsClose bool
}

func (d *driver)close()  {
	close(d.Ch)
	d.IsClose = true
}

// 一个人一个通道
type driverPool struct {
	M map[string]*driver
	X sync.Mutex
}

func newDriverPool() *driverPool {
	return &driverPool{
		M: make(map[string]*driver),
	}
}

func (p *driverPool)addDriver(name string) {
	p.X.Lock()
	defer p.X.Unlock()

	d := driver{
		Ch:      make(chan interface{}),
		IsClose: false,
	}
	go func(d *driver) {
		for true {
			if d.IsClose {
				return
			}
			select {
			case h := <- d.Ch:
				time.Sleep(1 * time.Second)
				fmt.Println(name, " ", h)
			default:
				//fmt.Println(name, " 休息")
				time.Sleep(100 * time.Microsecond)
			}
		}
	}(&d)

	p.M[name] = &d
}

func (p *driverPool)delDriver(name string) string {
	p.X.Lock()
	defer p.X.Unlock()

	if _, ok := p.M[name]; ok {
		p.M[name].close()
		delete(p.M, name)
		return "success"
	}

	return "empty"
}

func (p *driverPool)close() {
	if len(p.M) == 0 {
		return
	}
	for name, _ := range p.M {
		p.delDriver(name)
	}
}

func (p *driverPool)sendMsg(msg string)  {
	p.X.Lock()
	defer p.X.Unlock()

	if len(p.M) == 0 {
		return
	}
	for name, _ := range p.M {
		p.M[name].Ch <- msg
	}
}