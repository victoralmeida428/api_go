package middleware

import (
	"api/src/config"
	"fmt"
	"net/http"
)

func RecoverPanic(next http.Handler, cfg *config.Config) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if err := recover(); err != nil {
					w.Header().Set("Connection", "close")
					cfg.Error.ServerErrorResponse(w, r, fmt.Errorf("%s", err))
				}
			}()
			next.ServeHTTP(w, r)
		})
}
