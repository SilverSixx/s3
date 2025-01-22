package api

import (
	"log"
	"net/http"
	"time"

	"github.com/silversixx/s3-go/internal/interfaces"
	"github.com/silversixx/s3-go/pkg/middleware"
)

type H struct {
	f func(http.ResponseWriter, *http.Request)
}

func (h H) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.f(w, r)
}

type ApiServer struct {
	m *http.ServeMux
}

func (s *ApiServer) InitServer() {
	s.m = http.NewServeMux()
	interfaces.InitNewRoute(s.m)

}

func (s *ApiServer) Run() {

	srv := &http.Server{
		Handler:      middleware.LoggingMiddleware(middleware.CorsMiddleware(s.m)),
		Addr:         "0.0.0.0:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
