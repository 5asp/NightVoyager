.PHONY: gen 
gen:
	@goctl rpc protoc appsvc.proto --go_out=./ --go-grpc_out=./ --zrpc_out=.
.PHONY: run 
run:
	@go run send.go -f etc/send.yaml