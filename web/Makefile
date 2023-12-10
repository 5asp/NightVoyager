CURRENT_DIR := $(shell basename "$(CURDIR)")
.PHONY: g 
g:
	@goctl api go --api $(CURRENT_DIR).api --dir . --style go_zero
d:
    ifeq (,$(wildcard ./Dockerfile))
	@goctl docker 
    endif
.PHONY: run 
r:
	@go run $(CURRENT_DIR).go -f etc/$(CURRENT_DIR)-api.yaml
f:
	@goctl api format --dir $(CURRENT_DIR).api
s:
	@goctl api plugin -plugin goctl-swagger="swagger -filename $(CURRENT_DIR).json" -api $(CURRENT_DIR).api -dir .