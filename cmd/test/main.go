package main

import (
	"2021_1_Noskool_team/internal/microservices/auth/delivery/http"
	"2021_1_Noskool_team/internal/microservices/auth/usecase"
	"2021_1_Noskool_team/internal/pkg/server"
	"fmt"
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
