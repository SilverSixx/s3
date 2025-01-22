package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/silversixx/s3-go/pkg/logger"
)

func MethodNotAllowedHandler(w http.ResponseWriter, r *http.Request) {
	logger.Error(fmt.Sprintf("405 Method Not Allowed: %s %s", r.Method, r.URL.Path))
	w.Write([]byte(fmt.Sprintf("405 Method Not Allowed: %s %s", r.Method, r.URL.Path)))
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	logger.Error(fmt.Sprintf("404 Not Found: %s %s", r.Method, r.URL.Path))
	w.Write([]byte(fmt.Sprintf("404 Not Found: %s %s", r.Method, r.URL.Path)))
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// TODO: add health to path that not logged list
		if r.URL.Path != "/health" {
			logRequestDetails := []interface{}{
				r.Method, r.URL.Path, r.UserAgent(),
			}

			logger.InfoMultiField(append(
				[]interface{}{"Incoming request"},
				logRequestDetails...,
			)...)

			next.ServeHTTP(w, r)
			duration := time.Since(start)

			logResponseDetails := append(
				logRequestDetails, duration, r.Response,
			)

			logger.InfoMultiField(append(
				[]interface{}{"Outgoing response"},
				logResponseDetails...,
			)...)
		}
	})
}
