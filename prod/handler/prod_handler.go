package handler

import (
	"log"

	"github.com/gin-gonic/gin"
	"umbrella.github.com/go-micro.example/prod/service"
)

type Route struct{}

func NewRoute() *Route {
	return &Route{}
}

func (this *Route) InitRoute(router *gin.RouterGroup) {
	router.POST("/prods", this.prods)
}

func (this *Route) prods(ctx *gin.Context) {
	var pr service.ProdRequest
	err := ctx.Bind(&pr)
	if err != nil || pr.Size <= 0 {
		log.Println(err)
		pr = service.ProdRequest{Size: 2}
	}

	ctx.JSON(200, gin.H{"code": 200, "data": service.NewProdList(pr.Size)})
}
