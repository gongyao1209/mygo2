package main

import "fmt"

func main()  {

	//var i Lang
	//p := PHP{}
	////g := Go{}
	//
	//i = p
	//Say(i)
	//
	////Say(&g)

	var p People
	p = Gong{13}

	age := p.sayAge()
	p.upAge()
	age = p.sayAge()

	fmt.Println(age)
}

type Lang interface {
	SayHello()
	SayHi()
}

func Say(l Lang)  {
	l.SayHello()
	l.SayHi()
}

type PHP struct {
}
func (p PHP)SayHello() {
	fmt.Println("Hello, I am PHP")
}
func (p PHP)SayHi() {
	fmt.Println("Hi, I am PHP")
}

type Go struct {
}

func (g Go)SayHello()  {
	fmt.Println("Hello, I am Go")
}


type People interface {
	sayAge() int
	upAge()
}

type Gong struct {
	age int
}

func (g Gong)sayAge() int {
	return g.age
}

func (g Gong)upAge() {
	g.age++
}