package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"hello-client/pb"
	"io"
	"time"
)

func main() {
	conn, err := grpc.NewClient("127.0.0.1:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	c := hello_client.NewGreeterClient(conn)
	// 简单调用
	//callSayHello(c)

	// 流式调用
	callStreamSayHello(c)
}

// 普通函数调用
func callSayHello(client hello_client.GreeterClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err := client.SayHello(ctx, &hello_client.HelloRequest{Name: "JingPC", Age: 28})
	if err != nil {
		panic(err)
	}
	fmt.Println("对面请求的结果是：", resp)
}

// 请求流式server
func callStreamSayHello(client hello_client.GreeterClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	streamResp, err := client.LotsOfSayHello(ctx, &hello_client.HelloRequest{Name: "JingPC", Age: 28})
	if err != nil {
		panic(err)
	}
	for {
		resp, err := streamResp.Recv()
		if err == io.EOF {
			break
		}
		fmt.Println("对面请求的结果是：", resp)
		if err != nil {
			panic(err)
		}
	}
}
