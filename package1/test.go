package package1

import (
	"fmt"
	"math/rand"
)

type User struct {
	Name		string
	Mobile		string
	Age 		int64
	Birthplace	string
	Email		string
}

type User2 map[string]User

type Class struct {
	Student	*User2
	Sum		int64
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
func newVsMake()  {
	c1 := new([1]User)
	c2 := make([]User, 1)

	fmt.Println("new1 :", c1) //new1 : &[{  0  }]
	fmt.Println("make1 :", c2) //make1 : [{  0  }]

	for i, _ := range c1 {
		c1[i] = User{
			Name: "a",
			Mobile:randSeq(1),
			Age: rand.Int63(),
			Birthplace:randSeq(2),
			Email:randSeq(3),
		}
	}

	for i, _ := range c2 {
		c2[i] = User{
			Name: "b",
			Mobile:randSeq(1),
			Age: rand.Int63(),
			Birthplace:randSeq(2),
			Email:randSeq(3),
		}
	}

	fmt.Println("new2 :", c1) //new2 : &[{a X 8674665223082153551 lB zgb}]
	fmt.Println("make2 :", c2) //make2 : [{b a 894385949183117216 CM RAj}]

}

func varVsNew()  {
	u := User{
		Name:"gongyao",
		Mobile:randSeq(11),
		Age: rand.Int63(),
		Birthplace:randSeq(2),
		Email:randSeq(3),
	}


	u2 := new([1]User) //创建的是数组
	var u1 []User //创建的是切片
	u3 := make([]User, 0) //创建的是切片

	u1 = append(u1, u)
	u3 = append(u3, u)
	u2[0] = u

	fmt.Println("*u2 : ", *u2) //*u2 :  [{gongyao XVlBzgbaiCM 7504504064263669287 Aj Wwh}]
	fmt.Println("u1 : ", u1) //u1 :  [{gongyao XVlBzgbaiCM 7504504064263669287 Aj Wwh}]
	fmt.Println("u3 : ", u3) //u3 :  [{gongyao XVlBzgbaiCM 7504504064263669287 Aj Wwh}]

	//------上面是 切片

	var u4 User2
	u5 := new(User2) //new 返回指针, 指针只能修改原来分派的内存里面的数据，而不能修改存储的大小
	u7 := &User2{}
	u6 := User2{}

	//u4[u.Name] = u map不会 定义的时候给默认值，map得初始化。所以说 map并不是array、slice
	//u5[0][u.Name] = u
	u6[u.Name] = u

	fmt.Println(u4) //map[]
	fmt.Println("u5 : ", u5) //u5 :  &map[]
	fmt.Println("u7 : ", u7) //u7 :  &map[]
	fmt.Println(u6) //map[gongyao:{gongyao XVlBzgbaiCM 7504504064263669287 Aj Wwh}]

	u5 = &u6
	u6["b"] = u

	//u5[0] = User2{}
	//u5[0][u.Name] = u
	fmt.Println(u5) //&map[gongyao:{gongyao XVlBzgbaiCM 7504504064263669287 Aj Wwh} b:{gongyao XVlBzgbaiCM 7504504064263669287 Aj Wwh}]
}

func Test()  {

	varVsNew()
	return
	var i int64
	var u  User
	//var u1 map[string]User = map[string]User{}
	u1 := User2{}

	u1["a"] = u
	//u1 = append(u1, u)
	return

	//u1.Name = "wanghui"
	//fmt.Println(u1)
	return

	//temp_u2 := User2{}
	var temp_u2 User2
	fmt.Println(&temp_u2)

	temp_u2 = User2{}
	fmt.Println(&temp_u2) //只声明 没有初始化 不能被赋值
	return

	c := Class{}

	for true {
		if i == 10 {
			break
		}

		var t_name string
		if i == 2 {
			t_name = "gongyao"
		} else {
			t_name = randSeq(10)
		}

		temp_u := User{
			Name: t_name,
			Mobile:randSeq(1),
			Age: rand.Int63(),
			Birthplace:randSeq(2),
			Email:randSeq(3),
		}

		//fmt.Println(temp_u2)
		//return

		temp_u2[temp_u.Name] = temp_u

		c.Sum++

		i++
	}

	c.Student = &temp_u2

	fmt.Println(c)
	fmt.Println((*c.Student)["gongyao"])

	//if len(*c.Student) > 0 {
	//	for _, u := range c.Student {
	//		fmt.Println(u)
	//	}
	//}
}