package package1

import (
	"fmt"
	"math/rand"
	"strings"
	"sync"
	"testing"
	"time"
)

type User struct {
	Name       string
	Mobile     string
	Age        int64
	Birthplace string
	Email      string
}

type User2 map[string]User

type Class struct {
	Student *User2
	Sum     int64
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSeq(n int64) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

/**
 * new 和 make 的区别，new 返回的是指针，make返回的是引用。指针 和 引用的区别： 指针不能修改原来分配内存的大小，但是引用能修改
 */
func newVsMake() {
	c1 := new([1]User)
	c2 := make([]User, 1)

	fmt.Println("new1 :", c1)  //new1 : &[{  0  }]
	fmt.Println("make1 :", c2) //make1 : [{  0  }]

	for i, _ := range c1 {
		c1[i] = User{
			Name:       "a",
			Mobile:     randSeq(1),
			Age:        rand.Int63(),
			Birthplace: randSeq(2),
			Email:      randSeq(3),
		}
	}

	for i, _ := range c2 {
		c2[i] = User{
			Name:       "b",
			Mobile:     randSeq(1),
			Age:        rand.Int63(),
			Birthplace: randSeq(2),
			Email:      randSeq(3),
		}
	}

	fmt.Println("new2 :", c1)  //new2 : &[{a X 8674665223082153551 lB zgb}]
	fmt.Println("make2 :", c2) //make2 : [{b a 894385949183117216 CM RAj}]

}

/**
 * 声明 和 初始化，映射声明的时候 并没有初始化
 */
func varVsNew() {
	u := User{
		Name:       "gongyao",
		Mobile:     randSeq(11),
		Age:        rand.Int63(),
		Birthplace: randSeq(2),
		Email:      randSeq(3),
	}

	u2 := new([1]User)    //创建的是数组
	var u1 []User         //创建的是切片
	u3 := make([]User, 0) //创建的是切片

	u1 = append(u1, u)
	u3 = append(u3, u)
	u2[0] = u

	fmt.Println("*u2 : ", *u2) //*u2 :  [{gongyao XVlBzgbaiCM 7504504064263669287 Aj Wwh}]
	fmt.Println("u1 : ", u1)   //u1 :  [{gongyao XVlBzgbaiCM 7504504064263669287 Aj Wwh}]
	fmt.Println("u3 : ", u3)   //u3 :  [{gongyao XVlBzgbaiCM 7504504064263669287 Aj Wwh}]

	//------上面是 切片

	var u4 User2
	u5 := new(User2) //new 返回指针, 指针只能修改原来分派的内存里面的数据，而不能修改存储的大小
	u7 := &User2{}
	u6 := User2{}

	//u4[u.Name] = u map不会 定义的时候给默认值，map得初始化。所以说 map并不是array、slice
	//u5[0][u.Name] = u
	u6[u.Name] = u

	fmt.Printf("u4 的地址: %p\n", u4) //map[] 这个时候还没有分配内存
	fmt.Println("u5 : ", u5)       //u5 :  &map[]
	fmt.Println("u7 : ", u7)       //u7 :  &map[]
	fmt.Println(u6)                //map[gongyao:{gongyao XVlBzgbaiCM 7504504064263669287 Aj Wwh}]

	u5 = &u6
	u6["b"] = u

	//u5[0] = User2{}
	//u5[0][u.Name] = u
	fmt.Println(u5) //&map[gongyao:{gongyao XVlBzgbaiCM 7504504064263669287 Aj Wwh} b:{gongyao XVlBzgbaiCM 7504504064263669287 Aj Wwh}]
	fmt.Printf("u6 之前的地址：%p\n", &u6)

	u6 = u4
	fmt.Printf("u4 的地址：%p\n", &u4)
	fmt.Printf("u6 的地址：%p\n", &u6)
	u6 = User2{}
	u6["g"] = u
	fmt.Printf("u6 的地址：%p\n", &u6)
	//
	//fmt.Println(u4)
	//fmt.Println(u6)
}

func str_map(str string) map[string]int {
	str_map := make(map[string]int)
	str_arr := strings.Split(str, "")

	for _, value := range str_arr {
		if str_map[value] == 0 {
			str_map[value] = 1
		} else {
			str_map[value]++
		}
	}

	return str_map
}

func compare(str1 string, str2 string) bool {

	if str1 == "" || str2 == "" {
		return false
	}

	str1_map := str_map(str1)
	fmt.Println(str1_map)
	str2_map := str_map(str2)
	fmt.Println(str2_map)

	if len(str1_map) != len(str2_map) {
		return false
	}

	for key, value := range str1_map {
		if str2_map[key] != value {
			return false
		}
	}

	return true
}

//测试信道
func TestChan1() {
	str1 := "gongyao"
	map1 := str_map(str1)

	max := 10

	type jishu struct {
		Str   *map[string]int
		Count string
	}

	ch2 := make(chan jishu, max)

	for i := 0; i < max; i++ {
		go func() {
			str2 := randSeq(8)
			m := str_map(str2)

			ch2 <- jishu{Str: &m, Count: str2} //
		}()
		<-time.After(time.Second) //go 线程的调用机制由 runtime决定
	}

	var wg sync.WaitGroup

	for i := 0; i < max; i++ { //这个是发生在 所有的上面之后
		a := <-ch2

		wg.Add(1)
		go func() {
			defer wg.Done()

			time.Sleep(1 * time.Second)

			if len(map1) != len(*a.Str) {
				//return
			}
			for key, value := range map1 {
				if (*a.Str)[key] != value {
					//return
				}
			}

			fmt.Println(a.Str, a.Count)
		}()
	}

	wg.Wait()
}

// 正确的做法
func Test_Select_Chan1(t *testing.T) { //中间执行了什么操作？
	readerChannel := make(chan int)
	go func(readerChannel chan int) {
		for {
			select {
			// 判断管道是否关闭
			case _, ok := <-readerChannel:
				if !ok {
					fmt.Println("close")
					goto BB
					//return
				} else {
					fmt.Println("Normal")
				}
			}
			t.Log("for")
		}
	BB:
		fmt.Println("BB")
	}(readerChannel)
	close(readerChannel)
	<-time.After(time.Second * 2)
}

func Test() {

	t := testing.T{}

	Test_Select_Chan1(&t)
	//TestChan1()
	return
}
