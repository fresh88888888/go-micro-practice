package main

import (
	"log"

	"github.com/asim/go-micro/plugins/registry/consul/v3"
	httpServer "github.com/asim/go-micro/plugins/server/http/v3"
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/registry"
	"github.com/asim/go-micro/v3/server"
	"github.com/gin-gonic/gin"
	"umbrella.github.com/go-micro.example/prod/handler"
)

const (
	server_name = "http:prod_server"
)

func main() {
	consul_reg := consul.NewRegistry(registry.Addrs(":8500"))

	gin.SetMode(gin.ReleaseMode)
	gin_route := gin.New()
	gin_route.Use(gin.Recovery())
	gin_group := gin_route.Group("/v1")

	prod := handler.NewRoute()
	prod.InitRoute(gin_group)

	srv := httpServer.NewServer(
		server.Name(server_name),
		server.Address("8082"),
	)

	hd := srv.NewHandler(gin_route)
	if err := srv.Handle(hd); err != nil {
		log.Fatal(err)
	}

	service := micro.NewService(
		micro.Server(srv),
		micro.Name("prod-service"),
		micro.Registry(consul_reg),
	)

	service.Init()

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
