.PHONY: gen 
gen:
	@goctl rpc protoc ${name}.proto --go_out=./types --go-grpc_out=./types --zrpc_out=.

.PHONY: run 
run:
	@go run ${name}.go -f etc/${name}.yaml