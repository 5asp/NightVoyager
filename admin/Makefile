CURRENT_DIR := $(shell basename "$(CURDIR)")
.PHONY: gen 
gen:
	@goctl api go --api $(CURRENT_DIR).api --dir . --style go_zero
docker:
    ifeq (,$(wildcard ./Dockerfile))
	@goctl docker 
    endif
.PHONY: run 
run:
	@go run $(CURRENT_DIR).go -f etc/$(CURRENT_DIR)-api.yaml
swagger:
	@goctl api plugin -plugin goctl-swagger="swagger -filename $(CURRENT_DIR).json" -api $(CURRENT_DIR).api -dir .