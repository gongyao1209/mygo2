package test1014

import (
	"fmt"
	"runtime"
	"sync"
)

//----考察了 defer
// 讲解：https://tiancaiamao.gitbooks.io/go-internals/content/zh/03.4.html
func DeferCall() interface{} {

	defer func() {
		fmt.Println("打印前")
	}()

	defer func() {
		fmt.Println("打印中")
	}()

	defer func() {
		fmt.Println("打印后")
	}() //defer 放到栈里面，后面的先执行

	panic("出发异常")

	return "a"
}

func DeffCall_1() (r int) {
	defer func() {
		r++
	}()
	return 1
}

func DeffCall_2() (r int) {
	t := 5
	defer func() {
		t = t + 5
	}()
	return t
}

func DeffCall_3() (r int) {
	defer func(r *int) {
		*r += 1
	}(&r)
	return 1
}

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, " ", a, " ", b, " ", ret)
	return ret
}

func DefferCall_4() {
	a := 1
	b := 2

	//defer func(a, b int) {
	//	calc("1", a, calc("10", a, b))
	//}(a, b)

	//defer func() {
	//	calc("1", a, calc("10", a, b))
	//}() //

	defer func() {
		calc("1", a, calc("10", a, calc("100", a, b))) //2
	}()
	//defer calc("1", a, calc("10", a, calc("100", a, b))) //2

	a = 0
	defer calc("2", a, calc("20", a, b))
	b = 1
}

func add1(a, b *int) int {
	*a++
	*b++

	fmt.Printf("a = %d, b = %d\n", *a, *b)

	return *b
}

func DefferCall_5() {
	a := 1
	b := 2

	c := add1(&a, &b)

	defer add1(&a, &c)

}

//---------
type student struct {
	Name string
	Age  int
}

// 这道题主要考的是 for 循环里面 临时变量的创建
func PaseStudent() {
	m := make(map[string]*student)

	stus := []student{
		{Name: "gongyao1", Age: 1},
		{Name: "gongyao2", Age: 2},
		{Name: "gongyao3", Age: 3},
	}

	for k, stu := range stus { //stu 属于创建的临时变量，而且并不是每一次新建都会创建，而是只创建一次
		fmt.Printf("%d : %p\n", k, &stu)
		m[stu.Name] = &stu
	}
}

//---------

//主要考察 闭包函数 变量的作用域
func Test1014_01() {
	runtime.GOMAXPROCS(1)
	wg := sync.WaitGroup{}

	wg.Add(20)

	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println("go 1 :", i) //打印i
			wg.Done()
		}()
	}

	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println("go 2 :", i)
			wg.Done()
		}(i) //把i作为参数传递进行打印
	}

	wg.Wait()
}

//--------
type People struct {
	Name string
}

func (p *People) ShowA() {
	fmt.Println("show A")
	p.ShowB()
}
func (p *People) ShowB() {
	fmt.Println("People show B ", p.Name)
}

type Teacher struct {
	People
}

func (t *Teacher) ShowB() {
	fmt.Println("Teacher show B")
}

func Test1014_02() {
	t := Teacher{}

	t.ShowA() //主要考察Go中 虽然没有继承，但是这种语法怎么办？
	t.ShowB()
}

//------
func Test1014_03() {
	runtime.GOMAXPROCS(1)

	int_chan := make(chan int, 1)
	string_chan := make(chan string, 1)

	int_chan <- 1
	string_chan <- "hello"

	select { //主要考察了 select 语句 当满足要求的时候 执行的随机性
	case value := <-int_chan:
		fmt.Println(value)
	case value := <-string_chan:
		panic(value)
	}
}

//----
func Test1014_04() {
	s := make([]int, 5)

	s = append(s, []int{1, 2, 3}...)

	fmt.Println(s)
}

// ----

type UserAges struct {
	ages map[string]int
	sync.Mutex
}

func (u *UserAges) Add(name string, age int) {
	u.Lock()
	defer u.Unlock()

	u.ages[name] = age
}

func (u *UserAges) Get(name string) int {
	if age, ok := u.ages[name]; ok {
		return age
	}

	return -1
}

func Test1014_05() {
	u := UserAges{
		ages: make(map[string]int),
	}

	u.Add("gongyao", 26)
	fmt.Println(u)

	fmt.Println(u.Get("gongyao"))
}

//------

type People2 interface {
	Speak(string) string
}

type Student2 struct {
}

func (s *Student2) Speak(think string) (talk string) {
	if think == "bitch" {
		talk = "You are a good boy"
	} else {
		talk = "hi"
	}

	return
}

func Test1014_06() {
	//var p People2
	s := Student2{}
	fmt.Println(s.Speak("bitch"))

}
