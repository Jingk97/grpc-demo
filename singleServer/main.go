package main

import (
	"fmt"
	"google.golang.org/protobuf/proto"
	"jingpc/grpc-demo/singleServer/service"
)

// main
//
//	@Description:单点服务进行序列化和反序列化测试
func main() {
	u := &service.User{Username: "jingpc", Age: 27}
	// marshal 为序列化之后的内容
	marshal, err := proto.Marshal(u)
	if err != nil {
		panic(err)
	}

	newU := &service.User{}
	err = proto.Unmarshal(marshal, newU)
	if err != nil {
		panic(err)
	}
	fmt.Println("我是：", newU.Username, "\n年龄是：", newU.Age)
}
