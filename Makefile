
DOCKER_DIR := ${CURDIR}/build/package

build-docker:
	docker build -t music-service -f ${DOCKER_DIR}/music-service.Dockerfile .

build-and-run: build-docker
	docker-compose up


.PHONY: build_profile
build_profile:
	go build -v ./cmd/profiles



.PHONY: clean
clean:
	rm -rf *.o


.DEFAULT_GOAL := build_profile