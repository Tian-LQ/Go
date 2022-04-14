package main

import (
	"context"
	"firstProject/internal/grpc_class/pb/student"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"net"
)

var (
	port = flag.Int("port", 50052, "The server port")
)

type server struct {
	student.UnimplementedSearchServiceServer
}

func (s *server) Search(ctx context.Context, req *student.StudentReq) (res *student.StudentRes, err error) {
	name := req.GetName()
	res = &student.StudentRes{Name: fmt.Sprintf("receive message form caller [%s]", name)}
	err = nil
	return
}

func main() {
	flag.Parse()
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		panic(err)
	}
	// 实例化gRPC Server
	s := grpc.NewServer()
	student.RegisterSearchServiceServer(s, &server{})
	// 建立监听
	s.Serve(l)
}
