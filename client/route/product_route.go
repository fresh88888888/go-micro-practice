package route

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"umbrella.github.com/go-micro.example/client/handler"
	"umbrella.github.com/go-micro.example/client/service/protos"
)

type ProductRoute struct {
	service protos.ProdService
}

func NewProductRoute(service protos.ProdService) *ProductRoute {
	return &ProductRoute{service}
}

func ResetContext(service protos.ProdService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Keys = make(map[string]interface{})
		ctx.Keys["prodservice"] = service
		ctx.Next()
	}
}

func ErrorWrapper() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				ctx.JSON(500, gin.H{"status": fmt.Sprintf("s%\n", r)})
				ctx.Abort()
			}
		}()

		ctx.Next()
	}
}

func (this *ProductRoute) InitRoute() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	gin_route := gin.New()
	gin_route.Use(ResetContext(this.service), ErrorWrapper())
	gin_group := gin_route.Group("/v1")
	gin_group.POST("/prods", handler.ProdListHandle)
	gin_group.GET("/prods/:pid", handler.ProdDetailHandle)
	return gin_route
}
