
DOCKER_DIR := ${CURDIR}/build/package

build-all:
	go run cmd/sessions/main.go

build-docker:
	docker build --no-cache --network host -t sessions-service -f ${DOCKER_DIR}/sessions-service.Dockerfile .
	docker build --no-cache --network host -t music-service -f ${DOCKER_DIR}/music-service.Dockerfile .
	docker build --no-cache --network host -t profiles-service -f ${DOCKER_DIR}/profiles-service.Dockerfile .

build-and-run: build-docker
	docker-compose up

test-pr:
	go test -v -coverprofile cover.out ./...
	go tool cover -html=cover.out -o cover.html


coverage:
	go test -covermode=atomic -coverpkg=./... -coverprofile=cover ./...
	cat cover | fgrep -v "mock" | fgrep -v "pb.go" | fgrep -v "easyjson" | fgrep -v "start.go" > cover2
	go tool cover -func=cover2


clear-dockers:
	docker-compose down
	docker system prune -a
	docker volume prune

linter:
	 golangci-lint run

.PHONY: build_profile
build_profile:
	go build -v ./cmd/profiles



.PHONY: clean
clean:
	rm -rf *.o


.DEFAULT_GOAL := build_profile