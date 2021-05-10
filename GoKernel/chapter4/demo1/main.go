package main

import "fmt"

type myi interface {
	Hello(a string) string
}

type stu struct {}

func (s stu)Hello(a string) string {
	fmt.Println(a)
	return a
}

func main()  {

	var m myi
	m = stu{}
	fmt.Println(m.Hello("aaa"))
}