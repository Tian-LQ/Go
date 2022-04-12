package main

import (
	"context"
	hello_grpc "firstProject/internal/grpc_class/pb"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"net"
)

var (
	port = flag.Int("port", 8888, "The server port")
)

type server struct {
	hello_grpc.UnimplementedHelloGRPCServer
}

func (s *server) SayHello(ctx context.Context, req *hello_grpc.Req) (res *hello_grpc.Res, err error) {
	fmt.Println(req.GetMessage())
	return &hello_grpc.Res{Message: "Hello, again [grpc]"}, nil
}

func main() {
	flag.Parse()
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		panic(err)
	}
	// 注册服务
	s := grpc.NewServer()
	hello_grpc.RegisterHelloGRPCServer(s, &server{})
	// 建立监听
	s.Serve(l)
}
