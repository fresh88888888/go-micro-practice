package sidecar

type JsonRequest struct {
	Jsonp  string
	Method string
	Params []*Service
	Id     int
}

func NewJsonRequest(service *Service, endpoint string) *JsonRequest {
	return &JsonRequest{Jsonp: "2.0", Method: endpoint, Params: []*Service{service}, Id: 1}
}

func RegService(service Service) error {
	return requestRegister(NewJsonRequest(&service, "Registry.Register"))
}

func UnRegService(service Service) error {
	return requestRegister(NewJsonRequest(&service, "Registry.Deregister"))
}
