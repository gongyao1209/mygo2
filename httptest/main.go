package main

import (
	"fmt"
	"io"
	"net/http"
)

func main()  {
	arr()

	return
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "pong")
	})

	http.ListenAndServe(":9193", nil)
}

func arr()  {
	 //a := []string{"a", "b", "c"}

	 //a := []string{"a", "b", "c"}
	 a := make([]string, 10)
	 a = append(a, "d", "e")

	 fmt.Printf("a add %p\n", &a)
	 //fmt.Printf("b add %p\n", &b)
}