package impl

import (
	"context"
	"strconv"
	"time"

	"umbrella.github.com/go-micro.example/product/service/protos"
)

type ProdService struct {
}

func (this *ProdService) GetProdDetail(ctx context.Context, request *protos.ProdRequest, response *protos.ProdDetailResponse) error {
	time.Sleep(time.Second * 2)
	response.Data = NewProd(request.ProId, "tumaoto of food")

	return nil
}

func NewProd(id int32, name string) *protos.ProdModel {
	return &protos.ProdModel{ProdID: id, ProdName: name}
}

func (this *ProdService) GetProdsList(ctx context.Context, request *protos.ProdRequest, response *protos.ProdListResponse) error {
	time.Sleep(time.Second * 2)
	models := make([]*protos.ProdModel, 0)
	var i int32
	for i = 0; i < request.Size; i++ {
		models = append(models, &protos.ProdModel{ProdID: 100 + i, ProdName: "prod-" + strconv.Itoa(int(i))})
	}
	response.Data = models

	return nil
}
