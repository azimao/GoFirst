package main

import (
	"net"
	"net/rpc"
)

type UserService struct {
}

func (this *UserService) GetUserName(userId int, username *string) error {
	if userId == 101 {
		*username = "azimao"
	} else {
		*username = "tc"
	}
	return nil
}

func main() {
	lis, _ := net.Listen("tcp", ":8082")
	rpc.Register(new(UserService))
	client, _ := lis.Accept()
	rpc.ServeConn(client)
}
