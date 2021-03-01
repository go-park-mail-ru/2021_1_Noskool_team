FROM golang:1.13

WORKDIR /app

COPY . .

EXPOSE 8080 8081

RUN make build