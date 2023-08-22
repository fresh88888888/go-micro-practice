package handler

import (
	"context"

	"github.com/gin-gonic/gin"
	"umbrella.github.com/go-micro.example/client/service/protos"
)

func ProdListHandle(ctx *gin.Context) {
	var prodReq protos.ProdRequest
	err := ctx.BindJSON(&prodReq)
	prodService := ctx.Keys["prodservice"].(protos.ProdService)
	if err != nil {
		ctx.JSON(500, gin.H{"status": err.Error()})
	} else {

		resp, _ := prodService.GetProdsList(context.Background(), &prodReq)
		ctx.JSON(200, gin.H{"data": resp.GetData()})

		// hystrix.ConfigureCommand("getProds", hystrix.CommandConfig{
		// 	Timeout: 1000,
		// 	//MaxConcurrentRequests:  10,
		// 	//RequestVolumeThreshold: 20,
		// })

		// var resp *protos.ProdListResponse
		// hystrix.Do("getProds", func() error {
		// 	resp, err = prodService.GetProdsList(context.Background(), &prodReq)
		// 	return err
		// }, func(e error) error {
		// 	resp, err = DefaultProds()
		// 	return err
		// })

		// if err != nil {
		// 	ctx.JSON(500, gin.H{"status": err.Error()})
		// } else {
		// 	ctx.JSON(200, gin.H{"data": resp.GetData()})
		// }
	}
}

func ProdDetailHandle(ctx *gin.Context) {
	var prodReq protos.ProdRequest
	err := ctx.BindUri(&prodReq)
	if err != nil {
		panic(err)
	}
	prodService := ctx.Keys["prodservice"].(protos.ProdService)
	resp, _ := prodService.GetProdDetail(context.Background(), &prodReq)
	ctx.JSON(200, gin.H{"data": resp.GetData()})
}
