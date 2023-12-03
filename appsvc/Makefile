.PHONY: gen 
gen:
	@goctl rpc protoc appsvc.proto --go_out=./types --go-grpc_out=./types --zrpc_out=.

.PHONY: run 
run:
	@go run appsvc.go -f etc/appsvc.yaml