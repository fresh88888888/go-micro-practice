package wrapper

import (
	"context"
	"fmt"

	"github.com/asim/go-micro/v3/client"
	"github.com/asim/go-micro/v3/metadata"
)

type logWrapper struct {
	client.Client
}

func NewLogWrapper() client.Wrapper {
	return func(c client.Client) client.Client {
		return &logWrapper{c}
	}
}
func (this *logWrapper) Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error {
	md, _ := metadata.FromContext(ctx)
	fmt.Printf("[Log Wrapper] ctx: %v service: %s method: %s\n", md, req.Service(), req.Endpoint())
	return this.Client.Call(ctx, req, rsp, opts...)
}
