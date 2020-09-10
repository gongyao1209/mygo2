package main

import (
	"encoding/json"
	"fmt"
	"gocode/mygo2/rpc/rpc1/server/people"
	"net"
	"strings"
)

func main()  {
	listen, err := net.Listen("tcp", ":1239")
	if err != nil {
		fmt.Println(err.Error())
	}

	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println(err.Error())
		}

		go func() {
			var buf = make([]byte, 10)
			conn.Read(buf)

			//conn.Write([]byte(msg)) //写数据
			msg := strings.Trim(string(buf), " ")

			fmt.Println(msg, "|", strings.Compare(msg, "user"))

			if strings.Compare(msg, "user") == 1 {
				people.FindUserById()
				b, _ := json.Marshal(people.FindUserById())
				conn.Write(b)
			} else {
				conn.Write([]byte(accept(msg)))
			}
		}()

	}

}

func accept(msg string) string {
	str := "accept :" + msg

	fmt.Println(str)

	return str
}