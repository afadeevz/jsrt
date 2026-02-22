# Makefile for jsrt project

# default target
.PHONY: all
all: fmt generate

# run code generator
.PHONY: generate
generate:
	go run cmd/codegen.go > pkg/generated.go

# formatting and vetting
.PHONY: fmt
fmt:
	gofmt -s -w $(shell find . -name '*.go' -not -path './pkg/generated.go')
	go vet ./...

.PHONY: getDocs
getDocs:
	git clone git@github.com:mdn/content.git