FROM golang:latest


RUN go version
ENV GOPATH=/

COPY ./ ./

RUN go build -o serv ./cmd/test/main.go

EXPOSE 8080 6379

CMD ["./serv"]