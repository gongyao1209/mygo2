package main

import (
	"fmt"
	"gocode/mygo2/db"
	"time"
)

func main()  {
	s := db.GetData(1)
	time.Sleep(1 * time.Second)
	s2 := db.GetData(3)
	fmt.Println(s, s2)
}