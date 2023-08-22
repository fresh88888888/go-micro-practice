package impl

import (
	"context"
	"strconv"

	"umbrella.github.com/go-micro.example/grpc/service/protos"
)

type TestService struct {
}

func (*TestService) Call(ctx context.Context, request *protos.TestRequest, resp *protos.TestResponse) error {
	resp.Data = "test:" + strconv.Itoa(int(request.Id))

	return nil
}
