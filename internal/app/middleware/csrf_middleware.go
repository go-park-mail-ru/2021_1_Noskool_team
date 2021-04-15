package middleware

import (
	"2021_1_Noskool_team/internal/pkg/response"
	"github.com/sirupsen/logrus"
	"net/http"
)

func CheckCSRFMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		csrfFromHeader := r.Header.Get("X-Csrf-Token")
		csrfFromCookie, err := r.Cookie("csrf")
		if err != nil || csrfFromCookie.Value != csrfFromHeader {
			logrus.Info("Not CSRF Token")
			response.SendEmptyBody(w, http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	})
}
