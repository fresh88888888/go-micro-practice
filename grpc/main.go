package main

import (
	"log"

	"github.com/asim/go-micro/plugins/registry/etcd/v3"
	httpServer "github.com/asim/go-micro/plugins/server/http/v3"
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/registry"
	"github.com/asim/go-micro/v3/server"
	"github.com/gin-gonic/gin"
	_ "umbrella.github.com/go-micro.example/grpc/db"
)

func main() {

	register := etcd.NewRegistry(registry.Addrs(":2379"))

	gin_router := gin.Default()
	gin_group := gin_router.Group("/v1")
	gin_group.POST("/test", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"data": "teste kkk"})
	})

	srv := httpServer.NewServer(
		server.Address(":8082"),
	)

	mysservice := micro.NewService(
		micro.Server(srv),
		micro.Name("api.umbrella.myapp"),
		micro.Registry(register),
	)

	//protos.RegisterUserServiceHandler(myservice.Server(), new(impl.UserService))

	hd := srv.NewHandler(gin_router)
	if err := srv.Handle(hd); err != nil {
		log.Fatal(err)
	}
	mysservice.Init()

	if err := mysservice.Run(); err != nil {
		log.Fatal(err)
	}

}
