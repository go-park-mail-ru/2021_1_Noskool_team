FROM golang:latest


RUN go version
ENV GOPATH=/

COPY ./ ./

RUN go build -o sessions ./cmd/sessions/sessions.go

EXPOSE 8081 6379 8080 8082

CMD ["./sessions"]