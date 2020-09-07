package test1

import (
	"fmt"
	"math/rand"
)

//type boring struct {
//	name string
//	quit chan bool
//}

func TestQuit()  {

	quit := make(chan bool)

	c := boring("Joe", quit)

	for i := rand.Intn(10); i >= 0; i-- {
		fmt.Println(<-c)
	}

	quit <- true
}