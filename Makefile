
DOCKER_DIR := ${CURDIR}/build/package

build-docker:
	docker build -t music-service -f ${DOCKER_DIR}/music-service.Dockerfile .

build-and-run: build-docker
	docker-compose up