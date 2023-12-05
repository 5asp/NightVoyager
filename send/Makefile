.PHONY: gen 
gen:
	@goctl api go --api send.api --dir .
	@goctl docker
.PHONY: run 
run:
	@go run send.go -f etc/send-api.yaml