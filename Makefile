.PHONY: build
build:
	go build -v  -o "./artifacts/bin/client" ./cmd/httpserver

.DEFAULT_GOAL := build
