package locktest

import (
	"fmt"
	"sync"
	"time"
)

var m sync.Mutex
var wg sync.WaitGroup

func Test1() {
	m.Lock()

	wg.Add(2)

	go func() {
		defer func() {
			m.Unlock()
			wg.Done()
		}()

		time.Sleep(1 * time.Second)
		fmt.Println("111")
	}()

	go func() {
		defer func() {
			m.Unlock()
			wg.Done()
		}()

		m.Lock()
		fmt.Println("222")
	}()

	wg.Wait()
}
