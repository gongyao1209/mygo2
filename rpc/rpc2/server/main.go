package main

import (
	"encoding/json"
	"fmt"
	"net"
)

func main()  { // 服务端需要一直监听，客户端只需要一次
	lis, _ := net.Listen("tcp", ":1239")

	i := 1
	for {
		fmt.Println("ii := ", i)
		conn, _ := lis.Accept()
		fmt.Println("i := ", i)

		var buf = make([]byte, 10)

		go func() {
			conn.Read(buf)
			fmt.Println(string(buf))

			u := User{Id:i, Name:"gongyao"}
			b, _ := json.Marshal(u)
			conn.Write(b)
		}()

		i++
	}
}

type User struct {
	Id int `json:"id1"`
	Name string `json:"name1"`
}