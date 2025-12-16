package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"hello-server/pb"
	"io"
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

func (s *server) LotsOfSendHello(stream hello_server.Greeter_LotsOfSendHelloServer) error {
	respMsg := "你们好："
	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&hello_server.HelloResponse{
				Msg: respMsg,
			})
		}
		if err != nil {
			return err
		}
		respMsg += resp.Name + ","
	}
}

func (s *server) StreamSayHello(stream hello_server.Greeter_StreamSayHelloServer) error {
	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		err = stream.Send(&hello_server.HelloResponse{
			Msg: fmt.Sprintf("你好啊：%s，您的年龄是%d岁！", resp.Name, resp.Age),
		})
		if err != nil {
			return err
		}
	}
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
