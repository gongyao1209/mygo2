package main

import (
	"io"
	"net/http"
)

func main()  {
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "ping")
	})

	http.HandleFunc("/pong", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "pong")
	})

	http.ListenAndServe(":9193", nil)
}