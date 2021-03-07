package middleware

import (
	"2021_1_Noskool_team/configs"
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
		w.Header().Set("Access-Control-Allow-Methods",
			"POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length,"+
			" Accept-Encoding, X-CSRF-Token, csrf-token, Authorization")
		w.Header().Set("Access-Control-Allow-Origin", corsMiddlware.config.FrontendURL) //TODO поставить url с фронта
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		next.ServeHTTP(w, r)
	})
}
