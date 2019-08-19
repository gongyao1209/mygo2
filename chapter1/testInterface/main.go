package main

import (
	"fmt"
	"unsafe"
)

func main()  {

	var i Lang
	p := PHP{}
	////g := Go{}
	//
	i = p
	Say(i)

	return
	//
	////Say(&g)

	//testInterType()
	//testInterType2()
	//return

	//var p People
	//
	//fmt.Printf("p2: %T, %v\n", p, p) // 动态类型，动态值。
	//
	//p = &Gong{13}
	//fmt.Printf("p2: %T, %v\n", p, p)
	//
	//age := p.sayAge()
	//p.upAge()
	//age = p.sayAge()
	//
	//fmt.Println(age)
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

func testInterType()  {
	var p People
	var g *Gong
	
	fmt.Println(p == nil) // p 的动态类型 和 动态值 都为 nil
	fmt.Printf("p: %T, %v\n", p, p)
	
	fmt.Println(g == nil) // g 的动态值为 nil
	fmt.Printf("g: %T, %v\n", g, g)
	
	p = g // p 的动态类型变为 *main.Gong 动态值为nil
	fmt.Println(p == nil) // 此时 p 和 nil 的比较为false
	fmt.Printf("p2: %T, %v\n", p, p)

	var _ People = (*Gong)(nil)
}

type ifacee struct {
	itab, data uintptr
}

func testInterType2()  {
	var a interface{} = (*int)(nil)

	ia := *(*ifacee)(unsafe.Pointer(&a))

	fmt.Println(ia)

}