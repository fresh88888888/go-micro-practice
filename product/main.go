package main

import (
	"log"

	"github.com/asim/go-micro/plugins/registry/consul/v3"
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/registry"
	"umbrella.github.com/go-micro.example/product/service/impl"
	"umbrella.github.com/go-micro.example/product/service/protos"
)

func main() {
	consul_reg := consul.NewRegistry(registry.Addrs(":8500"))

	service := micro.NewService(
		micro.Name("product-service"),
		micro.Address(":8083"),
		micro.Registry(consul_reg),
	)

	service.Init()
	err := protos.RegisterProdServiceHandler(service.Server(), new(impl.ProdService))
	if err != nil {
		log.Fatal(err)
	}

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
