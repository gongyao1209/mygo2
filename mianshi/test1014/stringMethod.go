package test1014

import (
	"fmt"
	"strings"
)

func StringMethod() {
	s1 := "gongyao,dadad,adadq,3123 131"
	s2 := "yaoke"

	//d := strings.Compare(s1, s2)
	d := strings.Contains(s1, "gong")
	fmt.Println("Contains ", d)

	d2 := strings.IndexByte(s1, s2[0])

	fmt.Println(d2, s2[0])

	d3 := strings.ContainsAny(s1, "g")
	fmt.Println(d3)

	d4 := strings.Count(s1, "ad")
	fmt.Println(d4)

	d5 := strings.Index(s1, "ad")
	fmt.Println(d5)

	d6 := strings.Fields(s1)
	for _, v := range d6 {
		fmt.Println(v)
	}

	//将字符串打散到数组里面 可以自己填写规则
	d7 := strings.FieldsFunc(s1, func(r rune) bool {
		ds := []rune(", ")
		for k, _ := range ds {
			if r == ds[k] {
				return true
			}
		}
		return false
	})
	fmt.Println("==============")
	for _, v := range d7 {
		fmt.Println(v)
	}

	d8 := make([]byte, 10)
	copy(d8, s1)
	fmt.Println(string(d8))

	d9 := strings.Repeat(s1, 2)
	fmt.Println(d9)

	//strings.Map(func(r rune) rune {
	//
	//}, s1)

	d10 := strings.Title(d9)
	fmt.Println(d10)

	d11 := []rune(d10)
	d12 := []byte(d10)
	fmt.Println(d11)
	fmt.Println(d12)
}
