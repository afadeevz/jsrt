# Makefile for jsrt project

# default target
.PHONY: all
all: generate

# run code generator
.PHONY: generate
generate:
	go run cmd/codegen.go > pkg/generated.go
