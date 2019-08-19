package main

import "fmt"

type people interface {
	Say()
}

func Hi(p people)  {
	p.Say()
}

type ShanxiEr struct {
}

func (e ShanxiEr)Say()  {
	fmt.Println("你好，俺是山西人")
}

func main()  {
	var p people = ShanxiEr{}
	Hi(p)
}