package middleware

import (
	"2021_1_Noskool_team/internal/microservices/auth/delivery/grpc/client"
	"context"
	"fmt"
	"net/http"
	"strconv"
)

type SessionsMiddleware struct {
	sessionsClient client.AuthCheckerClient
}

func NewSessionMiddleware(grpcClient client.AuthCheckerClient) *SessionsMiddleware {
	sessMiddleware := &SessionsMiddleware{
		sessionsClient: grpcClient,
	}
	return sessMiddleware
}

func (sessMiddleware *SessionsMiddleware) CheckSessionMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sessionID, err := r.Cookie("session_id")

		if err != nil {
			fmt.Printf("Error in parsing cookie: %v\n", err)
			// http.Redirect(w, r, "/api/v1/musician/login/", http.StatusFound) //TODO::убрать редирект на стороне сервера
			return
		}

		userID, _ := strconv.Atoi(sessionID.Value)
		session, err := sessMiddleware.sessionsClient.Check(context.Background(), userID)
		fmt.Println("Result: = " + session.Status)

		if err == nil && session.ID == userID {
			w.Write([]byte("Кука есть и id у нее = " + strconv.Itoa(session.ID) + "\n"))
		} else {
			fmt.Printf("Error in checking session: %v\n", err)
			// http.Redirect(w, r, "/api/v1/musician/login/", http.StatusFound) //TODO::убрать редирект на стороне сервера
			return
		}
		next.ServeHTTP(w, r)
	})
}
