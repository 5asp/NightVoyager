.PHONY: gen 
gen:
	@goctl rpc protoc queuesvc.proto --go_out=./types --go-grpc_out=./types --zrpc_out=.

.PHONY: run 
run:
	@go run queuesvc.go -f etc/queuesvc.yaml