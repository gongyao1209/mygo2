package main

import (
	"bytes"
	"fmt"
	"strings"
)

func Add(a, b int) int {
	return a + b
}

func ForSlice(s []string) {
	len_ := len(s)
	for i := 0; i < len_; i++ {
		_, _ = i, s[i]
	}
}

func RangeForSlice(s []string) {
	for i, v := range s {
		_, _ = i, v
	}
}

// 字符串的几种拼接
func StringPlus() string {
	var s string
	s += "昵称" + ":" + "飞雪无情" + "\n"
	s += "博客" + ":" + "http://www.flysnow.org/" + "\n"
	s += "微信公众号" + ":" + "flysnow_org"
	return s
}

func StringPlusFtm() string {
	s := fmt.Sprint("昵称", ":", "飞雪无情", "\n", "博客", ":", "http://www.flysnow.org/", "\n", "微信公众号", ":", "flysnow_org")
	return s
}

func StringPlusJoin() string {
	s := []string{"昵称", ":", "飞雪无情", "\n", "博客", ":", "http://www.flysnow.org/", "\n", "微信公众号", ":", "flysnow_org"}
	return strings.Join(s, "")
}

func StringPlusBuffer() string {
	var b bytes.Buffer
	b.WriteString("昵称")
	b.WriteString(":")
	b.WriteString("飞雪无情")
	b.WriteString("\n")
	b.WriteString("博客")
	b.WriteString(":")
	b.WriteString("http://www.flysnow.org/")
	b.WriteString("\n")
	b.WriteString("微信公众号")
	b.WriteString(":")
	b.WriteString("flysnow_org")

	return b.String()
}
