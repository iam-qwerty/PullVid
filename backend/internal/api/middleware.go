package api

import (
	"net/http"
	"strings"
)

func CORS(allowedOriginsStr string) func(http.Handler) http.Handler {
	allowedOrigins := strings.Split(allowedOriginsStr, ",")
	for i := range allowedOrigins {
		allowedOrigins[i] = strings.TrimSpace(allowedOrigins[i])
	}

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			origin := r.Header.Get("Origin")
			allowedOrigin := ""

			if allowedOriginsStr == "*" {
				allowedOrigin = "*"
			} else {
				for _, o := range allowedOrigins {
					if o == origin {
						allowedOrigin = origin
						break
					}
				}
			}

			if allowedOrigin != "" {
				w.Header().Set("Access-Control-Allow-Origin", allowedOrigin)
			}

			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

			if r.Method == http.MethodOptions {
				w.WriteHeader(http.StatusNoContent)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
