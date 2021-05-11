package main

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"time"
)

func main()  {

	net.Listen("tcp", ":9194")

	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("begin...")
		time.Sleep(time.Second * 10)
		io.WriteString(w, "ping")
	})

	http.HandleFunc("/pong", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "pong")
	})

	http.ListenAndServe(":9193", nil)
}