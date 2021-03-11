FROM golang:latest


RUN go version
ENV GOPATH=/

COPY ./ ./

RUN go build -o music ./cmd/music/music.go

EXPOSE 8080 8081

CMD ["./music"]