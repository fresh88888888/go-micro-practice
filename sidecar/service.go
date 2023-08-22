package sidecar

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Service struct {
	Name string
	Node []*ServiceNode
}

type ServiceNode struct {
	Id      string
	Port    string
	Address string
}

func NewService(name string) *Service {
	return &Service{Name: name, Node: make([]*ServiceNode, 0)}
}

func NewServiceNode(id string, port string, address string) *ServiceNode {
	return &ServiceNode{Id: id, Port: port, Address: address}
}

func (this *Service) AddNode(id string, port string, address string) {
	this.Node = append(this.Node, NewServiceNode(id, port, address))
}

var RegisterUri = "198.19.37.126:2379"

func requestRegister(jsonRequest *JsonRequest) error {
	b, err := json.Marshal(jsonRequest)
	if err != nil {
		log.Fatal(err)
		return err
	}

	rsp, err := http.Post(RegisterUri, "application/json", bytes.NewReader(b))
	if err != nil {
		return err
	}
	defer rsp.Body.Close()

	res, err := io.ReadAll(rsp.Body)
	if err != nil {
		return err
	}

	fmt.Println(string(res))

	return nil
}
