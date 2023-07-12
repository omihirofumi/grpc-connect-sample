.PHONY: buflint
buflint:
	@( cd proto; buf lint; )

.PHONY: bufform
bufform:
	@( cd proto; buf format --write; )

.PHONY: bufgen
bufgen: buflint bufform
	@buf generate

.PHONY: start
start:
	@go run ./cmd/server/main.go