package main

import (
	"context"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
	"learnkitex/kitex_gen/api"
	user "learnkitex/kitex_gen/api/userservice"
	"log"
	"time"
)

func main() {
	c, err := user.NewClient("example", client.WithHostPorts("0.0.0.0:8888"))
	if err != nil {
		log.Fatal(err)
	}
	req := &api.UserRequest{
		UserId: 0,
		Token:  "",
	}
	resp, err := c.UserInfo(context.Background(), req, callopt.WithRPCTimeout(3*time.Second))
	if err != nil {
		log.Fatal(err)
	}
	log.Println(resp)
}
