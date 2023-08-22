package impl

import (
	"context"
	"time"

	"umbrella.github.com/go-micro.example/grpc/db"
	"umbrella.github.com/go-micro.example/grpc/entry"
	"umbrella.github.com/go-micro.example/grpc/service/protos"
)

type UserService struct {
}

func (this *UserService) UserReg(ctx context.Context, modle *protos.UserModel, resp *protos.RegResponse) error {
	user := entry.User{UserName: modle.UserName, UserPwd: modle.UserPassd, UserDate: time.Now()}
	if err := db.GetDB().Create(&user).Error; err != nil {
		resp.Status = "500"
		resp.Message = err.Error()
	} else {
		resp.Status = "200"
		resp.Message = "insert user success!"
	}

	return nil
}
