package main

import "fmt"

type People struct {
	Name string
	Age int
}

func (p *People)String() string { // 类型 T 只有接受者是 T 的方法；而类型 *T 拥有接受者是 T 和 *T 的方法。语法上 T 能直接调 *T 的方法仅仅是 Go 的语法糖。
	return fmt.Sprintf("name is %s, age is %d", p.Name, p.Age)
}

func main()  {
	gyao := People{
		Name:"gongyao",
		Age:27,
	}

	fmt.Println(gyao)
}