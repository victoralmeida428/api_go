package middleware

import (
	"api/src/config"
	"fmt"
	"net/http"
)

// Middleware para adicionar CORS a todas as requisições
func EnableCors(next http.Handler, cfg *config.Config) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE, PATCH")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Handle OPTIONS preflight requests
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		//printa as informaçoes do request
		msg := fmt.Sprintf("remote address: %s,  uri: %s", r.RemoteAddr, r.RequestURI)
		cfg.Logger.PrintInfo(msg, nil)

		// Chama o próximo handler
		next.ServeHTTP(w, r)
	})
}