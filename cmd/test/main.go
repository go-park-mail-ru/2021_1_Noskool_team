package main

import (
	"fmt"
	"testWorkWithAuth/internal/microservices/auth/delivery/http"
	"testWorkWithAuth/internal/microservices/auth/usecase"
	"testWorkWithAuth/internal/pkg/server"
)

func main() {
	config := server.Config{BindAddr: ":8080"}
	sessionUsecase := usecase.NewSessionsUsecase()
	handler := http.NewSessionHandler(&sessionUsecase)
	err := server.Start(config, handler)
	if err != nil {
		fmt.Println("fail start server")
	}
}
