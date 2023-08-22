package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/asim/go-micro/plugins/registry/consul/v3"
	httpServer "github.com/asim/go-micro/plugins/server/http/v3"
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/client"
	"github.com/asim/go-micro/v3/registry"
	"github.com/asim/go-micro/v3/selector"
	"github.com/asim/go-micro/v3/server"
	myhttp "umbrella.github.com/go-micro.example/client/http"
	"umbrella.github.com/go-micro.example/client/route"
	"umbrella.github.com/go-micro.example/client/service/protos"
	"umbrella.github.com/go-micro.example/client/wrapper"
	"umbrella.github.com/go-micro.example/prod/models"
)

func main() {

	// register := consul.NewRegistry(registry.Addrs("127.0.0.1:8500"))
	// c_selector := selector.NewSelector(
	// 	selector.SetStrategy(selector.RoundRobin),
	// 	selector.Registry(register),
	// )

	invokeGrpcClient()
	// getService, err := register.GetService("prod-service")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// next := selector.RoundRobin(getService)
	// node, err := next()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// res, err := invokeNativeClient(node.Address, "/v1/prods", "POST")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// log.Println(res)
}

func invokeNativeClient(addr, path, method string) (string, error) {
	var url string = fmt.Sprint("http://", addr, path)
	req, _ := http.NewRequest(method, url, nil)
	client := http.DefaultClient
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	buf, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return string(buf), nil
}

func invokeClient(s selector.Selector) {
	myclient := myhttp.NewClient(
		client.Selector(s),
		client.ContentType("application/json"),
	)
	//req := myclient.NewRequest("prod-service", "/v1/prods", map[string]int{"size": 4})
	req := myclient.NewRequest("prod-service", "/v1/prods", models.ProdRequest{Size: 3})
	//var rsp map[string]interface{}
	var rsp models.ProdListResponse
	err := myclient.Call(context.Background(), req, &rsp)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(rsp.GetData())
}

func invokeGrpcClient() {
	register := consul.NewRegistry(registry.Addrs("127.0.0.1:8500"))

	srv := httpServer.NewServer(
		server.Name("http-product-service"),
		server.Address(":8082"),
	)

	service := micro.NewService(
		micro.Name("productservice.client"),
		micro.Server(srv),
		micro.Registry(register),
		micro.WrapClient(wrapper.NewLogWrapper()),
		micro.WrapClient(wrapper.NewProdWrapper()),
	)

	productService := protos.NewProdService("product-service", service.Client())
	gin_route := route.NewProductRoute(productService).InitRoute()

	hd := srv.NewHandler(gin_route)
	if err := srv.Handle(hd); err != nil {
		log.Fatal(err)
	}

	service.Init()
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
