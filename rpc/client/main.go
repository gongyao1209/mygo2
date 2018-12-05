package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main()  {
	conn, err := net.Dial("tcp", "localhost:12345")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))

	var reply string
	str := "{\"a\":\"b\"}"
	err = client.Call("HelloService.Hello", str, &reply)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(reply)
}