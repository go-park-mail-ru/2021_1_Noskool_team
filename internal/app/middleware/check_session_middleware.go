package middleware

import (
	"2021_1_Noskool_team/internal/microservices/auth/delivery/grpc/client"
	"context"
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"
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
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		userID := sessionID.Value
		session, err := sessMiddleware.sessionsClient.Check(context.Background(), userID)
		fmt.Println("Result: = " + session.Status)

		if err == nil {
			//w.Write([]byte("Кука есть и id у нее = " + strconv.Itoa(session.ID) + "\n"))
		} else {
			fmt.Printf("Error in checking session: %v\n", err)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		r = r.WithContext(context.WithValue(r.Context(), "user_id", session)) //nolint
		next.ServeHTTP(w, r)
	})
}

func (sessMiddleware *SessionsMiddleware) CheckIsNotAuthorized(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sessionID, err := r.Cookie("session_id")
		if err != nil {
			logrus.Info("User not Authorized")
			next.ServeHTTP(w, r)
			return
		}
		userID := sessionID.Value
		session, err := sessMiddleware.sessionsClient.Check(context.Background(), userID)
		fmt.Println("Result: = " + session.Status)
		if err != nil || session.ID == "-1" {
			logrus.Info("User not Authorized")
			next.ServeHTTP(w, r)
			return
		}

		logrus.Error("Error, User is Authorized")
		w.WriteHeader(418)
	})
}
