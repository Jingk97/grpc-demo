package main

import (
	"google.golang.org/grpc"
	"jingpc/grpc-demo/csArch/service"
	"log"
	"net"
)

func main() {
	rpcServer := grpc.NewServer()
	service.RegisterProdServiceServer(rpcServer, service.ProductService)

	listener, err := net.Listen("tcp", ":8002")
	if err != nil {
		log.Fatal("服务监听端口失败", err)
	}
	_ = rpcServer.Serve(listener)
}
