package main

import (
	"context"
	"fmt"
	"log"

	"github.com/asim/go-micro/plugins/registry/etcd/v3"
	"github.com/asim/go-micro/v3/client"
	"github.com/asim/go-micro/v3/registry"
	"github.com/asim/go-micro/v3/selector"
	"umbrella.github.com/go-micro.example/client/http"
	"umbrella.github.com/go-micro.example/etcd/protos"
)

func main() {
	register := etcd.NewRegistry(registry.Addrs(":2379"))

	mySelector := selector.NewSelector(
		selector.Registry(register),
		selector.SetStrategy(selector.RoundRobin),
	)

	getClient := http.NewClient(client.Selector(mySelector), client.ContentType("application/json"))
	req := getClient.NewRequest("api.umbrella.myapp", "/v1/test", map[string]string{})

	rsp := new(protos.TestResp)
	err := getClient.Call(context.Background(), req, rsp)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(rsp)
}
