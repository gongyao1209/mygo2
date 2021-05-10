package main

import (
	"fmt"
	"strings"
)


func main()  {
	str1 := "abccc"
	str2 := "cccab"

	fmt.Println(isRegroup(str1, str2))
}

// 其实是 为了判断 两个字符串中每个字符出现的次数一样不
func isRegroup(string1 string, string2 string) bool {
	l1 := len([]rune(string1))
	l2 := len([]rune(string2))

	if l1 > 5000 || l2 > 5000 || l1 != l2 {
		return false
	}

	for _, v := range string1 {
		if strings.Count(string1, string(v)) != strings.Count(string2, string(v)) {
			return false
		}
	}

	return true
}