package interfaces

import (
	"net/http"
)

type H struct {
	f func(http.ResponseWriter, *http.Request)
}

func (h H) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.f(w, r)
}

var ListApi []Handler = []Handler{}

func AddRoute(hs ...Handler) {
	for _, h := range hs {
		ListApi = append(ListApi, h)
	}
}

func InitNewRoute(m *http.ServeMux) {

	for _, api := range ListApi {
		if api, ok := api.(GET); ok {
			m.HandleFunc("GET "+api.GetPath(), api.GetMiddleware(H{f: api.GET}, "GET").ServeHTTP)
		}
		if api, ok := api.(POST); ok {
			m.HandleFunc("POST "+api.GetPath(), api.GetMiddleware(H{f: api.POST}, "POST").ServeHTTP)
		}
		if api, ok := api.(DELETE); ok {
			m.HandleFunc("DELETE "+api.GetPath(), api.GetMiddleware(H{f: api.DELETE}, "DELTE").ServeHTTP)
		}
		if api, ok := api.(OPTIONS); ok {
			m.HandleFunc("OPTIONS "+api.GetPath(), api.GetMiddleware(H{f: api.OPTIONS}, "OPTIONS").ServeHTTP)
		}
		if api, ok := api.(PATCH); ok {
			m.HandleFunc("PATCH "+api.GetPath(), api.GetMiddleware(H{f: api.PATCH}, "PATCH").ServeHTTP)
		}
		if api, ok := api.(PUT); ok {
			m.HandleFunc("PUT "+api.GetPath(), api.GetMiddleware(H{f: api.PUT}, "PATCH").ServeHTTP)
		}
	}
}