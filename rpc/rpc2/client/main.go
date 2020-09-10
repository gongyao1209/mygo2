package main

import (
	"fmt"
	"net"
)

func main()  {
	fmt.Println(getUser())
}

func getUser() string {
	conn, _ := net.Dial("tcp", ":1239")

	conn.Write([]byte("1"))

	var u []byte = make([]byte, 100)

	conn.Read(u)

	return string(u)
}