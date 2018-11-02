package main

import (
	"fmt"
	"net"
	"time"
)

func handleConn(conn net.Conn)  {
	defer func() {
		conn.Close()
		fmt.Println("close conn ", conn.RemoteAddr().String())
	}()

	//conn.Write([]byte("hello socket\n"))

	for i:= 0; i < 10; i++ {
		fmt.Printf("accept i = %d\n", i)

		//设置 读取 阻塞的最迟时间
		conn.SetReadDeadline(time.Now().Add(time.Second * 10))
		var buf = make([]byte, 10)
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Printf("Read Error")
			if nerr, ok := err.(net.Error); ok && nerr.Timeout() { //如果超时 不输入数据那么直接返回
				return
			}
		}

		// 获取所传的数据
		msg := string(buf[:n])
		if msg == "close" {
			return
		}

		conn.Write([]byte(msg))
	}
}

func main()  {
	//一直在监听端口
	listener, err := net.Listen("tcp", ":12344")
	if err != nil {
		fmt.Println("Listen Err ", err)
	}

	//一直在监听端口 每次连接一次循环
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Conn Err ", err)
		}
		fmt.Println("conn Success", conn.RemoteAddr().String())

		go handleConn(conn)
	}

	//loopClient()
}

func loopClient()  {
	var cs []net.Conn

	for i := 0; i < 5; i++ {
		c := client(i)
		if c != nil {
			cs = append(cs, c)
		}
	}
}

func client(i int) net.Conn {
	conn, err := net.Dial("tcp", ":12344")
	if err != nil {
		fmt.Printf("i = %d, Conn Error %v", i, err)
	}

	fmt.Printf("i = %d, Conn Success\n", i)
	//defer conn.Close()
	return conn
}