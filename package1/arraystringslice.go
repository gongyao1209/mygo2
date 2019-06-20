package package1

import (
	"fmt"
	"unicode/utf8"
)

//1、----数组
func test1Array() {
	// 数组的定义
	var a [3]int
	var b = [...]int{1, 2, 3}
	var c = [...]int{2: 1, 1: 3}
	var d = [...]int{5: 3}

	fmt.Println(a, b, c, d)
}

//数组作为参数
func test2Array() {
	var a = [...]int{1, 2, 3}
	var b = &a

	fmt.Println("pre a: ", a)
	fun_a(a)
	fmt.Println("after a: ", a)

	//for i, v := range a {
	//	fmt.Println(i, v)
	//}
	fmt.Println("pre b: ", b)
	fun_b(b)
	fmt.Println("after b: ", b)
	fmt.Println("after2 a: ", a)
}

func fun_a(a [3]int) { //数组作为参数，传递的是 备份
	a[2] = 10
}
func fun_b(b *[3]int) {
	b[2] = 10
}

//多维数组
func test3Array() {
	var times [5][0]int //并没有分配内存

	for range times {
		fmt.Println("hello")
	}

	fmt.Printf("%p\n", times)
}

//2、----字符串
func test1String() {
	s1 := "hello, world"[:5]
	s2 := "hello, world"[7:]

	fmt.Println(s1)
	fmt.Println(s2)

	s3 := "世界abc"
	arr := []int32(s3)

	fmt.Println(arr)
	fmt.Println(runes2string(arr))

	//fmt.Println(arr)
}

func runes2string(s []int32) string {
	var p []byte
	buf := make([]byte, 3)
	for _, r := range s {
		n := utf8.EncodeRune(buf, r)
		p = append(p, buf[:n]...)
	}
	return string(p)
}

//3、----切片
//切片存储 cap 和 len
func test1Slice() {
	a := make([]int, 4, 4)
	b := make([]int, 4, 5)
	fun_c(a)
	fun_c(b)

	fmt.Printf("pre   a: %p \n", a)
	a = append(a, 10) //a 分配了 新存储空间
	fmt.Printf("after a: %p \n", a)

	fmt.Printf("pre   b: %p \n", b)
	b = append(b, 10) //b 并没有分配新的存储空间
	fmt.Printf("after b: %p \n", b)
}

func test2Slice() {
	a := make([]int, 4, 5)
	fun_c(a)

	fmt.Printf("addr1: %p\n", a)
	a = append(a, 0)
	fmt.Printf("addr2: %p\n", a)
	temp := 1
	copy(a[temp:], a)
	fmt.Printf("addr3: %p\n", a)

	b := make([]int, 4, 4)
	fun_c(b)
	c := append(b[:temp], append([]int{10}, b[temp:]...)...)
	fmt.Println(c)
	c2 := append(append(b[0:temp], []int{10}...), b[temp:]...)
	fmt.Println(c2)

	d := make([]int, 4, 5)
	//fun_c(d)
	d1 := append([]int{10, 11}, d[temp:]...)
	fmt.Println(d)
	fmt.Println(d1)

	fmt.Printf("d的地址1:%p, d的值%v\n", d, d)
	d2 := append(d[:temp], []int{10, 11, 12, 13}...)
	//d3:= append(d2, d[temp:]...)
	fmt.Printf("d的地址2:%p, d的值%v\n", d, d)
	fmt.Println(d2)

}

func fun_c(a []int) { //切片作为参数的时候 传递的是 sliceheader
	a[0] = 1
	a[1] = 2
	a[2] = 3
	a[3] = 4
}

func Test1() {
	//test2Slice()

	var a = []int{1, 2, 3}

	if true {
		a[1] = 10
	}

	fmt.Println(a)
}
