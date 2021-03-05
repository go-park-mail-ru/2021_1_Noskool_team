FROM golang:latest


RUN go version
ENV GOPATH=/

COPY ./ ./

RUN go build -o profiles ./cmd/profiles/main.go

EXPOSE 8082 6379 8081

CMD ["./profiles"]