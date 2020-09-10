package main

import (
	"fmt"
	"net"
)

func main()  {

	conn, err := net.Dial("tcp", ":1239")
	if err!= nil {
		fmt.Println(err.Error())
		return
	}

	buf := []byte("user")
	//buf := make([]byte, 0)
	//buf = append(buf, 12)
	//buf = append(buf, 23)
	//buf = append(buf, 34)

	fmt.Println(string(buf))

	n, err := conn.Write(buf)
	fmt.Println(n, err)

	var nbuf = make([]byte, 100)
	conn.Read(nbuf)
	fmt.Println(string(nbuf))
}