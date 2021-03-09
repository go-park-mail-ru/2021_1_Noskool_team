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
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Header().Set("Access-Control-Allow-Origin", "http://178.154.245.200:3000")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		if r.Method == "OPTIONS" {
			return
		}
		next.ServeHTTP(w, r)
	})
}

//func MyCORSMethodMiddleware(_ *mux.Router) mux.MiddlewareFunc {
//	return func(next http.Handler) http.Handler {
//		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
//			fmt.Printf("URL: %s, METHOD: %s", req.RequestURI, req.Method)
//			w.Header().Set("Access-Control-Allow-Headers", "*")
//			w.Header().Set("Access-Control-Allow-Origin", "http://178.154.245.200")
//			w.Header().Set("Access-Control-Allow-Credentials", "true")
//			w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
//			if req.Method == "OPTIONS" {
//				return
//			}
//			next.ServeHTTP(w, req)
//		})
//	}
//}