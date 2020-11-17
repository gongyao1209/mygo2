package main

import "fmt"

func main()  {
	s := "巩尧"

	ss, _ := reverString(s)

	fmt.Print(ss)
}

func reverString(s string) (string, bool) {
	str := []rune(s)

	fmt.Println(str)

	l := len(str)

	if l > 5000 {
		return "", false
	}

	for i := 0; i < l / 2; i++ {
		str[i], str[l - i - 1] = str[l - i - 1], str[i]
	}

	return string(str), true
}