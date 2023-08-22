package wrapper

import (
	"context"
	"strconv"

	"github.com/afex/hystrix-go/hystrix"
	"github.com/asim/go-micro/v3/client"
	"umbrella.github.com/go-micro.example/client/service/protos"
)

type ProdWrapper struct {
	client.Client
}

func DefaultProds(rsp interface{}) {
	models := make([]*protos.ProdModel, 0)
	var i int32
	for i = 0; i < 5; i++ {
		models = append(models, &protos.ProdModel{ProdID: 20 + i, ProdName: "prod-" + strconv.Itoa(int(i))})
	}
	response := rsp.(*protos.ProdListResponse)
	response.Data = models
}

func NewProdWrapper() client.Wrapper {
	return func(c client.Client) client.Client {
		return &ProdWrapper{c}
	}
}
func (this *ProdWrapper) Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error {
	cmd_name := req.Service() + "." + req.Endpoint()
	hystrix.ConfigureCommand(cmd_name, hystrix.CommandConfig{
		Timeout:                1000,
		RequestVolumeThreshold: 2,
		ErrorPercentThreshold:  50,
		SleepWindow:            5000,
	})

	return hystrix.Do(cmd_name, func() error {
		return this.Client.Call(ctx, req, rsp, opts...)
	}, func(e error) error {
		defaultRespData(rsp)
		return nil
	})
}

func defaultRespData(rsp interface{}) {
	switch t := rsp.(type) {
	case *protos.ProdListResponse:
		DefaultProds(rsp)
	case *protos.ProdDetailResponse:
		t.Data = &protos.ProdModel{ProdID: 12, ProdName: "default product 0"}
	}
}
