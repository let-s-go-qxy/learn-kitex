package main

import (
	"context"
	user "learnkitex/kitex_gen/api"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// UserInfo implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserInfo(ctx context.Context, req *user.UserRequest) (resp *user.UserResponse, err error) {
	// TODO: Your code here...
	msg := "ok"
	return &user.UserResponse{
		StatusCode: 0,
		StatusMsg:  &msg,
		User: &user.User{
			Id:            0,
			Name:          "name",
			FollowCount:   nil,
			FollowerCount: nil,
			IsFollow:      false,
		},
	}, nil
}
