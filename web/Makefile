CURRENT_DIR := $(shell basename "$(CURDIR)")
.PHONY: gen 
gen:
	@goctl api go --api $(CURRENT_DIR).api --dir .
.PHONY: run 
run:
	@go run $(CURRENT_DIR).go -f etc/$(CURRENT_DIR)-api.yaml