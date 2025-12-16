package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"hello-server/pb"
	"net"
)

type server struct {
	hello_server.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, request *hello_server.HelloRequest) (*hello_server.HelloResponse, error) {
	response := fmt.Sprintf("Hello %s,你的年龄是%d岁～", request.Name, request.Age)
	return &hello_server.HelloResponse{Msg: response}, nil
}

func (s *server) LotsOfSayHello(req *hello_server.HelloRequest, stream hello_server.Greeter_LotsOfSayHelloServer) (err error) {
	hellos := []string{
		"hello",
		"你好",
		"안녕하세요",
	}
	for _, hello := range hellos {
		resp := &hello_server.HelloResponse{
			Msg: fmt.Sprintf("%s,%s！您现在是：%d岁！", hello, req.Name, req.Age),
		}
		if err := stream.Send(resp); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	hello_server.RegisterGreeterServer(s, &server{})
	err = s.Serve(l)
	if err != nil {
		panic(err)
	}
}
