package middleware

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"net/http"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		msg := fmt.Sprintf("URL: %s, METHOD: %s, REMOTE_ADDR%s",
			r.RequestURI, r.Method, r.RemoteAddr)
		logrus.Info(msg)
		next.ServeHTTP(w, r)
	})
}
