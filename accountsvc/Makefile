CURRENT_DIR := $(shell basename "$(CURDIR)")

.PHONY: gen 
g:
	@goctl rpc protoc $(CURRENT_DIR).proto --go_out=./types --go-grpc_out=./types --zrpc_out=. --style go_zero -m
d:
    ifeq (,$(wildcard ./Dockerfile))
	@goctl docker 
    endif

.PHONY: run 
r:
	@go run $(CURRENT_DIR).go -f etc/$(CURRENT_DIR).yaml
e:
	@echo "当前目录名称: $(CURRENT_DIR)"
