package main

import (
	"context"
	"fmt"

	proto "github.com/micro/examples/service/proto"
	micro "github.com/micro/go-micro"
)

type Greeter struct {}

func (g *Greeter)Hello(ctx context.Context, req *proto.HelloRequest, rsp *proto.HelloResponse) error {
	rsp.Greeting = "Hello " + req.Name
	return nil
}

func main()  {
	service := micro.NewService(
		micro.Name("gretter"),
		)

	service.Init()

	proto.RegisterGreeterHandler(service.Server(), new(Greeter))

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}

