GO ?= go



.PHONY: default
default:
	GOOS=linux GOARCH=amd64 $(GO) build -trimpath -o release/uoss ./cmd/uoss