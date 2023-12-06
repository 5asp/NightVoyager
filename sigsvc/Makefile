CURRENT_DIR := $(shell basename "$(CURDIR)")

.PHONY: gen 
gen:
	@goctl rpc protoc $(CURRENT_DIR).proto --go_out=./types --go-grpc_out=./types --zrpc_out=. --style go_zero
docker:
    ifeq (,$(wildcard ./Dockerfile))
	@goctl docker 
    endif

.PHONY: run 
run:
	@go run $(CURRENT_DIR).go -f etc/$(CURRENT_DIR).yaml
echo:
	@echo "当前目录名称: $(CURRENT_DIR)"
