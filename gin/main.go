package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"umbrella.github.com/go-micro.example/sidecar"
)

func main() {
	gin_router := gin.Default()
	gin_group := gin_router.Group("/v1")
	gin_group.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"data": "teste kkk"})
	})

	http_server := http.Server{
		Addr:    ":8082",
		Handler: gin_router,
	}
	myservice := sidecar.NewService("api.umbrella.com.test")
	myservice.AddNode("test"+uuid.New().String(), "8082", "198.19.37.126:8082")
	handler := make(chan error)
	go func() {
		handler <- http_server.ListenAndServe()
	}()

	go func() {
		notify := make(chan os.Signal)
		signal.Notify(notify, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL)
		handler <- fmt.Errorf("%s", <-notify)
	}()

	go func() {
		err := sidecar.RegService(*myservice)
		if err != nil {
			log.Fatal(err)
		}
	}()
	getHandler := <-handler
	fmt.Println(getHandler.Error())

	err := sidecar.UnRegService(*myservice)
	if err != nil {
		log.Fatal(err)
	}

	err = http_server.Shutdown(context.Background())
	if err != nil {
		log.Fatal(err)
	}
}
