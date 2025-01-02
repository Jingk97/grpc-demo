package main

import (
	"errors"
	"log"
	"net"
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
type UserService struct {
}

func (us *UserService) GetUser(req GetUserReq, resp *GetUserResp) error {
	if u, ok := Users[req.ID]; ok {
		*resp = GetUserResp{
			ID:    u.Name,
			Name:  u.Name,
			Phone: u.Phone,
			Age:   u.Age,
		}
		return nil
	}
	return errors.New("查无此人")
}
func main() {
	userService := new(UserService)
	rpc.Register(userService)
	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatal("listen error:", err)
	}
	log.Println("listen success")
	for {
		connect, err := listener.Accept()
		if err != nil {
			log.Println("accept error:", err)
			continue
		}
		go rpc.ServeConn(connect)
	}
}
