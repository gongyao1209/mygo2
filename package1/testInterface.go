package package1

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type (
	push_driver chan interface{}
)
//
type Pusher struct {
	m sync.RWMutex
	buffer int
	timeout time.Duration
	drivers []push_driver
	quit chan bool
}

func p(v interface{})  {
	switch v.(type) {
	case int:
		fmt.Println("int ", v.(int))
	case string:
		fmt.Println("string ", v.(string))
	case Pusher:
		fmt.Println("Pusher ", v.(Pusher))
	case time.Duration:
		fmt.Println("time ", v)
	default:
		fmt.Println("other ", v)
	}
}

func Test101202()  {

	r := rand.New(rand.NewSource(time.Now().Unix()))

	sl := [10]int{}
	for i := 0; i < 10; i++ {
		sl[i] = r.Intn(1000)
		//sl = append(sl, r.Intn(1000))
	}

	fmt.Println(len(sl))
}