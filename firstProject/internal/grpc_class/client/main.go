package main

import (
	"context"
	hello_grpc "firstProject/internal/grpc_class/pb"
	"fmt"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:8888", grpc.WithInsecure())
	defer conn.Close()
	if err != nil {
		panic(err)
	}
	client := hello_grpc.NewHelloGRPCClient(conn)
	res, err := client.SayHello(context.Background(), &hello_grpc.Req{Message: "Hello again [client]"})
	if err != nil {
		panic(err)
	}
	fmt.Println(res.GetMessage())

}
