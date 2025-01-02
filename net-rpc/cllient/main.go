package main

import (
	"log"
	"net/rpc"
)

type (
	GetUserResp struct {
		ID    string `json:"id"`
		Name  string `json:"name"`
		Phone string `json:"phone"`
		Age   int    `json:"age"`
	}
	GetUserReq struct {
		ID string `json:"id"`
	}
)

func main() {
	connect, err := rpc.Dial("tcp", "localhost:8000")
	defer connect.Close()
	if err != nil {
		log.Fatal("dialing:", err)
	}
	var getReq = GetUserReq{ID: "2"}
	var getResp = GetUserResp{}
	err = connect.Call("UserService.GetUser", getReq, &getResp)
	if err != nil {
		log.Println("call error:", err)
		return
	}
	log.Println("getResp:", getResp)
}
