package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type HelloService struct{}

func (h *HelloService) Hello(request string, reply *string) error {
	*reply = "Hello " + request
	fmt.Println(*reply)
	return nil
}

func (h *HelloService) Hello2(request string, reply *string) error {
	*reply = "shabia " + request
	fmt.Println(*reply)
	return nil
}

func main() {
	rpc.RegisterName("HelloService", new(HelloService))

	listener, err := net.Listen("tcp", ":12345")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}

	//一直在监听端口 每次连接一次循环
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Conn Err ", err)
		}
		fmt.Println("conn Success", conn.RemoteAddr().String())

		// 里面是连接需要做什么
		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
		//handleConn(conn)
	}

	//conn, err := listener.Accept()
	//if err != nil {
	//	log.Fatal("Accept error:", err)
	//}
	//
	//for true {
	//	go rpc.ServeConn(conn)
	//}
}
