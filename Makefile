# note: call scripts from /scripts
.PHONY: build
build:
	go build -v ./cmd/profiles

.DEFAULT_GOAL := build