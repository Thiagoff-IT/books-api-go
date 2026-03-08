package middleware

import (
	"net/http"
	"os"
	"strings"
)

// CORS adiciona headers de Cross-Origin Resource Sharing.
// Origens permitidas podem ser configuradas via variável de ambiente CORS_ORIGINS (separadas por vírgula).
// Padrão: "*" (todas as origens).
func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origens := os.Getenv("CORS_ORIGINS")
		if origens == "" {
			origens = "*"
		}

		origin := r.Header.Get("Origin")
		if origens == "*" {
			w.Header().Set("Access-Control-Allow-Origin", "*")
		} else {
			for _, o := range strings.Split(origens, ",") {
				if strings.TrimSpace(o) == origin {
					w.Header().Set("Access-Control-Allow-Origin", origin)
					break
				}
			}
		}

		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}
