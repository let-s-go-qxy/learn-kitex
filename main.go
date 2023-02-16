package main

import (
	"github.com/cloudwego/kitex/server"
	user "learnkitex/kitex_gen/api/userservice"
	"log"
	"net"
)

func main() {
	addr, _ := net.ResolveTCPAddr("tcp", ":8081")
	var opts []server.Option
	opts = append(opts, server.WithServiceAddr(addr))

	svr := user.NewServer(new(UserServiceImpl), opts...)

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
