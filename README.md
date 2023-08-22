
# go-micro-example

1. go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
2. go install github.com/go-micro/generator/cmd/protoc-gen-micro@latest
3. export PATH="$PATH:$(go env GOPATH)/bin"
4. protoc --micro_out=.. --go_out=.. prods.proto
5. go install github.com/favadi/protoc-go-inject-tag@latest
6. protoc --micro_out=.. --go_out=.. prods.proto && protoc-go-inject-tag -input="../*.pb.go"
7. protoc --micro_out=.. --go_out=.. --go-grpc_out=.. --grpc-gateway_out=logtostderr=true:../../service *.proto && protoc-go-inject-tag -input="./*.pb.go"

run

- go run main.go --server_address :8081

protoc --micro_out=.. --go_out=.. --go-grpc_out=..  *.proto && protoc-go-inject-tag -input="./*.pb.go"
