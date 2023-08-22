package main

import (
	"log"
	"time"

	"github.com/asim/go-micro/plugins/registry/consul/v3"
	httpServer "github.com/asim/go-micro/plugins/server/http/v3"
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/registry"
	"github.com/asim/go-micro/v3/server"
	"github.com/gin-gonic/gin"
	"umbrella.github.com/go-micro.example/home/handler"
)

const (
	server_name = "http:home_server"
)

func main() {
	consul_reg := consul.NewRegistry(func(o *registry.Options) {
		o.Addrs = []string{"http://198.19.37.126:8500"}
	})

	gin.SetMode(gin.ReleaseMode)
	gin_route := gin.New()
	gin_route.Use(gin.Recovery())

	home := handler.NewHome()
	home.InitRoute(gin_route)

	srv := httpServer.NewServer(
		server.Name(server_name),
		server.Address(":8083"),
	)

	hd := srv.NewHandler(gin_route)
	if err := srv.Handle(hd); err != nil {
		log.Fatal(err)
	}

	service := micro.NewService(
		micro.Server(srv),
		micro.Name("home-service"),
		micro.Version("latest"),
		micro.Registry(consul_reg),
		micro.RegisterInterval(time.Second*10),
		micro.RegisterTTL(time.Second*20),
	)

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
