.PHONY: gen 
gen:
	@goctl api go --api ${name}.api --dir .
.PHONY: run 
run:
	@go run ${name}.go -f etc/${name}-api.yaml