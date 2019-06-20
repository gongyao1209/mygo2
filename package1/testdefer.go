package package1

import "fmt"

func Defer1019_01() int {
	i := 1
	defer func() {
		fmt.Printf("defer中 i的地址: %p\n", &i)
		i++
	}()
	fmt.Printf("i的地址: %p\n", &i)
	return i
}

func Defer1019_0101() *int {
	i := 1
	defer func() {
		fmt.Printf("defer中 i的地址: %p\n", &i)
		i++
	}()
	fmt.Printf("i的地址: %p\n", &i)
	return &i
}

func Defer1019_04() int {
	i := 1
	defer func(i int) {
		fmt.Printf("defer中 i的地址: %p\n", &i)
		i++
	}(i)
	fmt.Printf("i的地址: %p\n", &i)
	return i
}

// 校验defer方法
func Defer1019_10() (i int) {
	i = 1
	defer func() {
		fmt.Printf("defer中 i的地址: %p\n", &i) //defer中 i的地址: 0xc4200140a0
		i++
	}()
	fmt.Printf("i的地址: %p\n", &i) //i的地址: 0xc4200140a0
	return
}
func Defer1019_11() (i int) {
	i = 1
	defer func(i int) {
		fmt.Printf("defer中 i的地址: %p\n", &i) //defer中 i的地址: 0xc4200140c0
		i++
	}(i)
	fmt.Printf("i的地址: %p\n", &i) //i的地址: 0xc4200140a8
	return
}

func Defer1019_02() int {
	i := 1
	defer func(i int) {
		fmt.Printf("defer中 i的地址: %p\n", &i)
		i++
	}(i)

	fmt.Printf("i的地址: %p\n", &i)
	return i
}

func Defer1019_03() (i int) {
	i = 1
	defer func(i *int) {
		fmt.Printf("defer %p\n", i)
		*i++
	}(&i)

	fmt.Printf("i的地址: %p\n", &i)
	return
}

func Defer1019_05() {

}
