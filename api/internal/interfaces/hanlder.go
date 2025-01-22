package interfaces

import (
	"net/http"
)

type Handler interface {
	GetPath() string
	GetMiddleware(http.Handler, string) http.Handler
	Use(func(http.Handler) http.Handler)
}

type GET interface {
	Handler
	GET(w http.ResponseWriter, r *http.Request)
}

type POST interface {
	Handler
	POST(w http.ResponseWriter, r *http.Request)
}

type PUT interface {
	Handler
	PUT(w http.ResponseWriter, r *http.Request)
}

type DELETE interface {
	Handler
	DELETE(w http.ResponseWriter, r *http.Request)
}

type PATCH interface {
	Handler
	PATCH(w http.ResponseWriter, r *http.Request)
}

type OPTIONS interface {
	Handler
	OPTIONS(w http.ResponseWriter, r *http.Request)
}