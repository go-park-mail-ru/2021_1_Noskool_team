package middleware

import (
	"2021_1_Noskool_team/configs"
	"fmt"
	"net/http"
)

type CORSMiddleware struct {
	config *configs.Config
}

func NewCORSMiddleware(config *configs.Config) *CORSMiddleware {
	return &CORSMiddleware{
		config: config,
	}
}
func (corsMiddlware *CORSMiddleware) CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("URL: %s, METHOD: %s", r.RequestURI, r.Method)
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length,"+
			" Accept-Encoding, X-CSRF-Token, csrf-token, Authorization")
		w.Header().Set("Access-Control-Allow-Origin", corsMiddlware.config.FrontendURL)
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Vary", "Origin")
		w.Header().Set("Vary", "Access-Control-Request-Method")
		w.Header().Set("Vary", "Access-Control-Request-Headers")
		w.Header().Set("Access-Control-Max-Age", "600")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}
