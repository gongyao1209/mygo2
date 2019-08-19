package main

import (
	"fmt"
	"time"
)

func main()  {

	var t task = ti

	p := NewPool(3)

	go func() {
		i := 0
		End:
		for true {
			p.EntryChannel <- &t
			time.Sleep(1 * time.Second)
			i++

			if i == 10 {
				break End
			}
		}
	}()

	p.Run()
}

func ti() string {
	return time.Now().String()
}

type task interface {}

type Pool struct {
	// 空的 信道
	EntryChannel chan *task
	// 工作的go程数量
	workNum int
	// 工作 信道
	JobsChannel chan *task
}

func NewPool(i int) *Pool {
	p := Pool{
		EntryChannel: make(chan *task),
		workNum: i,
		JobsChannel: make(chan *task),
	}

	return &p
}

func (p *Pool)worker(id int)  {

	//End:
	//for true {
	//	select {
	//	case t :=<- p.JobsChannel:
	//		fmt.Println("worker ID ", id, " 执行完毕任务",  *t)
	//	case <-time.After(10):
	//		break End
	//	//default:
	//	//	fmt.Println("----------------------------------")
	//	}
	//}

	for t := range p.JobsChannel {
		fmt.Println("worker ID ", id, " 执行完毕任务",  *t)
	}
}

func (p *Pool)Run()  {
	for i := 0; i < p.workNum; i++ {
		go p.worker(i)
	}

	for t := range p.EntryChannel {
		p.JobsChannel <- t
	}

	close(p.EntryChannel)
	close(p.JobsChannel)
}
