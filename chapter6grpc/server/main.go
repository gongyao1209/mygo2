package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "gocode/mygo2/chapter6grpc/pb"
)
// SimpleService 定义我们的服务
type SimpleService struct{}

// Route 实现Route方法
func (s *SimpleService) Route(ctx context.Context, req *pb.SimpleRequest) (*pb.SimpleResponse, error) {
	a := int32(200)
	str := "hello " + string(req.Data)
	res := pb.SimpleResponse{
		Code:  a,
		Value: str,
	}
	return &res, nil
}

func (s *SimpleService) GetSum(ctx context.Context, req *pb.SumRequest) (*pb.SumResponse, error)  {
	log.Println(req)
	sum := req.Number1 + req.Number2
	res := pb.SumResponse{Sum: sum}
	return &res, nil
}


const (
	// Address 监听地址
	Address string = ":8000"
	// Network 网络通信协议
	Network string = "tcp"
)

func main() {
	// 监听本地端口
	listener, err := net.Listen(Network, Address)
	if err != nil {
		log.Fatalf("net.Listen err: %v", err)
	}
	log.Println(Address + " net.Listing...")
	// 新建gRPC服务器实例
	grpcServer := grpc.NewServer()

	// 在gRPC服务器注册我们的服务
	pb.RegisterSimpleServer(grpcServer, &SimpleService{})

	//用服务器 Serve() 方法以及我们的端口信息区实现阻塞等待，直到进程被杀死或者 Stop() 被调用
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalf("grpcServer.Serve err: %v", err)
	}
}
