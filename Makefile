
DOCKER_DIR := ${CURDIR}/build/package

build-all:
	go run cmd/sessions/main.go

build-docker:
	docker build -t music-service -f ${DOCKER_DIR}/music-service.Dockerfile .

build-and-run: build-docker
	docker-compose up

test-pr:
	go test -v -coverprofile cover.out ./...
	go tool cover -html=cover.out -o cover.html


.PHONY: build_profile
build_profile:
	go build -v ./cmd/profiles



.PHONY: clean
clean:
	rm -rf *.o


.DEFAULT_GOAL := build_profile