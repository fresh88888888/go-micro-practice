package main

import (
	"context"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"umbrella.github.com/go-micro.example/gateway/service"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	gRpcEndpoint := "localhost:8081"
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	err := service.RegisterTestServiceHandlerFromEndpoint(ctx, mux, gRpcEndpoint, opts)
	if err != nil {
		log.Fatal(err)
	}

	if err := http.ListenAndServe(":8082", mux); err != nil {
		log.Fatal(err)
	}
}
