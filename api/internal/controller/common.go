package controller

import (
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
)

const (
	DEFAULT_PAGE_SIZE = 20
)

type H func(http.Handler) http.Handler

// V is T DTO
type CommonController[T interface{}, V interface{}] struct {
	Path       string
	Middleware []H

	MiddlewareGET     []H
	MiddlewarePOST    []H
	MiddlewarePUT     []H
	MiddlewarePATCH   []H
	MiddlewareDELETE  []H
	MiddlewareHEAD    []H
	MiddlewareOPTIONS []H
}

func (h *CommonController[T, V]) GetPath() string {
	return h.Path
}

func (h *CommonController[T, V]) GetMiddleware(mainHandler http.Handler, method string) http.Handler {
	switch method {
	case "GET":
		for i := len(h.MiddlewareGET); i > 0; i-- {
			mainHandler = h.MiddlewareGET[i-1](mainHandler)
		}
		break
	case "POST":
		for i := len(h.MiddlewarePOST); i > 0; i-- {
			mainHandler = h.MiddlewarePOST[i-1](mainHandler)
		}
		break
	case "PUT":
		for i := len(h.MiddlewarePUT); i > 0; i-- {
			mainHandler = h.MiddlewarePUT[i-1](mainHandler)
		}
		break
	case "PATCH":
		for i := len(h.MiddlewarePATCH); i > 0; i-- {
			mainHandler = h.MiddlewarePATCH[i-1](mainHandler)
		}
		break
	}
	for i := len(h.Middleware); i > 0; i-- {
		mainHandler = h.Middleware[i-1](mainHandler)
	}
	return mainHandler
}

func (h *CommonController[T, V]) Use(middleware func(http.Handler) http.Handler) {
	h.Middleware = append(h.Middleware, middleware)
}

func (h *CommonController[T, V]) GetPaging(r *http.Request) (int, int) {
	page, err := strconv.ParseInt(r.URL.Query().Get("page"), 10, 32)
	if err != nil {
		page = 1
	}
	size, err := strconv.ParseInt(r.URL.Query().Get("size"), 10, 32)
	if err != nil {
		size = DEFAULT_PAGE_SIZE
	}
	return int(page), int(size)
}

func (h *CommonController[T, V]) Validate(i interface{}) error {
	validate := validator.New(validator.WithRequiredStructEnabled())
	return validate.Struct(i)
}