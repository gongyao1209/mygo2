package main

import (
	"log"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request)  {
}

func main()  {
	
	http.HandleFunc("/", sayHello)

	err := http.ListenAndServe(":9091", nil)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}