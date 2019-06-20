package test190125

import "fmt"

type People interface {
	Jump()
	Run()
}

type chinese struct {
	name string
	age  int
}

func (c chinese) Jump() {
	s := fmt.Sprintf("chinese, name: %s, age: %d, Jump", c.name, c.age)
	fmt.Println(s)
}
func (c chinese) Run() {
	s := fmt.Sprintf("chinese, name: %s, age: %d, Run", c.name, c.age)
	fmt.Println(s)
}

type usa struct {
	name string
	age  int
}

func (u usa) Jump() {
	s := fmt.Sprintf("usa, name: %s, age: %d, Jump", u.name, u.age)
	fmt.Println(s)
}
func (u usa) Run() {
	s := fmt.Sprintf("usa, name: %s, age: %d, Run", u.name, u.age)
	fmt.Println(s)
}

func isChinese(p People) {
	if c, ok := p.(chinese); ok {
		c.Jump()
	} else {
		fmt.Println("People is not Chinese")
	}
}

func Test() {
	var p People

	c := chinese{name: "gongyao", age: 28}
	p = &c
	isChinese(c)
	p.Jump()
	p.Run()

	p = &usa{name: "wanghui", age: 29}
	p.Jump()
	p.Run()

}
