package middleware

import (
	"net/http"

	"github.com/spf13/viper"
)

func CorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get the origin from the request
		origin := r.Header.Get("Origin")

		// List of allowed origins
		allowedOrigins := viper.GetStringSlice("allow_origins")

		// Check if the origin is allowed
		allowedOrigin := ""
		for _, allowed := range allowedOrigins {
			if origin == allowed {
				allowedOrigin = allowed
				break
			}
		}

		// Set CORS headers
		w.Header().Set("Access-Control-Allow-Origin", allowedOrigin)
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE, PATCH")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, Client-ID")
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		// Handle preflight requests
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}
