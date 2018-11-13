package funcTest

import "fmt"

type filter func(int, string) bool

func test(f filter, i int, s string) bool {
	return f(i, s)
}

func MyTest()  {
	fmt.Println(
		test(func(i int, s string) bool {
			
			return true
		},
		1, "aaa"))
}


//================================================================================
// 方法的调用
type Int int

func (i Int)Get()  {
	fmt.Println("Get")
}

func (i *Int)Set()  {
	fmt.Println("Set")
}

func Test()  {
	var i Int = 10

	(i).Get()
	(&i).Set()

	f := i.Get
	f()

	f1 := Int.Get
	f1(i)
}

//================================================================================

type T struct {
	a int
}

func (t T)Get()  {
	fmt.Println("Get T, value : ", t.a)
}

func (t *T)Set(in int)  {
	t.a = in
	fmt.Println("Set T, value : ", t.a)
}

func Test2()  {
	t := T{a:10}

	//方法表达式调用
	f := T.Get
	f(t)

	//普通方法调用
	t.Get()
	// 方法表达式调用
	f1 := t.Get
	f1()
	// 方法表达式调用
	f4 := (*T).Set
	f4(&t, 2)
	// 方法表达式调用
	(*T).Set(&t, 3)

	// 普通调用
	t.Set(1)
}