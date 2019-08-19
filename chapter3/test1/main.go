package main

import (
	"fmt"
	"sync"
)

type People struct {
	Name string
}

func (p People)baogao(g *sync.WaitGroup)  {
	defer g.Done()

	var i int = 1
	Han:
	for true {
		fmt.Println("my name is ", p.Name, ", baogao: ", i)
		i++
		if i == 100 {
			break Han
		}
	}
}

func main()  {
	p1 := People{Name:"gongyao"}
	p2 := People{Name:"wanghui"}

	w := &sync.WaitGroup{}

	w.Add(2)
	go p1.baogao(w)
	go p2.baogao(w)

	w.Wait()

	//time.Sleep(1 * time.Second)
}