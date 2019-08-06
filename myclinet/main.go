package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	/*	conn, err := net.Dial("tcp", "127.0.0.1:8099")

		if err != nil {
			fmt.Println(err)
			return
		}
		defer conn.Close()

		conn.Write([]byte("I am azimao"))*/

	client, _ := rpc.Dial("tcp", "127.0.0.1:8082")
	username := ""
	err := client.Call("UserService.GetUserName", 111, &username)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(username)
}
