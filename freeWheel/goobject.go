package freeWheel

import (
	"errors"
	"fmt"
	"reflect"
)

//import "fmt"

//结构体
type Person struct {
	Name string
	Age  int
}

type Employee struct {
	Person //通过组合类似继承
	Job    string
}

// Go通过组合实现类似继承，但是还有区别。Employee的实例并不是Person的实例
func playWith(p *Person) {

}

func TestStruct() {
	p := &Person{}
	e := &Employee{Job: "a"}

	playWith(p)
	//playWith(e) //cannot use e (type *Employee) as type *Person in argument to playWith
	fmt.Println(e)
}

//函数

// 编译器将 相同签名(参数和返回值列表相同)的函数视为同一类型
type operater func(int, int) int

func getOperator(op string) operater {
	switch op {
	case "+":
		return func(a, b int) int {
			return a + b
		}
	default:
		return nil
	}
}

func TestFun1() {
	f := getOperator("+")(1, 3)
	fmt.Println(f)
}

//带个初始值
func getOperatorWithValue(op string, a int) operater {
	switch op {
	case "+":
		return func(i1 int, i2 int) int {
			return a + i1 + i2
		}
	default:
		return nil
	}
}
func TestFun2() {
	f := getOperatorWithValue("+", 1)(1, 3)
	fmt.Println(f)
}

//可变长度的参数
type operaters func(...int) int

func getOperatorWithValue1(op string, v int) (fn operaters, err error) {
	switch op {
	case "+":
		fn := func(i ...int) int {
			sum := v
			for _, j := range i {
				sum += j
			}
			return sum
		}

		return fn, nil
	default:
		return nil, errors.New("不支持的操作符")
	}
}
func TestFun3() {
	if fu, err := getOperatorWithValue1("+", 0); err == nil {
		fmt.Println(fu(1, 2, 3))
		fmt.Println(fu(2, 3, 4, 5))
	}
}

//闭包
func getOperatorWithValue2(op string, v int) (fn operaters, err error) {
	switch op {
	case "+":
		fn := func(args ...int) int {
			for _, value := range args {
				v += value
			}

			return v
		}

		return fn, nil
	default:
		return nil, errors.New("不支持的操作符")
	}
}
func TestFun4() {
	if fu, err := getOperatorWithValue2("+", 0); err == nil {
		fmt.Println(fu(1, 2, 3))
		fmt.Println(fu(4, 5, 6))
	}
}

// Go通过设计模式解决重载
// 利用函数匹配参数列表
func Brew(i1 int, str string, i2 int) {
	fmt.Println(i1, str, i2)
}

func coffee1() (int, string, int) {
	return 3, "coffee1", 2
}
func coffee2(a int) (int, string, int) {
	return a, "coffee2", 2
}
func coffee3(name string) (int, string, int) {
	return 1, name, 2
}

func TestFun5() {
	Brew(coffee1())
	Brew(coffee2(10))
	Brew(coffee3("coffee3"))
}

// 方法 Method
// 方法可以看做是特殊的函数。与普通函数不同，方法需要与对象实例绑定，在定义语法上方法有前置的接受者
// 可以为除了接口和指针外的任何类型定义方法
type N int

func (n *N) double() {
	*n = *n * 2
}

func TestMethod1() {
	var i N = 10
	fmt.Println(i)

	i.double()
	fmt.Println(i)
}

func (p *Person) Talk() {
	fmt.Println("Person Talk : ", p)
}
func (e *Employee) Talk() {
	fmt.Println("Employee Talk : ", e)
}
func TestMethod2() {
	e := &Employee{}
	e.Talk()
}

//方法定义在函数上
type myfun func(int, int) int

func (mf myfun) Add(a, b int) int {
	return mf(a, b)
}

func TestMethod3() {
	fn1 := func(a, b int) int {
		return a * b
	}
	fn2 := func(a, b int) int {
		return a + b
	}

	fmt.Println(myfun(fn1).Add(2, 3))
	fmt.Println(myfun(fn2).Add(2, 3)) //需要强制类型转换
}

// 接口 interface{}
type Talker interface {
	Talk()
}

//通过接口实现多态
func justTalk(t Talker) {
	t.Talk()
}

func TestInterface1() {
	t := &Person{}
	justTalk(t)

	t2 := &Employee{}
	justTalk(t2)
}

// 反射
type X int

func TestReflect1() {
	var a X = 1
	t := reflect.TypeOf(a)

	fmt.Println(t.Name(), t.Kind())
}

type Foo struct {
	FirstName string `tag_name:"tag 1"`
	LastName  string `tag_name:"tag 2"`
	Age       int    `tag_name:"tag 3"`
}

func (f *Foo) reflect() {
	val := reflect.ValueOf(f).Elem()

	val.NumField()
	fmt.Println(val.NumField())

	for i := 0; i < val.NumField(); i++ {
		val_field := val.Field(i)
		type_field := val.Type().Field(i)

		tag := type_field.Tag

		fmt.Printf("Field Name: %s, \t Field Value : %v, \t Tag Value: %s\n", type_field.Name, val_field.Interface(), tag.Get("tag_name"))
	}
}

func TestReflect() {
	f := &Foo{
		FirstName: "gongyao",
		LastName:  "Yao",
		Age:       26,
	}
	f.reflect()
}
