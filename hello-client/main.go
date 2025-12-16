package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"hello-client/pb"
	"io"
	"log"
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

	// 流式调用(server)
	//callStreamSayHello(c)

	// 流式调用（client）
	//callStreamSendHello(c)

	// 双流式调用
	callStreamHello(c)
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

// 流式请求client
func callStreamSendHello(client hello_client.GreeterClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	streamClient, err := client.LotsOfSendHello(ctx)
	if err != nil {
		panic(err)
	}
	names := []string{"JingPC", "张三", "李四"}
	for _, name := range names {
		// 仅发送
		err = streamClient.Send(&hello_client.HelloRequest{Name: name})
		if err != nil {
			panic(err)
		}
	}
	// 关闭并且获得回复，因为对方只能一次性回复所以只能一次性处理回复
	resp, err := streamClient.CloseAndRecv()
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}

// 双向流式请求client
func callStreamHello(client hello_client.GreeterClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	streamClient, err := client.StreamSayHello(ctx)
	if err != nil {
		panic(err)
	}
	waitc := make(chan struct{})
	go func() {
		for {
			// 接收服务端返回的响应
			resp, err := streamClient.Recv()
			if err == io.EOF {
				// read done.
				close(waitc)
				return
			}
			if err != nil {
				log.Fatalf("c.BidiHello stream.Recv() failed, err: %v", err)
			}
			fmt.Printf("响应回答为：%s\n", resp.Msg)
		}
	}()
	names := []*hello_client.HelloRequest{
		{Name: "JingPC", Age: 28},
		{Name: "张三", Age: 30},
		{Name: "李四", Age: 40},
	}
	for _, man := range names {
		err = streamClient.Send(man)
		if err != nil {
			panic(err)
		}
	}
	streamClient.CloseSend()
	<-waitc
}
