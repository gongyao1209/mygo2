package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8}
	s1 := arr[2:5]
	s2 := s1[2:4]
	fmt.Println(s1, s2)
}