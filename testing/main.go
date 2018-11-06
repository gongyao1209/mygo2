package main

func Add(a,b int) int{
	return a+b
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